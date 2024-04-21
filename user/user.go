package user

import (
	"crypto/sha512"
	"errors"
	"git.jereileu.ch/gomessenger/gomessenger/gm-server/auth"
	"git.jereileu.ch/gomessenger/gomessenger/gm-server/conf"
	"git.jereileu.ch/gotables/client/go/gotables"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"math/rand"
	"time"
)

type User struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Status   string `mapstructure:"status"`
	Picture  string `mapstructure:"picture"`
}

func (u *User) FromRow(row map[string]any) error {
	var user User
	err := mapstructure.Decode(row, &user)
	if err != nil {
		return errors.New("failed to convert row to User")
	}
	*u = user
	return nil
}

func (u *User) ToRow() map[string]any {
	row := structs.Map(*u)
	return row
}

func ValidateCredentials(username string, password string, config conf.Conf) (string, error) {
	user, err := gotables.ShowTableConditions(
		[]string{"username == " + username, "password == " + auth.Sum(password)},
		[]string{"*"},
		"users",
		"gomessenger",
		"not implemented yet",
		config.DBMSConf)
	if err != nil {
		return "", err
	}
	if len(user.GetRows()) != 1 {
		return "", errors.New("wrong credentials")
	}
	id, sum := generateSessionid()
	err = storeSessionid(username, sum, config)
	return id, err
}

func GetUser(username string, sessionid string, config conf.Conf) (map[string]any, error) {
	if !CheckSessionid(username, sessionid, config) {
		return nil, errors.New("invalid sessionid")
	}
	user, err := gotables.ShowTableConditions(
		[]string{"username == " + username},
		[]string{"*"},
		"users",
		"gomessenger",
		"not implemented yet",
		config.DBMSConf)
	if err != nil {
		return nil, err
	}
	if len(user.GetRows()) != 1 {
		return nil, errors.New("invalid username")
	}
	return user.GetRows()[0], nil
}

func CheckSessionid(username string, id string, config conf.Conf) bool {
	sessions, err := getSessionids(username, config)
	if err != nil {
		return false
	}
	return validateSessionid(id, sessions)
}

func getSessionids(username string, config conf.Conf) ([][]string, error) {
	conditions := []string{
		"username == " + username,
	}
	columns := []string{
		"sum",
		"expiry",
	}
	sessionTable, err := gotables.ShowTableConditions(conditions, columns, "sessions", "gomessenger", "not implemented yet", config.DBMSConf)
	if err != nil {
		return nil, err
	}
	sessions := sessionTable.GetRows()
	ret := make([][]string, 0)
	for i := 0; i < len(sessions); i++ {
		ret = append(ret, []string{sessions[i]["sum"].(string), sessions[i]["expiry"].(string)})
	}
	return ret, nil
}

func generateSessionid() (string, string) {
	chars := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	idSlice := make([]rune, 64)
	for i := range idSlice {
		idSlice[i] = chars[rand.Intn(len(chars))]
	}
	id := string(idSlice)
	return id, auth.Sum(id)
}

func validateSessionid(id string, sums [][]string) bool {
	sumBytes := sha512.Sum512([]byte(id))
	sum := string(sumBytes[:])
	for i := 0; i < len(sums); i++ {
		if sum == sums[i][0] {
			t, err := time.Parse(time.RFC3339, sums[i][1])
			if err != nil {
				continue
			}
			if time.Now().Before(t) {
				return true
			}
		}
	}
	return false
}

func storeSessionid(username string, sum string, config conf.Conf) error {
	values := [][2]string{
		{"username", username},
		{"sum", sum},
		{"expiry", time.Now().Add(time.Duration(config.SessionLen) * time.Hour).String()},
	}
	_, err := gotables.CreateRow(values, "sessions", "gomessenger", "not implemented yet", config.DBMSConf)
	return err
}
