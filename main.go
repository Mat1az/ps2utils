package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"ps2-utils/models"
	"ps2-utils/services"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	gameService := services.NewGameService()

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "ps2-utils",
		Width:             1200,
		Height:            600,
		HideWindowOnClose: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			gameService.Startup(ctx)
		},
		Bind: []interface{}{
			app,
			gameService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func (a *App) SelectFile() []models.Game {
	fmt.Println("Test????????")
	return []models.Game{}
}
