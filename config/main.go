package config

import (
	"encoding/json"
	"fmt"
	"github.com/psolru/terrastate-http/safeclose"
	"log"
)

var Values layout

// Load loads the configuration file
func Load() {
	log.Println("[CONFIG] Loading...")

	f, err := getConfigFile()
	if err != nil {
		log.Fatalf("[CONFIG] Can't read config from disk: '%s'", err.Error())
	}
	defer safeclose.Close(f)

	log.Println(fmt.Sprintf("[CONFIG] Found config file '%s'", f.Name()))
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&Values)
	if err != nil {
		log.Fatalf("[CONFIG] Can't decode config: '%s'", err.Error())
	}

	err = Values.validate()
	if err != nil {
		log.Fatalf("[CONFIG] Config is invalid: '%s'", err.Error())
	}

	log.Printf("[CONFIG] Successful loaded from '%s'", f.Name())
}

// Basic Auth Credentials given
func IsAuthActive() bool {
	return Values.Username != "" && Values.Password != ""
}
