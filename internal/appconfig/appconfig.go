package appconfig

import (
	"log"

	"github.com/amrahman90/go-CRUD-api-sample/pkg/config"
	"github.com/spf13/viper"
)

func Init() *config.Config {
	var config config.Config
	{
		viperConfig := viper.New()
		viperConfig.SetConfigType("yml")

		viperConfig.SetConfigName("config.yml")
		viperConfig.AddConfigPath("./configs")
		err := viperConfig.ReadInConfig()

		if err != nil {
			log.Fatal(err)
		}
		err = viperConfig.Unmarshal(&config)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &config
}
