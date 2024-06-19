package config

import (
	"log"

	"github.com/spf13/viper"
)

func ConfigureViper() {
	println("[+] Processing: Initializing Viper Configuration [+]")

	// Specify the path to the configuration file
	viper.SetConfigFile("./internal/config/config.json")

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		println("[-] Error: Failed to read config file [-]")
		log.Fatal(err.Error())
	}
}
