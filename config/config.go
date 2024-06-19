package config

import (
	"github.com/pelletier/go-toml/v2"
	"os"
	"sync"
)

type Config struct {
	Database DatabaseConfig `toml:"database"`
	Server   ServerConfig   `toml:"server"`
	Security SecurityConfig `toml:"security"`
}

type DatabaseConfig struct {
	Driver string `toml:"driver"`
	DSN    string `toml:"dsn"`
}

type ServerConfig struct {
	IP    string       `toml:"ip"`
	Port  uint16       `toml:"port"`
	Users []UserConfig `toml:"users"`
}

type UserConfig struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type SecurityConfig struct {
	JwtKey    string `toml:"jwt_key"`
	JwtExpire int64  `toml:"jwt_expire"`
}

const (
	configFileName = "./config.toml"
)

var (
	cfg     *Config
	onceCfg sync.Once
)

// GetConfig 获取配置文件的配置
func GetConfig() *Config {
	onceCfg.Do(func() {
		data, err := os.ReadFile(configFileName)
		if err != nil {
			panic(err)
		}

		cfg = new(Config)
		err = toml.Unmarshal(data, cfg)
		if err != nil {
			panic(err)
		}
	})
	return cfg
}
