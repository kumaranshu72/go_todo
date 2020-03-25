package reader

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

const (
	configType = "toml"
	configName = "config" // By default viper looks for config.toml file
)

var (
	config *viper.viper
	once   sync.Once
)

// GetConfig is for fetching viper configs
func GetConfig(configPath string) *viper.Viper {
	once.Do(func() {
		config = viper.New()
		config.SetConfigName(configName)
		config.SetConfigType(configType)
		config.AddConfigPath(configPath)
		err := config.ReadInConfig()
		if err != nil {
			log.WithFields(log.Fields{"message": "error on parsing config file", "error": err}).
				Fatal("config error")
		}
	})

	return config
}
