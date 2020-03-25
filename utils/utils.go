package utils

import (
	"os"
	"strconv"
)

// GetEnv : get env variables
func GetEnv(key, def string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		value = def
	}
	return
}

// GetEnvInt : get int env
func GetEnvInt(key string, def int) (value int, err error) {
	if tmp := os.Getenv(key); tmp != "" {
		value, err = strconv.Atoi(tmp)
	} else {
		value = def
	}
	return
}
