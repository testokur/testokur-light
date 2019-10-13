package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	port = "PORT"
)

func init() {
	viper.SetDefault(port, 8066)
	viper.AutomaticEnv()
}

// GetPort returns Application port
func GetPort() string {
	return fmt.Sprintf(":%d", viper.GetInt(port))
}
