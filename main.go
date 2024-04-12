package main

import (
	"git.jereileu.ch/gomessenger/gomessenger/server/conf"
	"git.jereileu.ch/gomessenger/gomessenger/server/db"
	"git.jereileu.ch/gomessenger/gomessenger/server/server"
	"log"
	"os"
)

const (
	ConfigEnvvar = "GM_CONFIG"
)

func main() {
	log.Println("loading config...")
	location := os.Getenv(ConfigEnvvar)
	config, err := conf.Load(location)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("initialising database...")
	db.Initialize(config)
	log.Println("Successfully initialised database!")
	log.Println("trying to contact database...")
	err = db.TestConnection(config)
	if err != nil {
		log.Fatalln("failed to establish a connection with the database!")
	}
	server.Run(config)
}
