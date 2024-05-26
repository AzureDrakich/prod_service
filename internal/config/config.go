package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"false"`
	Listen        struct {
		Type   string `env:"Listen_Type" env-default:"port"`
		BindIp string `env:"Bind_Ip" env-default:"0.0.0.0"`
		Port   string `env:"Port" env-default:"10000"`
	}
	AppConfig struct {
		LogLevel  string
		AdminUser struct {
			Email    string `env:"Admin_Email" env-default:"admin"`
			Password string `env:"Admin_Password" env-default:"admin"`
		}
	}
	PostgreSQL struct {
		Username string `env:"PSQL_Username" env-required:"true"`
		Password string `env:"PSQL_Password"  env-required:"true"`
		Host     string `env:"PSQL_Host"  env-required:"true"`
		Port     string `env:"PSQL_Port"  env-required:"true"`
		Database string `env:"PSQL_Database"  env-required:"true"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("get config")
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "Monolith system"
			description, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(description)
			log.Fatal(err)
		}
	})
	return instance
}
