package config

import "github.com/Netflix/go-env"

type Config struct {
	RestPort int `env:"REST_PORT,default=8080"`
	GRPCPort int `env:"GRPC_PORT,default=8081"`
}

func ReadConfigFromEnv() (cfg Config, err error) {
	conf = Config{}
	_, err = env.UnmarshalFromEnviron(&conf)
	return
}
