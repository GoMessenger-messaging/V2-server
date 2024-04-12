package db

import (
	"git.jereileu.ch/gomessenger/gomessenger/server/conf"
	"git.jereileu.ch/gotables/client/go/gotables"
)

func Initialize(conf conf.Conf) {
	go gotables.RunServer(conf.DBConf)
}

func TestConnection(conf conf.Conf) error {
	var err error
	err = gotables.TestServer(conf.DBConf)
	return err
}
