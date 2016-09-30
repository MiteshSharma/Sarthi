package model

type Config struct  {
	ServerConfig ServerConfig
	DatabaseConfig DatabaseConfig
}

type ServerConfig struct  {
	Port string
}

type DatabaseConfig struct  {
	Host string
	DbName string
}

func (o *Config) SaveDefaultConfigParams() {
	if o.ServerConfig.Port == "" {
		o.ServerConfig.Port = ":8080"
	}
	if o.DatabaseConfig.Host == "" {
		o.DatabaseConfig.Host = "localhost"
	}
	if o.DatabaseConfig.DbName == "" {
		o.DatabaseConfig.DbName = "sarthi"
	}
}
