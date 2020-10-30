package rabbit

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Sterks/rXML/protocols44"
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
	msqs := c.GetXML(config, nameQueue)

	m := mongo.MongoDBNew(config)
	m.MongoDBConnect(config)
	stopChan := make(chan bool)
	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		var inf rabbit.InformationFile
		if nameQueue == "Notifications44OpenFile" {
			c.ChooseTemlateNotification44(msqs, inf, m)
		} else if nameQueue == "Protocols44OpenFile" {
			c.ChooseTemlateProtocol44(msqs, inf, m)
		}
	}()
	<-stopChan

	return msqs
}

func (c *ConsumerMQ) ChooseTemlateNotification44(msqs <-chan amqp.Delivery, inf rabbit.InformationFile, m *mongo.MongoDb) {
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
}

func (c *ConsumerMQ) ChooseTemlateProtocol44(msqs <-chan amqp.Delivery, inf rabbit.InformationFile, m *mongo.MongoDb) {
	for message := range msqs {
		if err := json.Unmarshal(message.Body, &inf); err != nil {
			log.Printf("Не могу прочитать - %v", err)
		}
		pattern := []string{
			"EF1",
			"EF2",
			"EF3",
			"EFInv",
			"EFSingleApp",
			"EOK1",
			"EOK2",
			"EOK3",
			"EOKOU1",
			"EOKOU2",
			"EZK1",
			"EZK2",
			"PP615",
			"PR615",
			"PZP",
			"VPP",
		}

		for _, k := range pattern {
			matched, err := regexp.MatchString(k, inf.NameFile)
			if err != nil {
				log.Fatal(err)
			}

			collection := "xml_protocols44"

			/// TODO Сделать одной функцией
			// str := fmt.Sprintf("%v", value)
			// fmt.Printf(str, " - ", k)
			if matched {
				switch typeNot := k; typeNot {
				case "EF1":
					var ef1 protocols44.ProtocolEF1
					if err := xml.Unmarshal(inf.FileZip, &ef1); err != nil {
						log.Println(err)
					}
					// fmt.Println(ea44.FcsNotificationEF.PurchaseNumber)
					m.Save("readerXML", collection, &ef1)
				case "EF2":
					var ief2 protocols44.ProtocolEF2
					xml.Unmarshal(inf.FileZip, &ief2)
					// fmt.Println(ea615)
					m.Save("readerXML", collection, &ief2)
				case "EF3":
					var ief3 protocols44.ProtocolEF3
					xml.Unmarshal(inf.FileZip, &ief3)
					// fmt.Println(ea615)
					m.Save("readerXML", collection, &ief3)
				case "EFInv":
					var iefinv protocols44.ProtocolEF3
					xml.Unmarshal(inf.FileZip, &iefinv)
					m.Save("readerXML", collection, &iefinv)
				case "EFSingleApp":
					var iefSingleApp protocols44.ProtocolEFSingle
					xml.Unmarshal(inf.FileZip, &iefSingleApp)
					// fmt.Println(ok504)
					m.Save("readerXML", collection, &iefSingleApp)
				case "EOK1":
					var ieok1 protocols44.ProtocolEOK1
					xml.Unmarshal(inf.FileZip, &ieok1)
					// fmt.Println(za44)
					m.Save("readerXML", collection, &ieok1)
				case "EOK2":
					var ieok2 protocols44.ProtocolEOK2
					xml.Unmarshal(inf.FileZip, &ieok2)
					// fmt.Println(zk504)
					m.Save("readerXML", collection, &ieok2)
				case "EOK3":
					var ieok3 protocols44.ProtocolEOK3
					xml.Unmarshal(inf.FileZip, &ieok3)
					// fmt.Println(zkk44)
					m.Save("readerXML", collection, &ieok3)
				case "EOKOU1":
					var ieokou1 protocols44.ProtocolEOKOU1
					xml.Unmarshal(inf.FileZip, &ieokou1)
					// fmt.Println(zp504)
					m.Save("readerXML", collection, &ieokou1)
				case "EOKOU2":
					var ieokou2 protocols44.ProtocolEOKOU2
					xml.Unmarshal(inf.FileZip, &ieokou2)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &ieokou2)
				case "EZK1":
					var iezk1 protocols44.ProtocolEZK1
					xml.Unmarshal(inf.FileZip, &iezk1)
					// fmt.Println(zp504)
					m.Save("readerXML", collection, &iezk1)
				case "EZK2":
					var iezk2 protocols44.ProtocolEZK2
					xml.Unmarshal(inf.FileZip, &iezk2)
					// fmt.Println(zp504)
					m.Save("readerXML", collection, &iezk2)
				case "PP615":
					var ipp615 protocols44.ProtocolPP615
					xml.Unmarshal(inf.FileZip, &ipp615)
					// fmt.Println(zp504)
					m.Save("readerXML", collection, &ipp615)
				case "PR615":
					var ipr615 protocols44.ProtocolPR615
					xml.Unmarshal(inf.FileZip, &ipr615)
					// fmt.Println(zp504)
					m.Save("readerXML", collection, &ipr615)
				case "PZP":
					var ipzp protocols44.ProtocolPZP
					xml.Unmarshal(inf.FileZip, &ipzp)
					// fmt.Println(zp504)
					m.Save("readerXML", collection, &ipzp)
				case "VPP":
					var ivpp protocols44.ProtocolVPP
					xml.Unmarshal(inf.FileZip, &ivpp)
					// fmt.Println(zp504)
					m.Save("readerXML", collection, &ivpp)
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
}

func (c *ConsumerMQ) GetXML(config *config.Config, nameQueue string) <-chan amqp.Delivery {
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
	return msqs
}

func (c *ConsumerMQ) Reader() {

}
