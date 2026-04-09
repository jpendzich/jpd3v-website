#!/bin/bash

set -e

BuildDir="$1"

if [ -z $BuildDir ]; then
    echo "Using default build directory"
    BuildDir="./build"
fi

if ! [ -d $BuildDir ]; then
    echo "Creating directory"
    mkdir $BuildDir
fi

go build -o $BuildDir/server ./cmd/server/main.go

cp -r ./templates $BuildDir/templates

cp -r ./static $BuildDir/static

echo "Build successful"
