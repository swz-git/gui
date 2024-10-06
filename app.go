package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/ncruces/zenity"
	rlbot "github.com/swz-git/go-interface"
	"github.com/swz-git/go-interface/flat"
)

// App struct
type App struct {
	ctx context.Context
}

func (a *App) IgnoreMe(
	_ BotInfo,
	_ PsyonixBotInfo,
	_ HumanInfo,
) {
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// // Greet returns a greeting for the given name
// func (a *App) Greet(name string) string {
// 	return fmt.Sprintf("Hello %s, It's show time!", name)
// }

func recursiveFileSearch(root, targetName string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(info.Name(), targetName) {
			matches = append(matches, path)
		}
		return nil
	})
	return matches, err
}

type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type StartMatchOptions struct {
	Map             string                `json:"map"`
	GameMode        string                `json:"gameMode"`
	BluePlayers     []PlayerJs            `json:"bluePlayers"`
	OrangePlayers   []PlayerJs            `json:"orangePlayers"`
	MutatorSettings flat.MutatorSettingsT `json:"mutatorSettings"`
}

func (a *App) StartMatch(options StartMatchOptions) Result {
	// TODO: Save this in App struct
	// TODO: Make dynamic, pull from env var?
	conn, err := rlbot.Connect("127.0.0.1:23234")
	if err != nil {
		return Result{false, "Failed to connect to rlbot"}
	}

	var gameMode flat.GameMode
	switch options.GameMode {
	case "Soccer":
		gameMode = flat.GameModeSoccer
	case "Hoops":
		gameMode = flat.GameModeHoops
	case "Dropshot":
		gameMode = flat.GameModeDropshot
	case "Hockey":
		gameMode = flat.GameModeHockey
	case "Rumble":
		gameMode = flat.GameModeRumble
	case "Heatseeker":
		gameMode = flat.GameModeHeatseeker
	case "Gridiron":
		gameMode = flat.GameModeGridiron
	default:
		println("No mode chosen, defaulting to soccer")
		gameMode = flat.GameModeSoccer
	}

	var playerConfigs []*flat.PlayerConfigurationT

	for _, playerInfo := range options.BluePlayers {
		println(playerInfo.ToPlayer().ToPlayerConfig(0))
		playerConfigs = append(playerConfigs, playerInfo.ToPlayer().ToPlayerConfig(0))
	}
	for _, playerInfo := range options.OrangePlayers {
		playerConfigs = append(playerConfigs, playerInfo.ToPlayer().ToPlayerConfig(1))
	}

	println(playerConfigs)

	conn.SendPacket(&flat.MatchSettingsT{
		AutoStartBots:        true,
		GameMapUpk:           options.Map,
		PlayerConfigurations: playerConfigs,
		GameMode:             gameMode,
		MutatorSettings:      &options.MutatorSettings,
		EnableRendering:      true,
		EnableStateSetting:   true,
	})

	return Result{true, ""}
}

func (a *App) StopMatch(shutdownServer bool) Result {
	// TODO: Save this in App struct
	// TODO: Make dynamic, pull from env var?
	conn, err := rlbot.Connect("127.0.0.1:23234")
	if err != nil {
		return Result{false, "Failed to connect to rlbot"}
	}

	conn.SendPacket(&flat.StopCommandT{
		ShutdownServer: shutdownServer,
	})

	return Result{true, ""}
}

func (a *App) PickFolder() string {
	path, err := zenity.SelectFile(zenity.Directory())
	if err != nil {
		println("ERR: File picker failed")
	}
	return path
}
