# NATS Function Tests - WASM

This directory contains the code for testing WASM function execution via NATS.

As the code is reasonably straightforward, only a couple of high level points
are noted here; specifically:
- the configuration file contains a JSON object which includes NATS co-ordinates;
  it also includes the NATS subject trigger which is in the function definition -
  this was an artefact of largely copying the Javascript implementation
- analogoues to the Javascript case, initialization of each WASM worker involves
  loading the function - in this case, it involves creating a WASM engine,
  compiling the module and creating an instance which can then be called.
- it is assumed that the WASM binary contains a function called `loop` and that
  this is exported in the WASM compilation process. Note that this can take a
  parameter as defined in the C source and as is visible from how it is called in
  go
- the manager simply prints out the rate at which it has been able to dispatch
  NATS messages to the workers; it does nothing in terms of handling channel
  blocking/overflow conditions

It is assumed that there is a WASM binary available which should be run each
time a message is received via NATS: this is specified in the configuration
file (although the name of the exported function is currently hard-coded).

The `c` subdirectory contains the source code for the function and a simple
script to build it.

To run the system, ensure that a valid WASM file is in the correct place,
ensure that the configuration points at a valid NATS deployment, run

```
$ go mod vendor
$ go build
$ ./function-executor-wasm --config-file ./config.json --workers 4
```

