package rabbit

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Sterks/rXML/models"
	"github.com/Sterks/rXML/notifications223"
	"github.com/Sterks/rXML/protocols223"
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
		} else if nameQueue == "Notifications223OpenFile" {
			c.ChooseTemlateNotification223(msqs, inf, m)
		} else if nameQueue == "Protocols223OpenFile" {
			c.ChooseTemlateProtocols223(msqs, inf, m)
		}
	}()
	<-stopChan

	return msqs
}

func (c *ConsumerMQ) ChooseTemlateNotification44(msqs <-chan amqp.Delivery, inf rabbit.InformationFile, m *mongo.MongoDb) {
	collection := "xml_notification44"
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
					m.Save("readerXML", collection, &ea44)
				case ".EA615_":
					var ea615 structs.NotificationEA615
					xml.Unmarshal(inf.FileZip, &ea615)
					// fmt.Println(ea615)
					m.Save("readerXML", collection, &ea615)
				case ".INM111_":
					var inm111 structs.NotificationNM111
					xml.Unmarshal(inf.FileZip, &inm111)
					m.Save("readerXML", collection, &inm111)
				case ".OK504_":
					var ok504 structs.NotificationOK504
					xml.Unmarshal(inf.FileZip, &ok504)
					// fmt.Println(ok504)
					m.Save("readerXML", collection, &ok504)
				case ".ZA44_":
					var za44 structs.NotificationZA44
					xml.Unmarshal(inf.FileZip, &za44)
					// fmt.Println(za44)
					m.Save("readerXML", collection, &za44)
				case ".ZK504_":
					var zk504 structs.NotificationZK504
					xml.Unmarshal(inf.FileZip, &zk504)
					// fmt.Println(zk504)
					m.Save("readerXML", collection, &zk504)
				case ".ZKK44_":
					var zkk44 structs.NotificationZKKP44
					xml.Unmarshal(inf.FileZip, &zkk44)
					// fmt.Println(zkk44)
					m.Save("readerXML", collection, &zkk44)
				case ".ZP504_":
					var zp504 structs.NotificationZP504
					xml.Unmarshal(inf.FileZip, &zp504)
					// fmt.Println(zp504)
					m.Save("readerXML", collection, &zp504)
				case ".OKU504_":
					var oku504 structs.NotificationOKU504
					xml.Unmarshal(inf.FileZip, &oku504)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &oku504)
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

