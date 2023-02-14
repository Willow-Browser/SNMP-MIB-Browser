package agent

import (
	"log"
	"time"

	g "github.com/gosnmp/gosnmp"
)

type AgentStorage struct {
	agents []AgentObj
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
