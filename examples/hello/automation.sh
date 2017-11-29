#!/bin/bash

White='\033[1;36m'
NC='\033[0m' # No Color

case "$1" in

    "build" )
        docker run --rm -v "$PWD":/app -w /app golang:1.9.2-alpine3.6 go build -v
        ;;
    * )
        echo "${White}build${NC}: build go project"

esac
