package config

import (
	"log"
	"os"

	"github.com/tkanos/gonfig"
)

//Configuration app
type Configuration struct {
	MongoHost   string
	MongoDBName string
	Port        string
}

var configuration Configuration

//Init config
func Init() {
	err := gonfig.GetConf("app.json", &configuration)
	if err != nil {
		log.Println(err)
		os.Exit(500)
	}
}

//GetConfig file configuration
func GetConfig() Configuration {
	return configuration
}
