package main

import (
	"fmt"
	"log"

	"github.com/go-git/config"
	"github.com/go-git/gitreport"
)

var (
	filePath = "config.json"
)

func main() {
	fmt.Println("Git worker started.")
	if err := Worker(); err != nil {
		log.Panic(err.Error())
	}
}

func Worker() (err error) {
	// read configuration from config.json file
	configuration, err := config.ReadConfigValues(filePath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("Configuration has been read from JSON (%s) successfully.\n", filePath)

	gitHandler := gitreport.GitWorker(configuration)
	summary, err := gitHandler.FetchGitPRSummary()
	if err != nil {
		log.Println(err.Error())
		return
	}

	// send mail to scrum master
	return gitHandler.SendMail(summary)
}
