package config

import (
	"errors"
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
	Users             []UserConfig      `toml:"users"`
	HttpRedirectHttps bool              `toml:"http_redirect_https"`
	Http              ServerHttpConfig  `toml:"http"`
	Https             ServerHttpsConfig `toml:"https"`
}

type ServerHttpConfig struct {
	On   bool   `toml:"on"`
	IP   string `toml:"ip"`
	Port uint16 `toml:"port"`
}

type ServerHttpsConfig struct {
	On                 bool   `toml:"on"`
	IP                 string `toml:"ip"`
	Port               uint16 `toml:"port"`
	CertificateFile    string `toml:"certificate_file"`
	CertificateKeyFile string `toml:"certificate_key_file"`
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

		// 检查配置的正确性
		// 检查http和https
		httpOn := cfg.Server.Http.On
		httpsOn := cfg.Server.Https.On
		// 同时关闭
		if !httpOn && !httpsOn {
			panic(errors.New("http和https不能同时关闭"))
		}
		httpPort := cfg.Server.Http.Port
		httpsPort := cfg.Server.Https.Port
		// 端口相同
		if httpOn && httpsOn && httpPort == httpsPort {
			panic(errors.New("http和https不能监听同一个端口"))
		}
		httpRedirectHttps := cfg.Server.HttpRedirectHttps
		// 开启重定向必须同时启用
		if httpRedirectHttps && !(httpOn && httpsOn) {
			panic(errors.New("开启http重定向https，必须同时启用http和https"))
		}
	})
	return cfg
}
