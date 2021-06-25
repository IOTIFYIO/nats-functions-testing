package main

import (
	"io/ioutil"
	"log"
	"sync"

	"github.com/wasmerio/wasmer-go/wasmer"
)

type FunctionExecutorWorker struct {
	workerId                int
	functionExecutorCounter int
	currentTime             int64
	lastTime                int64
	instance                *wasmer.Instance
}

func (w *FunctionExecutorWorker) Initialize(i int, filename string) {
	w.workerId = i
	wasmBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file %v - exiting... %v", cfg.WasmFile, err.Error())
	}

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, _ := wasmer.NewModule(store, wasmBytes)

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	w.instance, _ = wasmer.NewInstance(module, importObject)
}

func (w *FunctionExecutorWorker) Run(wg sync.WaitGroup, dataCh chan bool) {
	defer wg.Done()

	for done := range dataCh {

		if done {
			break
		}

		// we probably don't need to do this in the loop each time...
		sum, err := w.instance.Exports.GetFunction("loop")

		if err != nil {
			log.Printf("err = %v", err.Error())
		}

		// Calls that exported function with Go standard values. The WebAssembly
		// types are inferred and values are casted automatically.
		_, _ = sum(50)

		w.functionExecutorCounter++
	}
	log.Printf("Terminating worker...")
}
