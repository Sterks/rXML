package services

import (
	"github.com/Sterks/rXML/config"
	"github.com/Sterks/rXML/logger"
	"github.com/Sterks/rXML/mongo"
	"github.com/Sterks/rXML/rabbit"
	"github.com/streadway/amqp"
)

// XMLNotification Структура для XML
type XMLNotification struct {
	Logger  *logger.Logger
	Mongo   *mongo.MongoDb
	Config  *config.Config
	Rabbit  *rabbit.ConsumerMQ
	Channel <-chan amqp.Delivery
}

// NewXMLNotification ...
func NewXMLNotification() *XMLNotification {
	return &XMLNotification{
		Logger: logger.NewLogger(),
		Mongo:  &mongo.MongoDb{},
		Config: &config.Config{},
		Rabbit: &rabbit.ConsumerMQ{},
	}
}

// ConfigureXMLNotification ...
func (xml *XMLNotification) ConfigureXMLNotification(config *config.Config) {
	xml.Logger.ConfigureLogger(config)
	xml.Mongo.MongoDBConnect(config)
	channel := xml.Rabbit.ConsumerMQNow(config, "Notifications44OpenFile")
	xml.Channel = channel
}

// Connect ...
func (xml *XMLNotification) Connect(config *config.Config) {
	consumer := rabbit.NewConsumer(config)
	consumer.ConsumerMQNow(config, "Notifications44OpenFile")
	xml.Mongo.MongoDBConnect(config)
}

// ParseXML ...
func (xml *XMLNotification) ParseXML() {
	// var inf rabbit.InformationFile
	// go func() {
	// 	for message := range xml.Channel {

	// 		if err := json.Unmarshal(message.Body, &inf); err != nil {
	// 			xml.Logger.ErrorLog("Не могу прочитать - %v", err)
	// 		}

	// 		pattern := map[string]interface{}{
	// 			".EA44_":   structs.NotificationEA44{},
	// 			".EA615_":  structs.NotificationEA615{},
	// 			".INM111_": structs.NotificationNM111{},
	// 			".OK504_":  structs.NotificationOK504{},
	// 			".ZA44_":   structs.NotificationZA44{},
	// 			".ZK504_":  structs.NotificationZK504{},
	// 			".ZKK44_":  structs.NotificationZKKP44{},
	// 			".ZP504_":  structs.NotoficationZP504{},
	// 			".OKU504_": structs.NotificationOKU504{}}
	// 		for k, t := range pattern {
	// 			matched, err := regexp.MatchString(k, inf.NameFile)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	// 			if matched {
	// 				xml.Mongo.Save("readerXML", "xml_notification", t)
	// 				if err := message.Ack(false); err != nil {
	// 					log.Printf("Error acknowledging message : %s", err)
	// 				} else {
	// 					log.Printf("Acknowledged message")
	// 				}
	// 			}
	// 		}
	// 	}
	// }()
}

// Validation ...
func (xml *XMLNotification) Validation() {

}
