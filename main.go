package main

import (
	"encoding/gob"
	"go-credital/config"
	"go-credital/db"
	"go-credital/models"
	"go-credital/routes"
	"go-credital/utils"
	"log"
)

func main() {
	config.Init()
	db.InitMongo()
	utils.InitSettings()
	gob.Register(models.SysUser{})

	e := routes.Init()

	log.Println("/----------------------------------------------------/")
	log.Println("/-------------- CREDITAL PANEL ----------------------/")
	log.Println("/--------------  BY BANG HAJI  ----------------------/")
	log.Println("/----------------------------------------------------/")

	conf := config.GetConfig()
	e.Logger.Fatal(e.Start(conf.Port))
}
