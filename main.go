package main

import (
	"github.com/BurntSushi/toml"
	"github.com/Sterks/rXML/config"
	"github.com/Sterks/rXML/services"
	"os"
)

func main() {
	MainReader()
}

func MainReader() {
	configPath := ""
	getenv := os.Getenv("APPLICATION")
	if getenv == "production" {
		configPath = "config/config.prod.toml"
	} else {
		configPath = "config/config.toml"
	}
	conf := config.NewConfig()
	toml.DecodeFile(configPath, &conf)

	c := make(chan bool)
	go func() {
		parser := services.NewNotification44()
		services.Do(parser, conf)
	}()

	go func() {
		proto44 := services.NewProtocol44()
		services.Do(proto44, conf)
	}()

	go func() {
		not223 := services.NewNotifications223()
		services.Do(not223, conf)
	}()

	go func() {
		proto223 := services.NewProtocols223()
		services.Do(proto223, conf)
	}()
	<-c
}

// MainReader ...
// func MainReader() {
// 	configPath := "config/config.toml"
// 	conf := config.NewConfig()
// 	toml.DecodeFile(configPath, &conf)

// 	consumer := rabbit.NewConsumer(conf)
// 	msg1 := consumer.ConsumerMQNow(conf, "Notifications44OpenFile")

// 	var inf rabbit.InformationFile

// 	client := mongo.MongoDBNew(conf)
// 	client.MongoDBConnect()
// 	for message := range msg1 {
// 		if err := json.Unmarshal(message.Body, &inf); err != nil {
// 			log.Fatalf("Не могу прочитать - %v", err)
// 		}

// 		// Обработка извещений типа "Электронный аукцион"
// 		pattern1 := ".EA44_"

// 		var not structs.NotificationEA44
// 		reg := inf.NameFile
// 		matched, err := regexp.MatchString(pattern1, reg)
// 		if err != nil {
// 			log.Fatalf("Не работает regexp - %v", err)
// 		}
// 		if matched {
// 			var m models.Auctions
// 			if err6 := xml.Unmarshal(inf.FileZip, &not); err6 != nil {
// 				log.Fatalf("Не могу разобрать файл - %v", err6)
// 			}

// 			XMLCollection := client.MongoDb.Database("readerXML").Collection("xml_notification")
// 			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 			XMLCollection.InsertOne(ctx, not)

// 			// Реестровый номер
// 			m.RegistryNumber = not.FcsNotificationEF.PurchaseNumber
// 			// Лот 036030025432000001200
// 			m.Lot = not.FcsNotificationEF.Lot.PurchaseObjects.PurchaseObject.Quantity.Value
// 			// Название аукциона
// 			m.NameAu = not.FcsNotificationEF.Lot.PurchaseObjects.PurchaseObject.Name
// 			// НМЦ
// 			m.PriceStart, _ = strconv.ParseFloat(not.FcsNotificationEF.Lot.MaxPrice, 64)

// 			// Конечная цена
// 			// TODO пока не знаю где брать конечную дату
// 			//m.PriceEnd, _ = strconv.ParseFloat(not.FcsNotificationEF.)

// 			// Процент для обеспечения
// 			m.ProvisionAsPercent = not.FcsNotificationEF.Lot.CustomerRequirements.CustomerRequirement.ContractGuarantee.Part

// 			// Обеспечение
// 			m.Provision = not.FcsNotificationEF.Lot.CustomerRequirements.CustomerRequirement.ContractGuarantee.Amount
// 			if m.Provision == 0 {
// 				m.SetProvision()
// 			}

// 			fmt.Printf("Реестовый номер - %v \n", m.RegistryNumber)
// 			fmt.Printf("Лот - %v \n", m.Lot)
// 			fmt.Printf("Название аукциона - %s \n", m.NameAu)
// 			fmt.Printf("НМЦ - %6.2f \n", m.PriceStart)
// 			fmt.Printf("Обеспечение - %6.2f \n", m.Provision)
// 			fmt.Println("-----------------")

