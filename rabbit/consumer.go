package rabbit

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/Sterks/rXML/config"
	"github.com/Sterks/rXML/mongo"
	"github.com/Sterks/rXML/structs"
	"github.com/Sterks/rXmlReader/rabbit"

	"github.com/streadway/amqp"
)

// ConsumerMQ Подписчик
type ConsumerMQ struct {
	conn    *amqp.Connection
	config  *config.Config
	channel *amqp.Channel
}

// NewConsumer Инициализация
func NewConsumer(config *config.Config) *ConsumerMQ {
	return &ConsumerMQ{
		conn:    nil,
		channel: nil,
		config:  config,
	}
}

func (c *ConsumerMQ) ConsumerMQNow(config *config.Config, nameQueue string) <-chan amqp.Delivery {
	conn, err := amqp.Dial(config.Rabbit.ConnectRabbit)
	if err != nil {
		log.Fatalf("connection.open: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("connection.channel - %s", err)
	}
	//defer ch.Close()

	q, err2 := ch.QueueDeclare(
		nameQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err2 != nil {
		log.Fatalf("connection.QueueDeclare - %s", err2)
	}

	err = ch.Qos(1, 0, false)
	failOnError(err, "Qos - не работает")

	msqs, err := ch.Consume(
		q.Name,
		"",
		false, // автоответ отключен
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("connection.consumer - %s", err)
	}

	m := mongo.MongoDBNew(config)
	m.MongoDBConnect(config)
	stopChan := make(chan bool)
	go func() {

		log.Printf("Consumer ready, PID: %d", os.Getpid())
		var inf rabbit.InformationFile

		for message := range msqs {
			if err := json.Unmarshal(message.Body, &inf); err != nil {
				log.Printf("Не могу прочитать - %v", err)
			}
			pattern := []string{
				".EA44_",
				".EA615_",
				".INM111_",
				".OK504_",
				".ZA44_",
				".ZK504_",
				".ZKK44_",
				".ZP504_",
				".OKU504_",
			}

			for _, k := range pattern {
				matched, err := regexp.MatchString(k, inf.NameFile)
				if err != nil {
					log.Fatal(err)
				}
				// str := fmt.Sprintf("%v", value)
				// fmt.Printf(str, " - ", k)
				if matched {
					switch typeNot := k; typeNot {
					case ".EA44_":
						var ea44 structs.NotificationEA44
						if err := xml.Unmarshal(inf.FileZip, &ea44); err != nil {
							log.Println(err)
						}
						// fmt.Println(ea44.FcsNotificationEF.PurchaseNumber)
						m.Save("readerXML", "xml_notification", &ea44)
					case ".EA615_":
						var ea615 structs.NotificationEA615
						xml.Unmarshal(inf.FileZip, &ea615)
						// fmt.Println(ea615)
						m.Save("readerXML", "xml_notification", &ea615)
					case ".INM111_":
						var inm111 structs.NotificationNM111
						xml.Unmarshal(inf.FileZip, &inm111)
						m.Save("readerXML", "xml_notification", &inm111)
					case ".OK504_":
						var ok504 structs.NotificationOK504
						xml.Unmarshal(inf.FileZip, &ok504)
						// fmt.Println(ok504)
						m.Save("readerXML", "xml_notification", &ok504)
					case ".ZA44_":
						var za44 structs.NotificationZA44
						xml.Unmarshal(inf.FileZip, &za44)
						// fmt.Println(za44)
						m.Save("readerXML", "xml_notification", &za44)
					case ".ZK504_":
						var zk504 structs.NotificationZK504
						xml.Unmarshal(inf.FileZip, &zk504)
						// fmt.Println(zk504)
						m.Save("readerXML", "xml_notification", &zk504)
					case ".ZKK44_":
						var zkk44 structs.NotificationZKKP44
						xml.Unmarshal(inf.FileZip, &zkk44)
						// fmt.Println(zkk44)
						m.Save("readerXML", "xml_notification", &zkk44)
					case ".ZP504_":
						var zp504 structs.NotificationZP504
						xml.Unmarshal(inf.FileZip, &zp504)
						// fmt.Println(zp504)
						m.Save("readerXML", "xml_notification", &zp504)
					case ".OKU504_":
						var oku504 structs.NotificationOKU504
						xml.Unmarshal(inf.FileZip, &oku504)
						// fmt.Println(oku504)
						m.Save("readerXML", "xml_notification", &oku504)
					default:
						fmt.Println("Не могу определить тип извещения", typeNot)
					}
					fmt.Println(inf.NameFile)
				}
			}
			if err := message.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message - %v", inf.NameFile)
			}
		}
	}()
	<-stopChan

	return msqs
}

func (c *ConsumerMQ) Reader() {

}
