package main

import (
	"log"

	"github.com/ahmadalaik/music-catalog/internal/configs"
	"github.com/ahmadalaik/music-catalog/internal/models/memberships"
	membershipsRepo "github.com/ahmadalaik/music-catalog/internal/repository/memberships"
	membershipsSvc "github.com/ahmadalaik/music-catalog/internal/service/memberships"
	membershipsHandler "github.com/ahmadalaik/music-catalog/internal/handler/memberships"
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
		log.Fatalf("Failed connect database, err: %+v", err)
	}
	db.AutoMigrate(&memberships.User{})

	r := gin.Default()

	membershipRepo := membershipsRepo.NewRepository(db)
	
	membershipSvc := membershipsSvc.NewService(config, membershipRepo)

	memberhipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	memberhipHandler.RegisterRoute()

	r.Run(config.Service.Port)
}
