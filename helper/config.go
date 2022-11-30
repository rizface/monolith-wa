package helper

import (
	"log"

	"github.com/spf13/viper"
)

func ConfigLoader(path, filename, configtype string) {
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType(configtype)
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}
}
