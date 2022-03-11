package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./secrets")
	viper.AddConfigPath("../secrets")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("OK -", viper.Get("demo.key"))
}
