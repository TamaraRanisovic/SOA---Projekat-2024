package config

import "os"

type Config struct {
	UserServiceAddress string
	Address            string
}

func GetConfig() Config {
	return Config{
		UserServiceAddress: os.Getenv("USER_MANAGEMENT_SERVICE_ADDRESS"),
		Address:            os.Getenv("AUTH_SERVICE_ADDRESS"),
	}
}
