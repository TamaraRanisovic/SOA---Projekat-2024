package config

import "os"

type Config struct {
	Address            string
	TourServiceAddress string
	BlogServiceAddress string
	AuthServiceAddress string
	UserServiceAddress string
}

func GetConfig() Config {
	return Config{
		TourServiceAddress: os.Getenv("TOUR_SERVICE_ADDRESS"),
		BlogServiceAddress: os.Getenv("BLOG_SERVICE_ADDRESS"),
		AuthServiceAddress: os.Getenv("AUTH_SERVICE_ADDRESS"),
		UserServiceAddress: os.Getenv("USER_MANAGEMENT_SERVICE_ADDRESS"),
		Address:            os.Getenv("GATEWAY_ADDRESS"),
	}
}
