#!/bin/bash

# Build the binary outside src
cd src
go build -ldflags "-s -w" -o ..
cd ..
