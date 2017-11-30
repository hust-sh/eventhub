#!/bin/sh
White='\033[1;36m'
NC='\033[0m' # No Color
case "$1" in

    "build" )
        docker build -t goapp:latest .
    ;;
    "run" )
        docker run -d --rm -it -p 80:8080 --name eventhub goapp:latest
    ;;
    "stop" )
        docker container stop eventhub
    ;;
    * )
        echo "${White}build${NC}: build eventhub server"
        echo "${White}run  ${NC}: run eventhub server"
        echo "${White}stop ${NC}: stop eventhub server"

esac
