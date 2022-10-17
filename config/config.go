package config

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration holds config file data
type Configuration struct {
	SenderEmail   string `json:"senderEmail"`
	ReceiverEmail string `json:"receiverEmail"`
	GitURL        string `json:"giturl"`
	PreviousDays  int    `json:"previousDays"`
}

// ReadConfigValues read config data from JSON-file
func ReadConfigValues(filePath string) (config Configuration, err error) {
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("Configuration has been read from JSON (%s) successfully.\n", filePath)
	return
}
