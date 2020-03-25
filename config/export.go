package config

import (
	"errors"
	"go_todo/config/reader"
	"go_todo/pkg/model"
	"go_todo/utils"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

const (
	envLabel     = "ENVIRONMENT"
	envDevLabel  = "DEVELOPMENT"
	envProdLabel = "PRODUCTION"

	envServerPortLabel   = "SERVER_PORT"
	envMYSQLHost         = "MYSQL_HOST"
	envMYSQLHostPort     = "MYSQL_HOST_PORT"
	envMYSQLHostUser     = "MYSQL_HOST_USER"
	envMYSQLHostPassword = "MYSQL_HOST_PASSWORD"
	envGinRunModeLabel   = "GIN_RUN_MODE"

	serverPortLabel        = "ServerPort"
	mySQLHostLabel         = "MYSQLHost"
	mySQLHostPortLabel     = "MYSQLHostPort"
	mySQLHostUserLabel     = "MYSQLHostUser"
	mySQLHostPasswordLabel = "MYSQLHostPassword"
	ginRunModelLabel       = "GinRunMode"
)

var (
	//Config : instance for exporting application configuration
	Config *model.Configuration
	once   sync.Once
)

// GetConfig will return viper configuration
func GetConfig(configPath string) *model.Configuration {
	once.Do(func() {
		Config = setConfig(reader.GetConfig(configPath))
		// err := allConfigMembersHaveValue(Config)
		// if err != nil {
		// 	log.WithFields(log.Fields{"message": "error on parsing configuration file", "error": err}).
		// 		Fatal("config for field found")
		// }
	})
	return Config
}

// allConfigMembersHaveValue checks the given config have all necessary config fields
// func allConfigMembersHaveValue(config *model.Configuration) error {
// 	if config.GinRunMode == "" || config.MYSQLHost == "" ||
// 		config.Environment == "" {
// 		return errors.New("necessary fields not found")
// 	}
// 	// if config.MYSQLPassword == "" && config.MYSQLUser != "" {
// 	// 	return errors.New("necessary fields not found")
// 	// }
// 	if &config.MYSQLPort == nil || &config.RestPort == nil {
// 		return errors.New("necessary fields not found")
// 	}
// 	return nil
// }

// keyExistsInConfig will check given configuration header isn't empty.
// configuration header can be DEVELOPMENT, TESTING or RELEASE
func keyExistsInConfig(key string, m map[string]interface{}) error {
	err := errors.New("key not found")
	if len(m) == 0 {
		return err
	}
	return nil
}

// set configuration to model.Configuration
func setConfig(cfg *viper.Viper) *model.Configuration {
	c := new(model.Configuration)
	var value interface{}
	env := utils.GetEnv(envLabel, envDevLabel)
	err := errors.New("required field not found")
	var intErr error

	if err := keyExistsInConfig(env, cfg.GetStringMap(env)); err != nil {
		log.WithFields(log.Fields{"message": "error on parsing configuration file", "error": err}).
			Fatal("config for ", env, " not found")
	}
	c.Environment = env

	value = cfg.GetInt(env + "." + serverPortLabel)
	if c.RestPort, intErr = utils.GetEnvInt(envServerPortLabel, value.(int)); intErr != nil {
		log.WithFields(log.Fields{"message": "error on parsing configuration file", "error": intErr}).
			Fatal(envServerPortLabel, " config for field found")
	}
	// value = cfg.GetString(env + "." + mySQLHostLabel)
	// if c.MYSQLHost = utils.GetEnv(envMYSQLHost, value.(string)); c.MYSQLHost == "" {
	// 	log.WithFields(log.Fields{"message": "error on parsing configuration file", "error": err}).
	// 		Fatal(envMYSQLHost, " config for field found")
	// }
	// value = cfg.GetInt(env + "." + mySQLHostPortLabel)
	// if c.MYSQLPort, intErr = utils.GetEnvInt(envMYSQLHostPort, value.(int)); intErr != nil {
	// 	log.WithFields(log.Fields{"message": "error on parsing configuration file", "error": intErr}).
	// 		Fatal(envMYSQLHostPort, " config for field found")
	// }
	// value = cfg.GetString(env + "." + mySQLHostUserLabel)

	// c.MYSQLUser = utils.GetEnv(envMYSQLHostUser, value.(string))

	// value = cfg.GetString(env + "." + mySQLHostPasswordLabel)
	// c.MYSQLPassword = utils.GetEnv(envMYSQLHostPassword, value.(string))
	// if c.MYSQLPassword = utils.GetEnv(envMYSQLHostPassword, value.(string)); c.MYSQLPassword == "" && c.MYSQLUser != "" {
	// 	log.WithFields(log.Fields{"message": "error on parsing configuration file", "error": err}).
	// 		Fatal(envMYSQLHostPassword, " config for field found")
	// }
	value = cfg.GetString(env + "." + ginRunModelLabel)
	if c.GinRunMode = utils.GetEnv(envGinRunModeLabel, value.(string)); c.GinRunMode == "" {
		log.WithFields(log.Fields{"message": "error on parsing configuration file", "error": err}).
			Fatal(envGinRunModeLabel, " config for field found")
	}
	return c
}
