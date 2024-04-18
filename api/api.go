package api

import (
	"encoding/json"
	"git.jereileu.ch/gomessenger/gomessenger/gm-server/conf"
	"git.jereileu.ch/gomessenger/gomessenger/gm-server/user"
)

func Api(data []byte, path []string, config conf.Conf) ([]byte, int) {
	if len(path) < 3 {
		return nil, 404
	}

	switch path[0] {
	case "user":
		switch path[1] {
		case "auth":
			switch path[2] {
			case "login":
				return userAuthLogin(data, config)
			case "signup":
				return userAuthSignup(data, config)
			default:
				return nil, 404
			}
		default:
			return nil, 404
		}
	case "channel":
		return nil, 418
	case "community":
		return nil, 418
	default:
		return nil, 404
	}
}

func userAuthLogin(data []byte, config conf.Conf) ([]byte, int) {
	dataStruct := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err := json.Unmarshal(data, &dataStruct)
	if err != nil {
		return nil, 400
	}
	id, err := user.ValidateCredentials(dataStruct.Username, dataStruct.Password, config)
	if err != nil {
		return nil, 403
	}
	retStruct := struct {
		Sessionid string `json:"sessionid"`
	}{
		Sessionid: id,
	}
	data, err = json.Marshal(retStruct)
	if err != nil {
		return nil, 500
	}
	return data, 200
}

func userAuthSignup(data []byte, config conf.Conf) ([]byte, int) {
	return nil, 418
}
