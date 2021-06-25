# NATS Function Tests - Javascript

This directory contains the code for testing Javascript function execution
via NATS.

As the code is reasonably straightforward, only a couple of high level points
are noted here; specifically:
- the configuration file contains a JSON object which includes NATS co-ordinates
  and the function to be executed - note that no validation of this function is
  performed; it if is invalid Javascript, the v8 engine will return an error
  in execution
- the system comprises workers, each of which has its own v8 Isolate and a
  specific v8 Context for the function; this Context is reused each time (which
  may or may not be realistic)
- during worker initialization, the Function is created within the Isolate; as
  such, it is assumed that the Javascript code within the config file simply
  contains a valid Javascript function definition - further, it is assumed that
  this function is called `mean` as this is what is invoked each time a NATS
  message is received
- the manager simply prints out the rate at which it has been able to dispatch
  NATS messages to the workers; it does nothing in terms of handling channel
  blocking/overflow conditions

To build the system, you need a modern go toolchain and gcc and g++.

```
$ go mod vendor
$ go build
```

To run the system, simply ensure that the configuration points at a valid NATS
deployment, run

```
$ ./function-executor-javascript --config-file ./config.json --workers 4
```


