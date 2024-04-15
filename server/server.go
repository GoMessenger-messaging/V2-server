package server

import (
	"git.jereileu.ch/gomessenger/gomessenger/server/conf"
	"log"
	"net/http"
	"strings"
)

func Run(config conf.Conf) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(r.URL.Path, "/")
		log.Println(path)
	})
	if config.HTTPS {
		log.Fatalln(http.ListenAndServeTLS(config.Port, config.SSLCert, config.SSLKey, nil))
	} else {
		log.Fatalln(http.ListenAndServe(config.Port, nil))
	}
}
