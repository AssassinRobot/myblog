package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Cfg = Config{}

type Config struct {
	Mysql  MysqlConfig
	Jwt    JwtConfig
	Server ServerConfig
}
type MysqlConfig struct {
	Host     string
	Username string
	Password string
	Port     string
	DBname   string
}
type JwtConfig struct {
	Secret string
}
type ServerConfig struct {
	Port string
}

func InitConfig() {
	cfg := Config{}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %v", err))
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("Error unmarshaling config: %s\n", err)
		return
	}
	Cfg = cfg
}