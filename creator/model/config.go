package model

type Config struct  {
	ServerConfig ServerConfig
}

type ServerConfig struct  {
	Port string
}

func (o *Config) SaveDefaultConfigParams() {
	if o.ServerConfig.Port == "" {
		o.ServerConfig.Port = ":8080"
	}
}
