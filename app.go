package main

import (
	"context"
	"log"

	"github.com/alecthomas/repr"
	"github.com/sleepinggenius2/gosmi/parser"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/willowbrowser/snmpmibbrowser/internal/oidstorage"
)

// App struct
type App struct {
	ctx        context.Context
	loadedOids *oidstorage.LoadedOids
}

// NewApp creates a new App application struct
func NewApp() *App {
	loadedOids := oidstorage.NewLoadedOids()

	return &App{
		loadedOids: loadedOids,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
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
