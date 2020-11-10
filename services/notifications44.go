package services

import (
	"github.com/Sterks/rXML/config"
	"github.com/Sterks/rXML/logger"
	"github.com/Sterks/rXML/mongo"
	"github.com/Sterks/rXML/rabbit"
	"github.com/streadway/amqp"
)

// Notifications44 Структура для XML
type Notifications44 struct {
	Logger  *logger.Logger
	Mongo   *mongo.MongoDb
	Config  *config.Config
	Rabbit  *rabbit.ConsumerMQ
	Channel <-chan amqp.Delivery
}

// NewNotification44 ...
func NewNotification44() *Notifications44 {
	return &Notifications44{
		Logger: logger.NewLogger(),
		Mongo:  &mongo.MongoDb{},
		Config: &config.Config{},
		Rabbit: &rabbit.ConsumerMQ{},
	}
}

// ConfigureXMLNotification ...
func (xml *Notifications44) ConfigureXMLNotification(config *config.Config) {
	xml.Logger.ConfigureLogger(config)
	xml.Mongo.MongoDBConnect(config)
	channel := xml.Rabbit.ConsumerMQNow(config, "Notifications44OpenFile")
	xml.Channel = channel
}

// Connect ...
func (xml *Notifications44) Connect(config *config.Config) {
	consumer := rabbit.NewConsumer(config)
	consumer.ConsumerMQNow(config, "Notifications44OpenFile")
	xml.Mongo.MongoDBConnect(config)
}

// ParseXML ...
func (xml *Notifications44) ParseXML() {
}

// Validation ...
func (xml *Notifications44) Validation() {
}
