package main

import (
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/Sterks/rXML/config"
	"github.com/Sterks/rXML/rabbit"
)

func main() {
	configPath := "config/config.toml"
	conf := config.NewConfig()
	toml.DecodeFile(configPath, &conf)

	consumer := rabbit.NewConsumer(conf)
	msg1 := consumer.ConsumerMQNow(conf, "Notifications44OpenFile")
	forever := make(chan bool)

	var inf rabbit.InformationFile

	for message := range msg1 {
		err := json.Unmarshal(message.Body, &inf)
	}

}
