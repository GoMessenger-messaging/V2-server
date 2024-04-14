package main

import (
	"git.jereileu.ch/gomessenger/gomessenger/server/conf"
	"git.jereileu.ch/gomessenger/gomessenger/server/db"
	"git.jereileu.ch/gomessenger/gomessenger/server/server"
	"log"
	"os"
	"strconv"
	"time"
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
	if config.DBMSRun {
		log.Println("initialising dbms...")
		db.Initialize(config)
		log.Println("successfully initialised dbms!")
		log.Println("waiting for the dbms to boot... (" + strconv.Itoa(config.DBMSBootTime) + "s)")
		time.Sleep(time.Duration(config.DBMSBootTime) * time.Second)
	}
	log.Println("trying to contact dbms...")
	err = db.TestConnection(config)
	if err != nil {
		log.Fatalln("failed to establish a connection with the dbms!")
	}
	server.Run(config)
}
