package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name string
	} 
	Redis struct{
		Host string
		Port string
		Password string
	}
}

var AppConfig *Config

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("..")

	if err:=viper.ReadInConfig(); err!=nil{
		log.Fatalf("Read config error : %v",err)
	}

	AppConfig =&Config{}
	if err := viper.Unmarshal(AppConfig);err!=nil{
		log.Fatalf("Encoding config error: %v",err)
	}

	InitDb()
	InitRedis()
	 
}