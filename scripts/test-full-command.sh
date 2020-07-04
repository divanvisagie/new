#!/bin/bash

# make clean
# make
cd test/testdata
rm -rf ./testbed; python ../../new/main.py testbed https://github.com/divanvisagie/kotlin-tested-seed
cd ../..