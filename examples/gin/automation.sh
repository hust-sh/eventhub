#!/bin/sh
White='\033[1;36m'
NC='\033[0m' # No Color
case "$1" in

    "run" )
        docker run --rm -it -p 80:8080 goapp
    ;;
    * )
        echo "${White}run${NC}: run a go server"

esac
