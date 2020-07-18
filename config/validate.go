package config

import (
	"fmt"
	"log"
)

// c.validate runs a few validations against the loaded config file
func (c layout) validate() error {
	log.Println("[CONFIG] Validating...")
	var stateNames = make(map[string]struct{})
	for _, state := range c.States {
		// Address must be given
		if state.Address == "" {
			return fmt.Errorf("[CONFIG] no address: '%s'", state.Name)
		}
		// Address must be unique
		if _, exists := stateNames[state.Address]; exists {
			return fmt.Errorf("[CONFIG] statename conflict, please check: '%s'", state.Name)
		}
	}
	return nil
}
