package agent

import (
	"log"
	"time"

	g "github.com/gosnmp/gosnmp"
)

type AgentStorage struct {
	agents       []AgentObj
	currentAgent *g.GoSNMP
}

type AgentObj struct {
	Name  string `json:"name"`
	agent *g.GoSNMP
}

func NewAgentStore() *AgentStorage {
	return &AgentStorage{}
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
		Name:  input.AgentName,
		agent: agent,
	}

	a.agents = append(a.agents, agentObj)

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
