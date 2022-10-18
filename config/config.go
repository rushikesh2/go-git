package config

import (
	"encoding/json"
	"os"
)

// Configuration holds config file data
type Configuration struct {
	SenderEmail   string `json:"senderEmail"`
	ReceiverEmail string `json:"receiverEmail"`
	BaseURL       string `json:"baseUrl"`
	PreviousDays  int    `json:"previousDays"`
	Repository    string `json:"repository"`
}

// ReadConfigValues read config data from JSON-file
func ReadConfigValues(filePath string) (config Configuration, err error) {
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		return
	}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return
	}
	file.Close()
	return
}
