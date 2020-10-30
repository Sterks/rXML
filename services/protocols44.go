package services

import (
	"github.com/Sterks/rXML/config"
	"github.com/Sterks/rXML/logger"
	"github.com/Sterks/rXML/mongo"
	"github.com/Sterks/rXML/rabbit"
	"github.com/streadway/amqp"
)

type Protocol44 struct {
	Logger  *logger.Logger
	Mongo   *mongo.MongoDb
	Config  *config.Config
	Rabbit  *rabbit.ConsumerMQ
	Channel <-chan amqp.Delivery
}

func NewProtocol44() *Protocol44 {
	return &Protocol44{
		Logger: logger.NewLogger(),
		Mongo:  &mongo.MongoDb{},
		Config: &config.Config{},
		Rabbit: &rabbit.ConsumerMQ{},
	}
}

func (p Protocol44) ConfigureXMLNotification(config *config.Config) {
	p.Logger.ConfigureLogger(config)
	p.Mongo.MongoDBConnect(config)
	channel := p.Rabbit.ConsumerMQNow(config, "Protocols44OpenFile")
	p.Channel = channel
}

func (p Protocol44) ParseXML() {

}
