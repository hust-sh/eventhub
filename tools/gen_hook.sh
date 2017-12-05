#!/bin/bash

while getopts "s:u:" arg
do
    case ${arg} in
        s) 
            SITE=${OPTARG}
            ;;
        u)
            URL=${OPTARG}
            ;;
    esac
done


#curl -H "Content-Type: application/json" -d '{"url":"https://oapi.dingtalk.com/robot/send?access_token=4431ce3a5a8ac6d057b34615f254fd5e8df8d5eaa9e9c9303f6751eebd84fb31","site":"jira"}' -X POST localhost:3003/genhook

data="{\"site\":\"${SITE}\",\"url\":\"${URL}\"}"
curl -H "Content-Type: application/json" --data ${data} -X POST localhost:3003/genhook

echo ""
