package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type config struct {
	Database struct {
		Dbms   string
		User   string
		Pass   string
		Host   string
		Port   string
		DbName string
		Option string
	}
}

var c config

func Init() {
	viper.AddConfigPath("./backend/config/develop")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config file read error")
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&c); err != nil {
		fmt.Println("config file Unmarshal error")
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewConfig() *config {
	return &c
}
