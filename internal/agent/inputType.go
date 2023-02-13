package agent

type InputType struct {
	AgentAddress   string
	AgentPort      uint16
	AgentType      SelectedType
	ReadCommunity  string
	WriteCommunity string
	UsmUserName    string
	AuthType       SelectedType
	AuthKey        string
	PrivKey        string
}

type SelectedType struct {
	Id   uint32
	Name string
}
