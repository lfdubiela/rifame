package main

import (
	"github.com/go-ini/ini"
	"log"
	"rifame/driven"
	"rifame/driving"
)

type (
	Settings struct {
		Driven  driven.Settings
		Driving driving.Settings
	}

	Application struct {
		Driven   driven.Driven
		Driving  driving.Driving
		Settings Settings
	}
)

func main() {
	cfg, err := ini.Load("resources/app.ini")
	if err != nil {
		log.Fatalf("fail to read file: %v", err)
	}

	server := driving.ServerSettings{}
	database := driven.DatabaseSettings{}

	mapTo(cfg, "server", &server)
	mapTo(cfg, "database", &database)

	app := &Application{
		Driven:  driven.Driven{},
		Driving: driving.Driving{},

		Settings: Settings{
			Driving: driving.Settings{Server: server},
			Driven:  driven.Settings{Database: database},
		},
	}

	app.Driven.Setup(app.Settings.Driven)
	app.Driving.Setup(app.Driven, app.Settings.Driving)
}

// mapTo map section
func mapTo(cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
