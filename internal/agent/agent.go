package agent

import (
	"log"
	"time"

	g "github.com/gosnmp/gosnmp"
)

type AgentStorage struct {
	agents []*g.GoSNMP
}

func NewAgentStore() *AgentStorage {
	return &AgentStorage{}
}

func (a *AgentStorage) CloseAllConnections() {
	for _, agent := range a.agents {
		agent.Conn.Close()
	}
}

func (a *AgentStorage) CreateNewAgent(input InputType) {
	var agent *g.GoSNMP

	if input.AgentType.Id == 3 {
		agent = a.createV3_Agent(input)
	}

	a.agents = append(a.agents, agent)
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
