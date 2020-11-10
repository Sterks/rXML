package services

import (
	"github.com/Sterks/rXML/config"
	"github.com/Sterks/rXML/logger"
	"github.com/Sterks/rXML/mongo"
	"github.com/Sterks/rXML/rabbit"
	"github.com/streadway/amqp"
)

type Notifications223 struct {
	Logger  *logger.Logger
	Mongo   *mongo.MongoDb
	Config  *config.Config
	Rabbit  *rabbit.ConsumerMQ
	Channel <-chan amqp.Delivery
}

func NewNotifications223() *Notifications223 {
	return &Notifications223{
		Logger: logger.NewLogger(),
		Mongo:  &mongo.MongoDb{},
		Config: &config.Config{},
		Rabbit: &rabbit.ConsumerMQ{},
	}
}

func (n Notifications223) ConfigureXMLNotification(config *config.Config) {
	n.Logger.ConfigureLogger(config)
	n.Mongo.MongoDBConnect(config)
	channel := n.Rabbit.ConsumerMQNow(config, "Notifications223OpenFile")
	n.Channel = channel
}

func (n Notifications223) ParseXML() {
}