// 			if m.Provision == 0 {
// 				o, _ := os.Create("NotAmount.log")
// 				str := fmt.Sprintf("Нет обеспечения - %s", inf.NameFile)
// 				_, _ = o.WriteString(str)
// 			}
// 		}

// 		// Обработка извещений типа "Электронный аукцион"
// 		pattern2 := ".EA615_"
// 		var not2 structs.NotificationEA615
// 		matched2, err := regexp.MatchString(pattern2, reg)
// 		if err != nil {
// 			log.Fatalf("Не работает regexp - %v", err)
// 		}
// 		if matched2 {
// 			xml.Unmarshal(inf.FileZip, &not2)
// 		}

// 		// Электронный аукцион на оказание услуг или выполнение работ по капитальному ремонту общего имущества в многоквартирном доме
// 		pattern3 := ".INM111_"
// 		var not3 structs.NotificationNM111
// 		matched3, err := regexp.MatchString(pattern3, reg)
// 		if err != nil {
// 			log.Fatalf("Не работает regexp - %v", err)
// 		}
// 		if matched3 {
// 			xml.Unmarshal(inf.FileZip, &not3)
// 		}

// 		// Открытый конкурс в электронной форме
// 		pattern4 := ".OK504_"
// 		var not4 structs.NotificationOK504
// 		matched4, err := regexp.MatchString(pattern4, reg)
// 		if err != nil {
// 			log.Fatalf("Не работает regexp - %v", err)
// 		}
// 		if matched4 {
// 			xml.Unmarshal(inf.FileZip, &not4)
// 		}

// 		// Закрытый аукцион
// 		pattern5 := ".ZA44_"
// 		var not5 structs.NotificationZA44
// 		matched5, err := regexp.MatchString(pattern5, reg)
// 		if err != nil {
// 			log.Fatalf("Не работает regexp - %v", err)
// 		}
// 		if matched5 {
// 			xml.Unmarshal(inf.FileZip, &not5)
// 		}

// 		// Запрос котировок в электронной форме
// 		pattern6 := ".ZK504_"
// 		var not6 structs.NotificationZK504
// 		matched6, err := regexp.MatchString(pattern6, reg)
// 		if err != nil {
// 			log.Fatalf("Не работает regexp - %v", err)
// 		}
// 		if matched6 {
// 			xml.Unmarshal(inf.FileZip, &not6)
// 		}

// 		// Запрос котировок в электронной форме
// 		pattern7 := ".ZKK44_"
// 		var not7 structs.NotificationZKKP44
// 		matched7, err := regexp.MatchString(pattern7, reg)
// 		if err != nil {
// 			log.Fatalf("Не работает regexp - %v", err)
// 		}
// 		if matched7 {
// 			xml.Unmarshal(inf.FileZip, &not7)
// 		}

// 		// Запрос предложений в электронной форме
// 		pattern8 := ".ZP504_"
// 		var not8 structs.NotoficationZP504
// 		matched8, err := regexp.MatchString(pattern8, reg)
// 		if err != nil {
// 			log.Fatalf("Не работает regexp - %v", err)
// 		}
// 		if matched8 {
// 			xml.Unmarshal(inf.FileZip, &not8)
// 		}

// 		// Открытый конкурс в электронной форме
// 		pattern9 := ".OKU504_"
// 		var not9 structs.NotificationOKU504
// 		matched9, err := regexp.MatchString(pattern9, reg)
// 		if err != nil {
// 			log.Fatalf("Не работает regexp - %v", err)
// 		}
// 		if matched9 {
// 			xml.Unmarshal(inf.FileZip, &not9)
// 		}
// 		fmt.Println(inf.NameFile)
// 	}
// }

// ChooseType ...
// func ChooseType(pattern string, inf rabbit.InformationFile, not interface{}) {
// 	matched, err := regexp.MatchString(pattern, inf.NameFile)
// 	if err != nil {
// 		log.Fatalf("Не работает regexp", err)
// 	}
// 	if matched {
// 		xml.Unmarshal(inf.FileZip, &not)
// 	}
// 	fmt.Println(not)
// }
