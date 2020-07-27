#!/bin/bash

# make clean
# make
pip install -e .
cd test/testdata
# rm -rf ./testbed; new testbed https://github.com/divanvisagie/kotlin-tested-seed
rm -rf ./testbed; new testbed divanvisagie/new
cd ../..