package main

import (
	"flag"
	"fmt"
	"github.com/alhaos/quick-menu-api/internal/config"
	"github.com/alhaos/quick-menu-api/internal/controller"
	"github.com/alhaos/quick-menu-api/internal/database"
	"github.com/alhaos/quick-menu-api/internal/repository"
	"github.com/gin-gonic/gin"
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

	// Init repo
	repo := repository.New(dbConnection)

	// Init controller
	ctrl := controller.New(repo)

	// Init router
	router := gin.Default()
	controller.SetupRouter(router, ctrl)

	// Run service
	err = router.Run(fmt.Sprintf("%s:%d", conf.Address.Hostname, conf.Address.Port))
	if err != nil {
		panic(err)
	}

}
