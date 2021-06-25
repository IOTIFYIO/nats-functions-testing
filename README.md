# NATS Function Tests

This repo was used for testing function invocation with [NATS](https://www.nats.io); the objective was
to compare Javascript and [WASM](https://webassembly.org/) performance for lightweight function invocation.

For the Javascript context, v8 was used; for the WASM context, [wasmer](https://wasmer.io) was used
(as it was a little easier for us to pass a WASM binary to the executor than,
for example, running WASM directly within v8). In both cases, we had a very
simple function which just calculated a mean of the numbers 1-50 and returned
a response.

For each case, we had a function executor which executed the function whenever
a message was received on NATS; note that the contents of the message was
ignored and, in particular, it was not passed to the function - the point here
was to perform simple load testing in a somewhat rudimentary way.

The function executors were written in go: in the Javascript case the [v8-go](https://github.com/rogchap/v8go)
module was used and in the WASM case the [wasmer-go](https://github.com/wasmerio/wasmer-go) module was used. In both cases,
the number of concurrent executors is configurable but as v8 is designed to
operate on a single thread and WASM thread support is relatively new (and not
used here), there is little gain to be had from having a number of executors
significantly larger than the number of possible concurrent threads in the system.

The experimental setup was simple: the function executor was run which would
listen on a NATS subject and `natsbench` was used to publish to that subject.
All of this was run on a single VM.  As there was no co-ordination between the
function executor and `natsbench`, (ie the function executor did not know when
an experiment 'started' and 'ended') the executor just simply printed out the
rate of function executions after every 1000 function executions - this was
sufficient for us to see the difference in the rate of function executions
between the JS context and the WASM context.

Sample output from each of the two experiments is shown below - this was run
in a single VM with 4 vCPUs and 8GB RAM.

Below is the output for running the javascript version with 4 workers.

```
2021/06/25 15:11:58 Executed 648000 functions (rate 293317/sec)
2021/06/25 15:11:58 Executed 649000 functions (rate 333448/sec)
2021/06/25 15:11:58 Executed 650000 functions (rate 318977/sec)
2021/06/25 15:11:58 Executed 651000 functions (rate 283696/sec)
2021/06/25 15:11:58 Executed 652000 functions (rate 246697/sec)
2021/06/25 15:11:58 Executed 653000 functions (rate 329070/sec)
2021/06/25 15:11:58 Executed 654000 functions (rate 364161/sec)
2021/06/25 15:11:58 Executed 655000 functions (rate 351889/sec)
2021/06/25 15:11:58 Executed 656000 functions (rate 298963/sec)
2021/06/25 15:11:58 Executed 657000 functions (rate 355048/sec)
2021/06/25 15:11:58 Executed 658000 functions (rate 294956/sec)
2021/06/25 15:11:58 Executed 659000 functions (rate 334852/sec)
2021/06/25 15:11:58 Executed 660000 functions (rate 239363/sec)
2021/06/25 15:11:58 Executed 661000 functions (rate 259989/sec)
2021/06/25 15:11:58 Executed 662000 functions (rate 137818/sec)
2021/06/25 15:11:58 Executed 663000 functions (rate 284385/sec)
2021/06/25 15:11:58 Executed 664000 functions (rate 251784/sec)
2021/06/25 15:11:58 Executed 665000 functions (rate 393819/sec)
2021/06/25 15:11:58 Executed 666000 functions (rate 266860/sec)
2021/06/25 15:11:58 Executed 667000 functions (rate 286899/sec)
2021/06/25 15:11:58 Executed 668000 functions (rate 347201/sec)
2021/06/25 15:11:58 Executed 669000 functions (rate 359996/sec)
2021/06/25 15:11:58 Executed 670000 functions (rate 480824/sec)
2021/06/25 15:11:58 Executed 671000 functions (rate 442599/sec)
2021/06/25 15:11:58 Executed 672000 functions (rate 313409/sec)
2021/06/25 15:11:58 Executed 673000 functions (rate 228326/sec)
2021/06/25 15:11:58 Executed 674000 functions (rate 302962/sec)
2021/06/25 15:11:58 Executed 675000 functions (rate 381376/sec)
2021/06/25 15:11:58 Executed 676000 functions (rate 375202/sec)
2021/06/25 15:11:58 Executed 677000 functions (rate 346653/sec)
2021/06/25 15:11:58 Executed 678000 functions (rate 233409/sec)
2021/06/25 15:11:58 Executed 679000 functions (rate 241470/sec)
2021/06/25 15:11:58 Executed 680000 functions (rate 252902/sec)
2021/06/25 15:11:58 Executed 681000 functions (rate 357827/sec)
2021/06/25 15:11:58 Executed 682000 functions (rate 325747/sec)
2021/06/25 15:11:58 Executed 683000 functions (rate 335310/sec)
2021/06/25 15:11:58 Executed 684000 functions (rate 287617/sec)
2021/06/25 15:11:58 Executed 685000 functions (rate 375001/sec)
2021/06/25 15:11:58 Executed 686000 functions (rate 300716/sec)
```

Similar output for wasm variant, also with 4 workers.
```
2021/06/25 15:14:20 Executed 660000 functions (rate 580547/sec)
2021/06/25 15:14:20 Executed 661000 functions (rate 592077/sec)
2021/06/25 15:14:20 Executed 662000 functions (rate 588861/sec)
2021/06/25 15:14:20 Executed 663000 functions (rate 570530/sec)
2021/06/25 15:14:20 Executed 664000 functions (rate 595196/sec)
2021/06/25 15:14:20 Executed 665000 functions (rate 574569/sec)
2021/06/25 15:14:20 Executed 666000 functions (rate 493640/sec)
2021/06/25 15:14:20 Executed 667000 functions (rate 536504/sec)
2021/06/25 15:14:20 Executed 668000 functions (rate 469394/sec)
2021/06/25 15:14:20 Executed 669000 functions (rate 455750/sec)
2021/06/25 15:14:20 Executed 670000 functions (rate 365056/sec)
2021/06/25 15:14:20 Executed 671000 functions (rate 405657/sec)
2021/06/25 15:14:20 Executed 672000 functions (rate 535934/sec)
2021/06/25 15:14:20 Executed 673000 functions (rate 587167/sec)
2021/06/25 15:14:20 Executed 674000 functions (rate 466539/sec)
2021/06/25 15:14:20 Executed 675000 functions (rate 338960/sec)
2021/06/25 15:14:20 Executed 676000 functions (rate 573562/sec)
2021/06/25 15:14:20 Executed 677000 functions (rate 597145/sec)
2021/06/25 15:14:20 Executed 678000 functions (rate 402840/sec)
2021/06/25 15:14:20 Executed 679000 functions (rate 307189/sec)
2021/06/25 15:14:20 Executed 680000 functions (rate 590044/sec)
2021/06/25 15:14:20 Executed 681000 functions (rate 477518/sec)
2021/06/25 15:14:20 Executed 682000 functions (rate 547422/sec)
2021/06/25 15:14:20 Executed 683000 functions (rate 473975/sec)
2021/06/25 15:14:20 Executed 684000 functions (rate 367740/sec)
2021/06/25 15:14:20 Executed 685000 functions (rate 459700/sec)
2021/06/25 15:14:20 Executed 686000 functions (rate 411708/sec)
2021/06/25 15:14:20 Executed 687000 functions (rate 482606/sec)
2021/06/25 15:14:20 Executed 688000 functions (rate 458527/sec)
2021/06/25 15:14:20 Executed 689000 functions (rate 566629/sec)
2021/06/25 15:14:20 Executed 690000 functions (rate 600877/sec)
2021/06/25 15:14:20 Executed 691000 functions (rate 372204/sec)
2021/06/25 15:14:20 Executed 692000 functions (rate 560664/sec)
2021/06/25 15:14:20 Executed 693000 functions (rate 549158/sec)
2021/06/25 15:14:20 Executed 694000 functions (rate 592713/sec)
2021/06/25 15:14:20 Executed 695000 functions (rate 658067/sec)
2021/06/25 15:14:20 Executed 696000 functions (rate 630333/sec)
2021/06/25 15:14:20 Executed 697000 functions (rate 340579/sec)
```

More details on how to run the specific function executors are provided in the
subdirectories within this repo.

