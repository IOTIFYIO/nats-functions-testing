# Compiling to WASM

This directory contains a few bits and pieces that were used when trying to
understand the best approach to create a WASM executable from a simple C
program.

The main point here is the `loop.c` file which is used to create the
`loop.wasm` executable which is used by the Function Executor in the directory
above. The shell script `./create_loop_wasm.sh` can be used to create the
binary. Note that it assumes `clang` and `llvm` are installed. (Also note on
ubuntu that installing modern `clang`, `llvm` and `lld` installs the correct
binaries, but `wasm-ld` may have a link missing and `clang` may complain that
it cannot find `wasm-ld` - it may be necessary to
`ln -s /usr/bin/wasm-ld-10 /usr/bin/wasm-ld`


