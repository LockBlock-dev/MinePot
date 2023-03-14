#!/bin/bash

# Build the binary outside src
cd src
go build -o ..
cd ..
