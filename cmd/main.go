package main

import (
	"log"

	"github.com/ahmadalaik/music-catalog/internal/configs"
	"github.com/ahmadalaik/music-catalog/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	var config *configs.Config

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Failed initialize config ", err)
	}
	config = configs.Get()
	log.Println("config", config)

	db, err := internalsql.Connect(config.Database.DataSourceName)
	if err != nil {
		log.Fatal("Failed connect database, err: %+v", err)
	}

	r := gin.Default()

	r.Run(config.Service.Port)
}
