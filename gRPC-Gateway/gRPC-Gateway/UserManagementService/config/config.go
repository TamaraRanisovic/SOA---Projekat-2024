package config

import "os"

type Config struct {
	AuthServiceAddress string
	Address            string
}

func GetConfig() Config {
	return Config{
		Address:            os.Getenv("USER_MANAGEMENT_SERVICE_ADDRESS"),
		AuthServiceAddress: os.Getenv("AUTH_SERVICE_ADDRESS"),
	}
}
