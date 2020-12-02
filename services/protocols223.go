package services

import (
	"github.com/Sterks/rXML/config"
	"github.com/Sterks/rXML/logger"
	"github.com/Sterks/rXML/mongo"
	"github.com/Sterks/rXML/rabbit"
	"github.com/streadway/amqp"
)

type Protocols223 struct {
	Logger  *logger.Logger
	Mongo   *mongo.MongoDb
	Config  *config.Config
	Rabbit  *rabbit.ConsumerMQ
	Channel <-chan amqp.Delivery
}

func NewProtocols223() *Protocols223 {
	return &Protocols223{
		Logger: logger.NewLogger(),
		Mongo:  &mongo.MongoDb{},
		Config: &config.Config{},
		Rabbit: &rabbit.ConsumerMQ{},
	}
}

func (p Protocols223) ConfigureXMLNotification(config *config.Config) {
	p.Logger.ConfigureLogger(config)
	p.Mongo.MongoDBConnect(config)
	channel := p.Rabbit.ConsumerMQNow(config, "Protocols223OpenFile")
	p.Channel = channel
}

func (n Protocols223) ParseXML() {
}
