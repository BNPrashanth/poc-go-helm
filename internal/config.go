package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

/*
InitializeViper Function initializes viper to read config.yml file and environment variables in the application.
*/
func InitializeViper() {

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	cfgPath := os.Getenv(
		strings.ToUpper(strings.Replace("poc-helm", "-", "_", -1)) + "_CONFIG_PATH",
	)
	fmt.Println("cfgPath", cfgPath)
	if cfgPath != "" {
		viper.SetConfigFile(cfgPath)
	}

	err := viper.ReadInConfig()
	if err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Error reading config:", err.Error())
	}

}
