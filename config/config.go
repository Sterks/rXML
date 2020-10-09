package config

// Config - конфигурация
type Config struct {
	Rabbit       Rabbit
	MongoDb      MongoDb
	MainSettings MainSettings
}

// Rabbit подключение
type Rabbit struct {
	ConnectRabbit string `toml:"connection_rabbit"`
}

// MainSettings ...
type MainSettings struct {
	LogLevel string `toml:"log_level"`
}

// MongoDb ...
type MongoDb struct {
	ConnectMongoDB string `toml:"connect_mongo"`
}

// NewConfig инициализация
func NewConfig() *Config {
	return &Config{
		Rabbit:       Rabbit{},
		MongoDb:      MongoDb{},
		MainSettings: MainSettings{},
	}
}
