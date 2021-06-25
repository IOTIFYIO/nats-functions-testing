#! /usr/bin/env bash

clang --target=wasm32 --no-standard-libraries -Wl,--export-all -Wl,--no-entry -o loop.wasm loop.c

