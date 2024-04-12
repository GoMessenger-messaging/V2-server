package server

import (
	"git.jereileu.ch/gomessenger/gomessenger/server/conf"
	"log"
	"net/http"
)

func Run(config conf.Conf) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	if config.HTTPS {
		log.Fatalln(http.ListenAndServeTLS(":"+config.Port, config.SSLCert, config.SSLKey, nil))
	} else {
		log.Fatalln(http.ListenAndServe(":"+config.Port, nil))
	}
}
