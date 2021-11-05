package pkg

import (
	"flag"
)

type Config struct {
	MeowerServiceAddress string
}

func ProvideConfig() *Config {
	config := &Config{
		MeowerServiceAddress: *flag.String("meowerServiceAddress", "127.0.0.1:8080", "Meower service address"),
	}

	flag.Parse()

	return config
}
