package main

import (
	"log"
	"sync"

	"rogchap.com/v8go"
)

type FunctionExecutorWorker struct {
	workerId                int
	v8Context               *v8go.Context
	functionExecutorCounter int
	currentTime             int64
	lastTime                int64
}


func (w *FunctionExecutorWorker) Initialize(i int, fd string) {
	w.workerId = i
	w.setupV8Context(fd, "test-function")
}

func (w *FunctionExecutorWorker) Run(wg sync.WaitGroup, dataCh chan bool) {
	defer wg.Done()

	for done := range dataCh {

		if done {
			break
		}

		_, err := w.v8Context.RunScript("mean(50);", "test-file") // return a value in JavaScript back to Go
		if err != nil {
			log.Printf("Error running script: %v", err.Error())
		}
		//log.Printf("val = %v", val)

		w.functionExecutorCounter++
	}
	log.Printf("Terminating worker...")
}

func (w *FunctionExecutorWorker) setupV8Context(fd string, filename string) {
	iso, _ := v8go.NewIsolate()                   // create a new VM
	global, _ := v8go.NewObjectTemplate(iso)      // a template that represents a JS Object
	w.v8Context, _ = v8go.NewContext(iso, global) // creates a new V8 context with a new Isolate aka VM
	_ = w.v8Context.Global()

	log.Printf("Running initial script %v", fd)
	_, err := w.v8Context.RunScript(fd, filename) // return a value in JavaScript back to Go
	if err != nil {
		log.Printf("Error running script: %v", err.Error())
	}
}
