package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/swz-git/go-interface/flat"
	"github.com/wailsapp/mimetype"
)

type PlayerJs struct {
	Sort   string          `json:"sort"`
	Player json.RawMessage `json:"player"`
}

func (self PlayerJs) ToPlayer() Player {
	switch self.Sort {
	case "rlbot":
		var correct BotInfo
		if err := json.Unmarshal([]byte(self.Player), &correct); err != nil {
			log.Fatal("unable to unmarshal PlayerJs")
		}
		return correct
	case "psyonix":
		var correct PsyonixBotInfo
		if err := json.Unmarshal([]byte(self.Player), &correct); err != nil {
			log.Fatal("unable to unmarshal PlayerJs")
		}
		return correct
	case "human":
		var correct HumanInfo
		if err := json.Unmarshal([]byte(self.Player), &correct); err != nil {
			log.Fatal("unable to unmarshal PlayerJs")
		}
		return correct
	}
	log.Println("ERROR: invalid sort field in PlayerJs")
	return PsyonixBotInfo{}
}

type Player interface {
	ToPlayerConfig(team uint32) *flat.PlayerConfigurationT
}

type PsyonixBotInfo struct {
	// Beginner: 0, Rookie: 1, Pro: 2, AllStar: 3
	Skill byte `json:"skill"`
}

func (info PsyonixBotInfo) ToPlayerConfig(team uint32) *flat.PlayerConfigurationT {
	return &flat.PlayerConfigurationT{
		Variety: &flat.PlayerClassT{
			Type: flat.PlayerClassPsyonix,
			Value: &flat.PsyonixT{
				BotSkill: flat.PsyonixSkill(info.Skill),
			},
		},
		Name:       "Psyonix Bot",
		Team:       team,
		Location:   "",
		RunCommand: "",
		Loadout:    &flat.PlayerLoadoutT{},
		SpawnId:    0,
		Hivemind:   false,
	}
}

type HumanInfo struct{}

func (info HumanInfo) ToPlayerConfig(team uint32) *flat.PlayerConfigurationT {
	return &flat.PlayerConfigurationT{
		Variety: &flat.PlayerClassT{
			Type:  flat.PlayerClassHuman,
			Value: &flat.HumanT{},
		},
		Name:       "Human",
		Team:       team,
		Location:   "",
		RunCommand: "",
		Loadout:    &flat.PlayerLoadoutT{},
		SpawnId:    0,
		Hivemind:   false,
	}
}

type BotInfo struct {
	Config   BotConfig `json:"config"`
	TomlPath string    `json:"tomlPath"`
}

func (self BotInfo) ToPlayerConfig(team uint32) *flat.PlayerConfigurationT {
	var runCommand string
	if runtime.GOOS == "windows" {
		runCommand = self.Config.Settings.RunCommand
	} else if runtime.GOOS == "linux" {
		runCommand = self.Config.Settings.RunCommandLinux
	}

	return &flat.PlayerConfigurationT{
		Variety: &flat.PlayerClassT{
			Type:  flat.PlayerClassRLBot,
			Value: &flat.RLBotT{},
		},
		Name:       self.Config.Settings.Name,
		Team:       team,
		Location:   self.Config.Settings.Location,
		RunCommand: runCommand,
		// TODO: read player loadout file from LooksConfig
		Loadout:  &flat.PlayerLoadoutT{},
		SpawnId:  0, // let core do this
		Hivemind: self.Config.Settings.Hivemind,
	}
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
	Tags []string `toml:"tags" json:"tags"`
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
