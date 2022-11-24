package env

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig("../../")

	fmt.Println(viper.Get("DB_NAME"))

}
