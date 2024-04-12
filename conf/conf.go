package conf

import (
	"encoding/json"
	"git.jereileu.ch/gotables/client/go/gotables"
	"os"
)

type Conf struct {
	DBConf  gotables.Config `json:"db_conf"`
	Port    string          `json:"port"`
	HTTPS   bool            `json:"https"`
	SSLCert string          `json:"ssl_cert"`
	SSLKey  string          `json:"ssl_key"`
}

func Load(location string) (Conf, error) {
	data, err := os.ReadFile(location)
	if err != nil {
		return Conf{}, err
	}
	var conf Conf
	err = json.Unmarshal(data, &conf)
	return conf, err
}
