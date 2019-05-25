package main

import (
	"encoding/gob"
	"go-credital/config"
	"go-credital/db/mongo"
	models "go-credital/models/mongo"
	"go-credital/routes"
	"log"
)

func main() {
	config.Init()
	mongo.InitMongo()
	gob.Register(models.SysUser{})

	e := routes.Init()

	log.Println("/----------------------------------------------------/")
	log.Println("/-------------- CREDITAL PANEL ----------------------/")
	log.Println("/--------------  BY BANG HAJI  ----------------------/")
	log.Println("/----------------------------------------------------/")

	conf := config.GetConfig()
	e.Logger.Fatal(e.Start(conf.Port))
}
