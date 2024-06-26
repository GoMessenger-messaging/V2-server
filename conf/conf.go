package conf

import (
	"encoding/json"
	"git.jereileu.ch/gotables/client/go/gotables"
	"git.jereileu.ch/gotables/server/gt-server/fs"
	"os"
)

type Conf struct {
	DBMSConf     gotables.Config `json:"dbms_conf"`
	Port         string          `json:"port"`
	AssetDir     string          `json:"asset_dir"`
	UploadDir    string          `json:"upload_dir"`
	HTTPS        bool            `json:"https"`
	SSLCert      string          `json:"ssl_cert"`
	SSLKey       string          `json:"ssl_key"`
	DBMSRun      bool            `json:"dbms_run"`
	DBMSBootTime int             `json:"dbms_boot_time"`
	SessionLen   int             `json:"session_len"`
}

func Load(location string) (Conf, error) {
	data, err := os.ReadFile(location)
	if err != nil {
		return defaultConf(), err
	}
	conf := defaultConf()
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return defaultConf(), err
	}
	return conf, nil
}

func defaultConf() Conf {
	return Conf{
		DBMSConf: gotables.Config{
			Conf: fs.Conf{
				Port:            ":5678",
				Dir:             "dbms/srv",
				LogDir:          "dbms/log",
				HTTPSMode:       false,
				SSLCert:         "",
				SSLKey:          "",
				EnableGTSyntax:  true,
				EnableSQLSyntax: false,
			},
			Host: "127.0.0.1",
		},
		Port:         ":8080",
		AssetDir:     "assets",
		UploadDir:    "uploads",
		HTTPS:        false,
		SSLCert:      "",
		SSLKey:       "",
		DBMSRun:      false,
		DBMSBootTime: 10,
		SessionLen:   5,
	}
}
