package config

import "os"

type Config struct {
	Address string
}

func GetConfig() Config {
	return Config{
		Address: os.Getenv("USER_MANAGEMENT_SERVICE_ADDRESS"),
	}
}
