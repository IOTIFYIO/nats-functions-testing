package main

import (
	"flag"
	"log"
	"os"
)

var (
	cfg                      Configuration
	defaultConfigurationFile = "./config.json"
	numberOfWorkers          int
)

func init() {
	flag.StringVar(&defaultConfigurationFile, "config-file", "/etc/iotify/load-testing/config.json", "Configuration file for this service")
	flag.IntVar(&numberOfWorkers, "workers", 4, "Number of wasm workers to use")
}

func main() {
	flag.Parse()

	cfg, err := readConfigurationFile(defaultConfigurationFile)
	if err != nil {
		log.Printf("Error reading configuration file...exiting...")
		os.Exit(1)
	}

	i := createFunctionExecutorWasm(cfg)

	// initialize sets up the NATS subscriptions
	i.Initialize(numberOfWorkers, cfg.TestFunctions)

	select {}
}