func (c *ConsumerMQ) ChooseTemlateNotification223(msqs <-chan amqp.Delivery, inf rabbit.InformationFile, m *mongo.MongoDb) {
	collection := "xml_notification223"
	for message := range msqs {
		if err := json.Unmarshal(message.Body, &inf); err != nil {
			log.Printf("Не могу прочитать - %v", err)
		}
		pattern := []string{
			"Notice_",
			"NoticeAE_",
			"NoticeAE94_",
			"NoticeAESMBO_",
			"NoticeEP_",
			"NoticeIS_",
			"NoticeKESMBO_",
			"NoticeOA_",
			"NoticeOK_",
			"NoticeZK_",
			"NoticeZKESMBO_",
			"NoticeZPESMBO_",
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
				case "Notice_":
					var not notifications223.PurchaseNotice
					if err := xml.Unmarshal(inf.FileZip, &not); err != nil {
						log.Println(err)
					}
					// fmt.Println(ea44.FcsNotificationEF.PurchaseNumber)
					m.Save("readerXML", collection, &not)
				case "NoticeAE_":
					var notAE notifications223.PurchaseNoticeAE
					xml.Unmarshal(inf.FileZip, &notAE)
					// fmt.Println(ea615)
					m.Save("readerXML", collection, &notAE)
				case "NoticeAE94_":
					var notAE94FZ notifications223.PurchaseNoticeAE94FZ
					xml.Unmarshal(inf.FileZip, &notAE94FZ)
					m.Save("readerXML", collection, &notAE94FZ)
				case "NoticeAESMBO_":
					var notAESMBO notifications223.PurchaseNoticeAESMBO
					xml.Unmarshal(inf.FileZip, &notAESMBO)
					// fmt.Println(ok504)
					m.Save("readerXML", collection, &notAESMBO)
				case "NoticeEP_":
					var notEP notifications223.PurchaseNoticeEP
					xml.Unmarshal(inf.FileZip, &notEP)
					// fmt.Println(za44)
					m.Save("readerXML", collection, &notEP)
				case "NoticeIS_":
					var notIS notifications223.PurchaseNoticeIS
					xml.Unmarshal(inf.FileZip, &notIS)
					// fmt.Println(zk504)
					m.Save("readerXML", collection, &notIS)
				case "NoticeKESMBO_":
					var notKESMBO notifications223.PurchaseNoticeKESMBO
					xml.Unmarshal(inf.FileZip, &notKESMBO)
					// fmt.Println(zkk44)
					m.Save("readerXML", collection, &notKESMBO)
				case "NoticeOA_":
					var notOA notifications223.PurchaseNoticeOA
					xml.Unmarshal(inf.FileZip, &notOA)
					// fmt.Println(zp504)
					m.Save("readerXML", collection, &notOA)
				case "NoticeOK_":
					var notOK notifications223.PurchaseNoticeOK
					xml.Unmarshal(inf.FileZip, &notOK)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &notOK)
				case "NoticeZK_":
					var notZK notifications223.PurchaseNoticeZK
					xml.Unmarshal(inf.FileZip, &notZK)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &notZK)
				case "NoticeZKESMBO_":
					var notZKESMBO notifications223.PurchaseNoticeZKESMBO
					xml.Unmarshal(inf.FileZip, &notZKESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &notZKESMBO)
				case "NoticeZPESMBO_":
					var notZPESMBO notifications223.PurchaseNoticeZPESMBO
					xml.Unmarshal(inf.FileZip, &notZPESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &notZPESMBO)
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

func (c *ConsumerMQ) ChooseTemlateProtocols223(msqs <-chan amqp.Delivery, inf rabbit.InformationFile, m *mongo.MongoDb) {
	collection := "xml_protocols223"
	for message := range msqs {
		if err := json.Unmarshal(message.Body, &inf); err != nil {
			log.Printf("Не могу прочитать - %v", err)
		}
		pattern := []string{
			"Protocol_",
			"ProtocolCancellation_",
			"ProtocolCCAESMBO_",
			"ProtocolCCKESMBO_",
			"ProtocolCCZKESMBO_",
			"ProtocolCCZPESMBO_",
			"ProtocolCollationAESMBO_",
			"ProtocolEvasionAESMBO_",
			"ProtocolEvasionKESMBO_",
			"ProtocolEvasionZKESMBO_",
			"ProtocolEvasionZPESMBO_",
			"ProtocolFCDKESMBO_",
			"ProtocolIP_",
			"ProtocolOSZ_",
			"ProtocolPA_AE_",
			"ProtocolPA_OA_",
			"ProtocolPAAE_",
			"ProtocolPAAE94_",
			"ProtocolPAEP_",
			"ProtocolPAOA_",
			"ProtocolRZ1AE_",
			"ProtocolRZ1AESMBO_",
			"ProtocolRZ1KESMBO_",
			"ProtocolRZ1ZPESMBO_",
			"ProtocolRZ2AE_",
			"ProtocolRZ2AESMBO_",
			"ProtocolRZ2KESMBO_",
			"ProtocolRZ2ZPESMBO_",
			"ProtocolRZAE_",
			"ProtocolRZOA_",
			"ProtocolRZOK_",
			"ProtocolRZZKESMBO_",
			"ProtocolSummingupAESMBO_",
			"ProtocolSummingupKESMBO_",
			"ProtocolSummingupZKESMBO_",
			"ProtocolSummingupZPESMBO_",
			"ProtocolVK_Moskva_",
			"ProtocolZK_Moskva_",
			"ProtocolZRPZAESMBO_",
			"ProtocolZRPZZKESMBO_",
			"ProtocolZRPZZPESMBO_",
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
				case "Protocol_":
					var proto223 protocols223.PurchaseProtocol
					if err := xml.Unmarshal(inf.FileZip, &proto223); err != nil {
						log.Println(err)
					}
					m.Save("readerXML", collection, &proto223)
				case "ProtocolCancellation_":
					var protoCancel protocols223.ProtocolCancellation
					xml.Unmarshal(inf.FileZip, &protoCancel)
					// fmt.Println(ea615)
					m.Save("readerXML", collection, &protoCancel)
				case "ProtocolCCAESMBO_":
					var protoCCAESMBO protocols223.PurchaseProtocolCCAESMBO
					xml.Unmarshal(inf.FileZip, &protoCCAESMBO)
					m.Save("readerXML", collection, &protoCCAESMBO)
				case "ProtocolCCKESMBO_":
					var protoCCKESMBO protocols223.PurchaseProtocolCCKESMBO
					xml.Unmarshal(inf.FileZip, &protoCCKESMBO)
					// fmt.Println(ok504)
					m.Save("readerXML", collection, &protoCCKESMBO)
				case "ProtocolCCZKESMBO_":
					var protoCCZKESMBO protocols223.PurchaseProtocolCCZKESMBO
					xml.Unmarshal(inf.FileZip, &protoCCZKESMBO)
					// fmt.Println(za44)
					m.Save("readerXML", collection, &protoCCZKESMBO)
				case "ProtocolCCZPESMBO_":
					var protoCCZPESMBO protocols223.PurchaseProtocolCCZPESMBO
					xml.Unmarshal(inf.FileZip, &protoCCZPESMBO)
					// fmt.Println(zk504)
					m.Save("readerXML", collection, &protoCCZPESMBO)
				case "ProtocolCollationAESMBO_":
					var protoCollation protocols223.PurchaseProtocolCollationAESMBO
					xml.Unmarshal(inf.FileZip, &protoCollation)
					// fmt.Println(zkk44)
					m.Save("readerXML", collection, &protoCollation)
				case "ProtocolEvasionAESMBO_":
					var protoEvasionAESMBO protocols223.PurchaseProtocolEvasionAESMBO
					xml.Unmarshal(inf.FileZip, &protoEvasionAESMBO)
					// fmt.Println(zp504)
					m.Save("readerXML", collection, &protoEvasionAESMBO)
				case "ProtocolEvasionKESMBO_":
					var protoEvasionKESMBO protocols223.PurchaseProtocolEvasionKESMBO
					xml.Unmarshal(inf.FileZip, &protoEvasionKESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protoEvasionKESMBO)
				case "ProtocolEvasionZKESMBO_":
					var protoEvasionZKESMBO protocols223.PurchaseProtocolEvasionZKESMBO
					xml.Unmarshal(inf.FileZip, &protoEvasionZKESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protoEvasionZKESMBO)
				case "ProtocolEvasionZPESMBO_":
					var protocolEvasionZPESMBO protocols223.PurchaseProtocolEvasionZPESMBO
					xml.Unmarshal(inf.FileZip, &protocolEvasionZPESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolEvasionZPESMBO)
				case "ProtocolOSZ_":
					var protocolOSZ protocols223.PurchaseProtocolOSZ
					xml.Unmarshal(inf.FileZip, &protocolOSZ)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolOSZ)
				case "ProtocolPA_AE_":
					var protocolPAAE protocols223.PurchaseProtocolPAAE
					xml.Unmarshal(inf.FileZip, &protocolPAAE)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolPAAE)
				case "ProtocolPAAE94_":
					var protocolPAAE94 protocols223.PurchaseProtocolPAAE94FZ
					xml.Unmarshal(inf.FileZip, &protocolPAAE94)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolPAAE94)
				case "ProtocolPAOA_":
					var protocolPAOA protocols223.PurchaseProtocolPAEP
					xml.Unmarshal(inf.FileZip, &protocolPAOA)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolPAOA)
				case "ProtocolRZ1AE_":
					var protocolRZ1AE protocols223.PurchaseProtocolRZ1AE94FZ
					xml.Unmarshal(inf.FileZip, &protocolRZ1AE)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolRZ1AE)
				case "ProtocolRZ1AESMBO_":
					var protocolRZ1AESMBO protocols223.PurchaseProtocolRZ2AESMBO
					xml.Unmarshal(inf.FileZip, &protocolRZ1AESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolRZ1AESMBO)
				case "ProtocolRZ1KESMBO_":
					var protocolRZ1KESMBO protocols223.PurchaseProtocolRZ1KESMBO
					xml.Unmarshal(inf.FileZip, &protocolRZ1KESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolRZ1KESMBO)
				case "ProtocolRZ1ZPESMBO_":
					var protocolRZ1ZPESMBO protocols223.PurchaseProtocolRZ1ZPESMBO
					xml.Unmarshal(inf.FileZip, &protocolRZ1ZPESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolRZ1ZPESMBO)
				case "ProtocolRZ2AE_":
					var protocolRZ2AE protocols223.PurchaseProtocolRZ2AE94FZ
					xml.Unmarshal(inf.FileZip, &protocolRZ2AE)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolRZ2AE)
				case "ProtocolRZ2AESMBO_":
					var protocolRZ2AESMBO protocols223.PurchaseProtocolRZ2AESMBO
					xml.Unmarshal(inf.FileZip, &protocolRZ2AESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolRZ2AESMBO)
				case "ProtocolRZ2KESMBO_":
					var protocolRZ2KESMBO protocols223.PurchaseProtocolRZ2KESMBO
					xml.Unmarshal(inf.FileZip, &protocolRZ2KESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolRZ2KESMBO)
				case "ProtocolRZ2ZPESMBO_":
					var protocolRZ2ZPESMBO protocols223.PurchaseProtocolRZ2ZPESMBO
					xml.Unmarshal(inf.FileZip, &protocolRZ2ZPESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolRZ2ZPESMBO)
				case "ProtocolRZAE_":
					var protocolRZAE protocols223.PurchaseProtocolRZAE
					xml.Unmarshal(inf.FileZip, &protocolRZAE)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolRZAE)
				case "ProtocolRZOA_":
					var protocolRZOA protocols223.PurchaseProtocolRZOA
					xml.Unmarshal(inf.FileZip, &protocolRZOA)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolRZOA)
				case "ProtocolRZOK_":
					var protocolRZOK protocols223.PurchaseProtocolRZOK
					xml.Unmarshal(inf.FileZip, &protocolRZOK)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolRZOK)
				case "ProtocolRZZKESMBO_":
					var protocolRZZKESMBO protocols223.PurchaseProtocolRZZKESMBO
					xml.Unmarshal(inf.FileZip, &protocolRZZKESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolRZZKESMBO)
				case "ProtocolSummingupAESMBO_":
					var protocolSummingupAESMBO protocols223.PurchaseProtocolSummingupAESMBO
					xml.Unmarshal(inf.FileZip, &protocolSummingupAESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolSummingupAESMBO)
				case "ProtocolSummingupKESMBO_":
					var protocolSummingupKESMBO protocols223.PurchaseProtocolSummingupKESMBO
					xml.Unmarshal(inf.FileZip, &protocolSummingupKESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolSummingupKESMBO)
				case "ProtocolSummingupZKESMBO_":
					var protocolSummingupZKESMBO protocols223.PurchaseProtocolSummingupZKESMBO
					xml.Unmarshal(inf.FileZip, &protocolSummingupZKESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolSummingupZKESMBO)
				case "ProtocolSummingupZPESMBO_":
					var protocolSummingupZPESMBO protocols223.PurchaseProtocolSummingupZPESMBO
					xml.Unmarshal(inf.FileZip, &protocolSummingupZPESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolSummingupZPESMBO)
				case "ProtocolVK_":
					var protocolVK protocols223.PurchaseProtocolVK
					xml.Unmarshal(inf.FileZip, &protocolVK)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolVK)
				case "ProtocolZK_":
					var protocolZK protocols223.PurchaseProtocolZK
					xml.Unmarshal(inf.FileZip, &protocolZK)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolZK)
				case "ProtocolZRPZAESMBO_":
					var protocolZRPZAESMBO protocols223.PurchaseProtocolZRPZAESMBO
					xml.Unmarshal(inf.FileZip, &protocolZRPZAESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolZRPZAESMBO)
				case "ProtocolZRPZKESMBO_":
					var protocolZRPZKESMBO protocols223.PurchaseProtocolZRPZKESMBO
					xml.Unmarshal(inf.FileZip, &protocolZRPZKESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolZRPZKESMBO)
				case "ProtocolZRPZZKESMBO_":
					var protocolZRPZZKESMBO protocols223.PurchaseProtocolZRPZZKESMBO
					xml.Unmarshal(inf.FileZip, &protocolZRPZZKESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolZRPZZKESMBO)
				case "ProtocolZRPZZPESMBO_":
					var protocolZRPZZPESMBO protocols223.PurchaseProtocolZRPZZPESMBO
					xml.Unmarshal(inf.FileZip, &protocolZRPZZPESMBO)
					// fmt.Println(oku504)
					m.Save("readerXML", collection, &protocolZRPZZPESMBO)
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
					g := ef1.FcsProtocolEF1.ProtocolLot.Applications.Application
					for _, value := range g {
						var z models.Participant44FZ
						z.Reest = ef1.FcsProtocolEF1.PurchaseNumber
						z.Application = value
						m.Save("readerXML", "participant", z)
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
					var iefinv protocols44.ProtocolEFInv
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
