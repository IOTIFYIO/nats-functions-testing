package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	NatsServer    string     `json:"natsServer"`
	NatsPort      int        `json:"natsPort"`
	TestFunctions []Function `json:"testFunctions"`
}

func readFile(filename string) ([]byte, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file %v: %v", filename, err.Error())
		return nil, err
	}
	return dat, nil
}

func validateConfiguration(c *Configuration) bool {
	return true
}

func dumpConfiguration(c *Configuration) bool {
	return true
}

func readConfigurationFile(filename string) (*Configuration, error) {
	configBytes, err := readFile(filename)
	if err != nil {
		return nil, err
	}

	c := Configuration{}
	err = json.Unmarshal(configBytes, &c)

	if err != nil {
		log.Printf("Error unmarshalling configuration: %v", err.Error())
		return nil, err
	}
	return &c, nil
}
