package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

type FunctionExecutorWasm struct {
	NatsServer            string
	NatsServerPort        int
	FunctionSubscriptions map[string](*nats.Subscription)
	Nc                    *nats.Conn
	workers               []FunctionExecutorWorker
	dataCh                chan bool
	messagesProcessed     int
	currentTime           int64
	lastTime              int64
	wasmFile              string
}

func createFunctionExecutorWasm(c *Configuration) *FunctionExecutorWasm {
	log.Printf("configuration = %v", *c)
	return &FunctionExecutorWasm{
		NatsServer:     c.NatsServer,
		NatsServerPort: c.NatsPort,
		wasmFile:       c.WasmFile,
	}
}

func (i *FunctionExecutorWasm) setupNats() {
	natsUrl := fmt.Sprintf("nats://%v:%v", i.NatsServer, i.NatsServerPort)
	var err error
	i.Nc, err = nats.Connect(natsUrl)
	if err != nil {
		log.Printf("Unable to connect to NATS Server - error: %v...Exiting...", err.Error())
		os.Exit(1)
	}
}

func (i *FunctionExecutorWasm) createNatsSubscription(f Function) {
	if i.Nc == nil {
		i.setupNats()
	}

	log.Printf("Creating NATS subscription on subject %v, triggering function %v", f.NatsSubjectTrigger, f.Name)
	_, err := i.Nc.Subscribe(f.NatsSubjectTrigger, i.messageHandler)
	if err != nil {
		log.Printf("Error subscribing to NATS channel %v", err.Error())
	}
}

func (i *FunctionExecutorWasm) initializeWorkers(numberOfWorkers int, functions []Function) {

	var wg sync.WaitGroup

	i.dataCh = make(chan bool, numberOfWorkers)
	for idx := 0; idx < numberOfWorkers; idx++ {
		wg.Add(1)
		worker := FunctionExecutorWorker{}
		worker.Initialize(idx, i.wasmFile)
		i.workers = append(i.workers, worker)
		go worker.Run(wg, i.dataCh)
	}

	i.createNatsSubscription(functions[0])

	wg.Wait()
	log.Printf("Finished!")

}

func (i *FunctionExecutorWasm) Initialize(numberOfWorkers int, functions []Function) {
	i.FunctionSubscriptions = make(map[string](*nats.Subscription))
	i.initializeWorkers(numberOfWorkers, functions)
}

func (i *FunctionExecutorWasm) messageHandler(m *nats.Msg) {

	//log.Printf("Sending true to channel...")
	i.messagesProcessed++
	if i.messagesProcessed > 999999 {
		// exit - enough work has been done...
		i.dataCh <- true
	} else {
		// send any message which is not true to trigger function execution
		i.dataCh <- false
	}

	if i.messagesProcessed%1000 == 0 {
		i.currentTime = time.Now().UnixNano()
		timeDelta := i.currentTime - i.lastTime
		functionExecutionRate := (1000 * 1000000000) / timeDelta
		log.Printf("Executed %v functions (rate %v/sec)", i.messagesProcessed, functionExecutionRate)
		i.lastTime = i.currentTime
	}
}
