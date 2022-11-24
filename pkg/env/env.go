package env

import (
	"github.com/spf13/viper"
	"log"
)

//type Config struct {
//	DBName string `mapstructure:"DB_Name"`
//	DBURL  string `mapstructure:"DB_URL"`
//}

func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}
	return
}
