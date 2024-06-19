package service

import (
	"errors"
	"minimark/config"
	"minimark/utils"
)

func UserLogin(username, password string) (string, error) {
	cfg := config.GetConfig()
	users := cfg.Server.Users
	for _, user := range users {
		if user.Username == username && user.Password == password {
			jwt, err := utils.CreateJwtToken(map[string]any{
				"username": username,
			})
			if err != nil {
				return "", err
			}
			return jwt, nil
		}
	}
	return "", errors.New("用户名或密码错误")
}
