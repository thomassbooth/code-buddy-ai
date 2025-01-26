package config

type ServerConfig struct {
	Port string
}

type Config struct {
	Server ServerConfig
}

var AppConfig = Config{
	Server: ServerConfig{
		Port: "8065",
	},
}
