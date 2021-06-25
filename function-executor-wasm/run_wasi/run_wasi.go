package main

import (
	"fmt"
	"io/ioutil"
	"log"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

var filename = "../c/loop.wasm"

func main() {
	wasmBytes, _ := ioutil.ReadFile(filename)

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, _ := wasmer.NewModule(store, wasmBytes)

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)

	if err != nil {
		log.Printf("err = %v", err.Error())
	}

	// Gets the `sum` exported function from the WebAssembly instance.
	log.Printf("instance = %v", instance)
	sum, err := instance.Exports.GetFunction("loop")

	if err != nil {
		log.Printf("err = %v", err.Error())
	}

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, _ := sum(50)

	fmt.Println(result) // 42!
}
