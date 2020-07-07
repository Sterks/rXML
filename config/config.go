package config

// Config - конфигурация
type Config struct {
	Rabbit Rabbit
}

// Rabbit подключение
type Rabbit struct {
	ConnectRabbit string `toml:"connection_rabbit"`
}

// NewConfig инициализация
func NewConfig() *Config {
	return &Config{}
}
