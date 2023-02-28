package agent

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	g "github.com/gosnmp/gosnmp"
	"github.com/xujiajun/nutsdb"

	"github.com/willowbrowser/snmpmibbrowser/internal/oidstorage"
)

const (
	bucket = "agents"
)

type AgentStorage struct {
	agents       []AgentObj
	currentAgent *g.GoSNMP
	db           *oidstorage.DB
}

type AgentObj struct {
	Name          string `json:"name"`
	OriginalInput InputType
	agent         *g.GoSNMP
}

func NewAgentStore(db *oidstorage.DB) *AgentStorage {
	gob.Register(AgentObj{})
	var newAgents []AgentObj

	agentStore := AgentStorage{
		db: db,
	}

	if err := db.Test().View(func(tx *nutsdb.Tx) error {
		entries, err := tx.GetAll(bucket)
		if err != nil {
			if err.Error() == "bucket is empty" {
				return nil
			}
			return err
		}

		for _, entry := range entries {
			input := entry.Value
			buf := bytes.NewBuffer(input)
			dec := gob.NewDecoder(buf)

			var a AgentObj
			err := dec.Decode(&a)
			// TODO : when decoding, the pointer is always nil...duh it's a pointer
			if err != nil {
				log.Printf("decode error: %v", err)
				// return nil
			} else {
				a.agent = agentStore.createAgent(a.OriginalInput)

				if err := a.agent.Connect(); err != nil {
					log.Fatalf("Connect() err: %v", err)
				}

				newAgents = append(newAgents, a)
			}
		}
		return nil
	}); err != nil {
		log.Fatalln(err)
	}

	agentStore.agents = newAgents

	return &agentStore
}

func (a *AgentStorage) GetAllCurrentAgents() []AgentObj {
	return a.agents
}

func (a *AgentStorage) CloseAllConnections() {
	for _, agent := range a.agents {
		agent.agent.Conn.Close()
	}
}

func (a *AgentStorage) SelectCurrentAgent(name string) {
	for _, agent := range a.agents {
		if agent.Name == name {
			a.currentAgent = agent.agent
			return
		}
	}

	log.Fatalf("Invalid agent selected")
}

func (a *AgentStorage) PerformSnmpGet() {
	oids := []string{"1.3.6.1.2.1.1.4.0"}

	result, err := a.currentAgent.Get(oids)
	if err != nil {
		log.Fatalf("Get() err: %v", err)
	}

	for i, variable := range result.Variables {
		log.Printf("%d: oid: %s ", i, variable.Name)

		switch variable.Type {
		case g.OctetString:
			log.Printf("string: %s\n", string(variable.Value.([]byte)))
		default:
			log.Printf("number: %d\n", g.ToBigInt(variable.Value))
		}
	}
}

func (a *AgentStorage) CreateNewAgent(input InputType) {
	var agent *g.GoSNMP

	if input.AgentType.Id == 3 {
		agent = a.createV3_Agent(input)
	}

	agentObj := AgentObj{
		Name:          input.AgentName,
		OriginalInput: input,
		agent:         agent,
	}

	if err := a.db.Test().Update(func(tx *nutsdb.Tx) error {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		enc.Encode(&agentObj)
		key := []byte(input.AgentName)
		value := buf.Bytes()
		if err := tx.Put(bucket, key, value, 0); err != nil {
			return err
		}

		newBuf := bytes.NewBuffer(buf.Bytes())
		dec := gob.NewDecoder(newBuf)

		var a AgentObj
		dec.Decode(&a)

		return nil
	}); err != nil {
		log.Fatalf("Update error: %v", err)
	}

	a.agents = append(a.agents, agentObj)
}

func (a *AgentStorage) createAgent(input InputType) *g.GoSNMP {
	var agent *g.GoSNMP

	if input.AgentType.Id == 3 {
		agent = a.createV3_Agent(input)
	}

	return agent
}

func (a *AgentStorage) createV3_Agent(input InputType) *g.GoSNMP {
	var authType g.SnmpV3MsgFlags

	if input.AuthType.Name == "noAuthNoPriv" {
		authType = g.NoAuthNoPriv
	} else if input.AgentType.Name == "authNoPriv" {
		authType = g.AuthNoPriv
	} else if input.AuthType.Name == "authPriv" {
		authType = g.AuthPriv
	}

	agent := &g.GoSNMP{
		Target:        input.AgentAddress,
		Port:          input.AgentPort,
		Version:       g.Version3,
		SecurityModel: g.UserSecurityModel,
		MsgFlags:      authType,
		Timeout:       time.Duration(5) * time.Second,
		SecurityParameters: &g.UsmSecurityParameters{
			UserName:                 input.UsmUserName,
			AuthenticationProtocol:   g.NoAuth,
			AuthenticationPassphrase: input.AuthKey,
			PrivacyProtocol:          g.NoPriv,
			PrivacyPassphrase:        input.PrivKey,
		},
	}

	if err := agent.Connect(); err != nil {
		log.Fatalf("Connect() err: %v", err)
	}

	log.Printf("Agent create!")

	return agent
}
