package server

import (
	"git.jereileu.ch/gomessenger/gomessenger/server/assets"
	"git.jereileu.ch/gomessenger/gomessenger/server/conf"
	"log"
	"net/http"
	"strings"
)

func Run(config conf.Conf) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(r.URL.Path, "/")
		path = path[2:]

		w.WriteHeader(418)
	})
	http.HandleFunc("/assets", func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(r.URL.Path, "/")
		path = path[2:]

		assets.GetAsset(config)
		w.WriteHeader(418)
	})
	if config.HTTPS {
		log.Fatalln(http.ListenAndServeTLS(config.Port, config.SSLCert, config.SSLKey, nil))
	} else {
		log.Fatalln(http.ListenAndServe(config.Port, nil))
	}
}
