package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/alecthomas/repr"
	"github.com/sleepinggenius2/gosmi/parser"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/willowbrowser/snmpmibbrowser/internal/agent"
	"github.com/willowbrowser/snmpmibbrowser/internal/oidstorage"
)

// App struct
type App struct {
	ctx         context.Context
	loadedOids  *oidstorage.LoadedOids
	db          *oidstorage.DB
	agentStores *agent.AgentStorage
}

// NewApp creates a new App application struct
func NewApp() *App {
	db := oidstorage.InitializeDb()
	loadedOids := oidstorage.NewLoadedOids(db)

	return &App{
		loadedOids:  loadedOids,
		db:          db,
		agentStores: agent.NewAgentStore(db),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.EventsOn(a.ctx, "selectedAgent", func(optionalData ...interface{}) {
		data := optionalData[0]

		name := data.(map[string]interface{})["name"].(string)

		fmt.Printf("%s\n", name)

		a.agentStores.SelectCurrentAgent(name)
	})

	runtime.EventsOn(a.ctx, "createAgent", func(optionalData ...interface{}) {
		data := optionalData[0]

		agentTypeIdFloat := data.(map[string]interface{})["agentType"].(map[string]interface{})["id"].(float64)
		authTypeIdFloat := data.(map[string]interface{})["authType"].(map[string]interface{})["id"].(float64)

		agentType := agent.SelectedType{
			Id:   uint32(math.Round(agentTypeIdFloat)),
			Name: data.(map[string]interface{})["agentType"].(map[string]interface{})["name"].(string),
		}

		authType := agent.SelectedType{
			Id:   uint32(math.Round(authTypeIdFloat)),
			Name: data.(map[string]interface{})["authType"].(map[string]interface{})["name"].(string),
		}

		agentPortStr := data.(map[string]interface{})["agentPort"].(string)
		agentPortInt, _ := strconv.Atoi(agentPortStr)

		input := agent.InputType{
			AgentName:      data.(map[string]interface{})["agentName"].(string),
			AgentAddress:   data.(map[string]interface{})["agentAddress"].(string),
			AgentPort:      uint16(agentPortInt),
			AgentType:      agentType,
			ReadCommunity:  data.(map[string]interface{})["readCommunity"].(string),
			WriteCommunity: data.(map[string]interface{})["writeCommunity"].(string),
			UsmUserName:    data.(map[string]interface{})["usmUserName"].(string),
			AuthType:       authType,
			AuthKey:        data.(map[string]interface{})["authKey"].(string),
			PrivKey:        data.(map[string]interface{})["privKey"].(string),
		}

		a.agentStores.CreateNewAgent(input)

		runtime.EventsEmit(a.ctx, "agentsUpdated")
	})
}

func (a *App) shutdown(ctx context.Context) {
	a.db.CloseDb()
	a.agentStores.CloseAllConnections()
}

func (a *App) ParseMib() {
	result, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory: "./",
		Title:            "Select a mib",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Mib Files (*.*)",
				Pattern:     "*.*",
			},
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	if result != "" {
		module, err := parser.ParseFile(result)
		if err != nil {
			log.Fatalln(err)
		}

		repr.Println(module)
	}
}

func (a *App) GetCurrentOids() []oidstorage.Oid {
	return a.loadedOids.GetLoadedOids()
}

func (a *App) GetAllCurrentAgents() []agent.AgentObj {
	return a.agentStores.GetAllCurrentAgents()
}

// sending a simple get from selected agent
// selectable OID not yet available
func (a *App) SendGetRequest() {
	a.agentStores.PerformSnmpGet()
}

func (a *App) SendGetRequestWithOid(oid string) {
	a.agentStores.PerformSnmpGetOid(oid)
}
