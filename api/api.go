package api

import (
	"git.jereileu.ch/gomessenger/gomessenger/gm-server/conf"
)

func Api(path []string, config conf.Conf) ([]byte, int) {
	if len(path) < 3 {
		return nil, 404
	}

	switch path[0] {
	case "user":
		switch path[1] {
		case "auth":
			switch path[2] {
			case "login":
				return nil, 418
			case "signup":
				return nil, 418
			default:
				return nil, 404
			}
		default:
			return nil, 404
		}
	case "chat":
		return nil, 418
	case "group":
		return nil, 418
	default:
		return nil, 404
	}
}
