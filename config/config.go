package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App           AppAccount
	Connection    ConnectionAccount
	Logger        LoggerAccount
}

type AppAccount struct {
	Name          string
}

type ConnectionAccount struct {
	DatabaseApp DatabaseAccount
}

type DatabaseAccount struct {
	DriverName      string
	DriverSource    string
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
}

type LoggerAccount struct {
	Logrus    LogrusAccount
}

type LogrusAccount struct {
	Level string
}

//=================================================================================================================

// * Init Config
func InitConfig(env string) *Config {
	configPath := GetConfigPath(env)

	confFile, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	conf, err := ParseConfig(confFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	return conf
}

// * Get config path for local or docker
func GetConfigPath(configPath string) string {
	return "./config/config"
}

// * Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// * Parse config ifle
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
