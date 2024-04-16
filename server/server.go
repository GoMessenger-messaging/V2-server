package server

import (
	"git.jereileu.ch/gomessenger/gomessenger/gm-server/assets"
	"git.jereileu.ch/gomessenger/gomessenger/gm-server/conf"
	"log"
	"net/http"
	"strings"
)

func Run(config conf.Conf) {
	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(r.URL.Path, "/")
		path = path[2:]

		data, contentType, err := assets.GetAsset(connectPath(path), config.AssetDir)
		if err != nil {
			w.WriteHeader(404)
		} else {
			w.Header().Set("Content-Type", contentType)
			_, err := w.Write(data)
			if err != nil {
				w.WriteHeader(500)
			}
		}
	})
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(r.URL.Path, "/")
		path = path[2:]

		w.WriteHeader(418)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})
	if config.HTTPS {
		log.Fatalln(http.ListenAndServeTLS(config.Port, config.SSLCert, config.SSLKey, nil))
	} else {
		log.Fatalln(http.ListenAndServe(config.Port, nil))
	}
}

func connectPath(path []string) string {
	var out string
	for i := 0; i < len(path); i++ {
		out += path[i]
		if i != len(path)-1 {
			out += "/"
		}
	}
	return out
}
