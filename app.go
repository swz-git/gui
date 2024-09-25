package main

import (
	"context"
	"encoding/base64"
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/ncruces/zenity"
	rlbot "github.com/swz-git/go-interface"
	"github.com/swz-git/go-interface/flat"
	"github.com/wailsapp/mimetype"
)

// App struct
type App struct {
	ctx context.Context
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
		if info.Name() == targetName {
			matches = append(matches, path)
		}
		return nil
	})
	return matches, err
}

func BotInfoToPlayerConfig(botInfo BotInfo, team uint32) *flat.PlayerConfigurationT {
	var runCommand string
	if runtime.GOOS == "windows" {
		runCommand = botInfo.Config.Settings.RunCommand
	} else if runtime.GOOS == "linux" {
		runCommand = botInfo.Config.Settings.RunCommandLinux
	}

	return &flat.PlayerConfigurationT{
		Variety: &flat.PlayerClassT{
			Type:  flat.PlayerClassRLBot,
			Value: &flat.RLBotT{},
		},
		Name:       botInfo.Config.Settings.Name,
		Team:       team,
		Location:   botInfo.Config.Settings.Location,
		RunCommand: runCommand,
		// TODO: read player loadout file from LooksConfig
		Loadout:  &flat.PlayerLoadoutT{},
		SpawnId:  0, // let core do this
		Hivemind: botInfo.Config.Settings.Hivemind,
	}
}

type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type StartMatchOptions struct {
	Map             string                `json:"map"`
	GameMode        string                `json:"gameMode"`
	BlueBots        []BotInfo             `json:"blueBots"`
	OrangeBots      []BotInfo             `json:"orangeBots"`
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

	for _, botInfo := range options.BlueBots {
		playerConfigs = append(playerConfigs, BotInfoToPlayerConfig(botInfo, 0))
	}
	for _, botInfo := range options.OrangeBots {
		playerConfigs = append(playerConfigs, BotInfoToPlayerConfig(botInfo, 1))
	}

	conn.SendPacket(&flat.MatchSettingsT{
		AutoStartBots:        true,
		GameMapUpk:           options.Map,
		PlayerConfigurations: playerConfigs,
		GameMode:             gameMode,
		MutatorSettings: &flat.MutatorSettingsT{
			MatchLength: flat.MatchLengthUnlimited,
		},
		EnableRendering:    true,
		EnableStateSetting: true,
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

type BotInfo struct {
	Config   BotConfig `json:"config"`
	TomlPath string    `json:"tomlPath"`
}

type BotConfig struct {
	Settings BotSettings `toml:"settings" json:"settings"`
	Details  BotDetails  `toml:"details" json:"details"`
}

type BotSettings struct {
	// In-game name of the bot
	Name string `toml:"name" json:"name"`
	// Path to looks.toml, describing the bots "loadout"
	LooksConfig string `toml:"looks_config" json:"looksConfig"`
	// Optional working dir of the bot
	Location string `toml:"location" json:"location"`
	// Path to the logo of the bot, if ignored, RLBot will look for logo.png
	LogoFile string `toml:"logo_file" json:"logoFile"`
	// The command RLBot will call to start your bot on Windows
	RunCommand string `toml:"run_command" json:"runCommand"`
	// The command RLBot will call to start your bot on Linux
	// If not defined, RLBot may try to run your bot under wine
	RunCommandLinux string `toml:"run_command_linux" json:"runCommandLinux"`
	// If bot can handle multiple agents with one client
	Hivemind bool `toml:"hivemind" json:"hivemind"`
}

type BotDetails struct {
	// Short description of thebot
	Description string `toml:"description" json:"description"`
	// A fun fact about the bot
	FunFact string `toml:"fun_fact" json:"funFact"`
	// Link to the source code of the bot (if its avaliable)
	SourceLink string `toml:"source_link" json:"sourceLink"` // TODO: Rename this field to repo?
	// Name(s) of the bot developer(s)
	Developer string `toml:"developer" json:"developer"`
	// Programming language the bot is written in.
	// (RLGym for example is also valid even though it is written in Python)
	Language string `toml:"language" json:"language"`
	// ALL POSSIBLE TAGS: 1v1, teamplay, goalie, hoops, dropshot, snow-day, spike-rush, heatseaker, memebot
	// NOTE: Only add the goalie tag if your bot only plays as a goalie; this directly contrasts with the teamplay tag!
	// NOTE: Only add a tag for a special game mode if you bot properly supports it
	Tags []string `toml:"Tags" json:"tags"`
}

func (a *App) GetBots(paths []string) []BotInfo {
	// TODO: Search recusrive in paths
	potentialConfigs := []string{}

	for _, path := range paths {
		new, err := recursiveFileSearch(path, "bot.toml")
		if err != nil {
			println("WARN: failed to search path: " + path)
			continue
		}
		potentialConfigs = append(potentialConfigs, new...)
	}

	infos := []BotInfo{}

	for _, potentialConfigPath := range potentialConfigs {
		data, err := os.ReadFile(potentialConfigPath)
		if err != nil {
			println("WARN: skipping config, couldn't read config at " + potentialConfigPath)
			continue
		}
		var conf BotConfig
		toml.Decode(string(data), &conf)

		// make location path relative to parent of bot.toml
		conf.Settings.Location = filepath.Join(filepath.Dir(potentialConfigPath), conf.Settings.Location)

		// Read logo file and convert it to data url so the frontend can use it
		if conf.Settings.LogoFile != "" {
			conf.Settings.LogoFile = filepath.Join(conf.Settings.Location, conf.Settings.LogoFile)
			logo_data, err := os.ReadFile(conf.Settings.LogoFile)
			if err != nil {
				println("WARN: failed to read logo file at " + conf.Settings.LogoFile)
			} else {
				mtype := mimetype.Detect(logo_data)
				b64data := base64.StdEncoding.EncodeToString(logo_data)
				conf.Settings.LogoFile = "data:" + mtype.String() + ";base64," + b64data
			}
		}

		infos = append(infos, BotInfo{
			Config:   conf,
			TomlPath: potentialConfigPath,
		})
	}

	return infos
}
