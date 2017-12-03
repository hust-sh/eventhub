# Eventhub

This is a golang hobby project.

## Get started

        docker-compose up -d
        curl localhost:3003/
        curl -H "Content-Type:application/json" -X POST -d '{"display_name":"hub", "build":{"url":"job/hub/1", "phase":"STARTED"}}' 0.0.0.0:3003/webhook/jenkins/123
