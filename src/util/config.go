package util

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Env           string   // "development" or "production"
	Port          string   // application listen port
	PublicKeyPath string   // auth server's public key
	Database      DBConfig `toml:"database"`
}

type DBConfig struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
	Debug    bool
}

func LoadConfig(filePath string) Config {
	var config Config
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		panic(err)
	}
	return config
}
