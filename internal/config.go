package internal

import (
	"github.com/spf13/viper"
)

/*
InitializeViper Function initializes viper to read config.yml file and environment variables in the application.
*/
func InitializeViper() {

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

}
