package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"regexp"

	"github.com/BurntSushi/toml"
	"github.com/Sterks/rXML/config"
	"github.com/Sterks/rXML/rabbit"
	"github.com/Sterks/rXML/structs"
)

func main() {
	configPath := "config/config.toml"
	conf := config.NewConfig()
	toml.DecodeFile(configPath, &conf)

	consumer := rabbit.NewConsumer(conf)
	msg1 := consumer.ConsumerMQNow(conf, "Notifications44OpenFile")
	var inf rabbit.InformationFile

	for message := range msg1 {
		if err := json.Unmarshal(message.Body, &inf); err != nil {
			log.Fatalf("Не могу прочитать", err)
		}

		pattern := ".EA44_"
		var not structs.NotificationEA44
		reg := inf.NameFile
		matched, err := regexp.MatchString(pattern, reg)
		if err != nil {
			log.Fatalf("Не работает regexp", err)
		}
		if matched {
			xml.Unmarshal(inf.FileZip, &not)
			fmt.Println(not.FcsNotificationEF.PurchaseNumber)
		}

	}

}
