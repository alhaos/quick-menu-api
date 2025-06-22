package main

import (
	"flag"
	"github.com/alhaos/quick-menu-api/internal/config"
	"github.com/alhaos/quick-menu-api/internal/database"
	"github.com/alhaos/quick-menu-api/internal/repo"
)

func main() {

	// Parse arguments
	configFileNamePointer := flag.String("config", "config.yml", "config file path")
	flag.Parse()
	filename := *configFileNamePointer

	// Init config
	conf, err := config.New(filename)
	if err != nil {
		panic(err)
	}

	// Init database
	dbConnection, err := database.New(conf.Database)
	if err != nil {
		panic(err)
	}

	defer dbConnection.Close()

	// Init repository
	_ = repo.New(dbConnection)

}
