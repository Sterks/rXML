package services

import (
	"github.com/Sterks/rXML/config"
)

// ParserXML Интерфейс сервиса
type ParserXML interface {
	// Connect(*config.Config)

	ConfigureXMLNotification(*config.Config)
	ParseXML()
	// Validation()
}

//Do Парсинг XML
func Do(parser ParserXML, config *config.Config) {
	// parser.Connect(config)
	parser.ConfigureXMLNotification(config)
	stopChan := make(chan bool)
	go parser.ParseXML()
	<-stopChan
	// parser.Validation()
}
