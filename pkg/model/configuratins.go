package model

// Configuration for application
type Configuration struct {
	Environment string
	RestPort    int
	GinRunMode  string
}
