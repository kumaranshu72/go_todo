package main

import (
	"go_todo/config"
)

const (
	configPath = "config/reader"
)

func main() {
	config.GetConfig(configPath)

}
