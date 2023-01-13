package main

import (
	"embed"
	"fmt"

	"github.com/willowbrowser/snmpmibbrowser/internal/mibreader"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), func(cd *menu.CallbackData) {
		file, err := runtime.OpenFileDialog(app.ctx, runtime.OpenDialogOptions{
			DefaultDirectory: "./",
			Title:            "Select File",
			Filters: []runtime.FileFilter{
				{
					DisplayName: "Text (*.txt)",
					Pattern:     "*.txt",
				},
			},
		})

		if err != nil {
			fmt.Errorf(fmt.Sprintf("Error opening file: %v", err))
		}

		if file != "" {
			mibReader := mibreader.NewMibReader(app.loadedOids)
			mibReader.ReadMib(file)
			runtime.EventsEmit(app.ctx, "mibsLoaded")
		}
	})
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(cd *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "test-wails",
		Width:     1280,
		Height:    780,
		Assets:    assets,
		OnStartup: app.startup,
		Menu:      AppMenu,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		fmt.Errorf(fmt.Sprintf("Error closing mib-reader: %v", err))
	}
}
