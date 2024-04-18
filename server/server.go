package server

import (
	"git.jereileu.ch/gomessenger/gomessenger/gm-server/api"
	"git.jereileu.ch/gomessenger/gomessenger/gm-server/assets"
	"git.jereileu.ch/gomessenger/gomessenger/gm-server/conf"
	"io"
	"log"
	"net/http"
	"strings"
)

func Run(config conf.Conf) {
	http.HandleFunc("/uploads/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		path = path[1:]

		data, contentType, err := assets.GetAsset(connectPath(path), config.UploadDir)
		if err != nil {
			w.WriteHeader(403)
		} else {
			w.Header().Set("Content-Type", contentType)
			_, err := w.Write(data)
			if err != nil {
				w.WriteHeader(500)
			}
		}
	})
	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		path = path[1:]

		data, contentType, err := assets.GetAsset(connectPath(path), config.AssetDir)
		if err != nil {
			w.WriteHeader(403)
		} else {
			w.Header().Set("Content-Type", contentType)
			_, err := w.Write(data)
			if err != nil {
				w.WriteHeader(500)
			}
		}
	})
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		path = path[1:]

		if len(path) == 0 {
			w.WriteHeader(404)
		} else {
			data, statusCode := api.Api(readBody(r), path, config)
			w.WriteHeader(statusCode)
			_, err := w.Write(data)
			if err != nil {
				w.WriteHeader(500)
			}
		}
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

func readBody(r *http.Request) []byte {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil
	}
	return data
}
