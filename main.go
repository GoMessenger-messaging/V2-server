package main

import (
	"git.jereileu.ch/gomessenger/gomessenger/gm-server/conf"
	"git.jereileu.ch/gomessenger/gomessenger/gm-server/db"
	"git.jereileu.ch/gomessenger/gomessenger/gm-server/server"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	ConfigEnvvar = "GM_CONFIG"
	Version      = "0.0.0"
)

func main() {
	log.Println("loading config...")
	location := os.Getenv(ConfigEnvvar)
	config, err := conf.Load(location)
	if err != nil {
		log.Println("warning: failed to find config file! using default configuration!")
	} else {
		log.Println("successfully loaded config!")
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
	log.Println("successfully contacted dbms!")
	log.Println("==================== GoTables server " + Version + " ====================")
	if config.HTTPS {
		log.Println("server started at https://127.0.0.1" + config.Port)
	} else {
		log.Println("server started at http://127.0.0.1" + config.Port)
	}
	log.Println("press 'Ctrl' + 'C' to stop this program")
	end := ""
	for i := 0; i < 58+len(Version); i++ {
		end += "="
	}
	log.Println(end)
	log.Println("")
	server.Run(config)
}
