package roborooney

import (
	"encoding/json"
	"log"
	"os"
)

const configFilePath = "config.json"

func (cred *Credentials) Read() {
	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal("Couldn't open config file!", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&cred)
	if err != nil {
		log.Fatal("Couldn't parse config file: ", err.Error())
	}
}
