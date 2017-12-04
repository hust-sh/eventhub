# Eventhub

This is a golang hobby project.

## Get started

        docker-compose up -d
        curl localhost:3003/
        curl -H "Content-Type:application/json" -X POST -d '{"display_name":"hub", "build":{"url":"job/hub/1", "phase":"STARTED"}}' 0.0.0.0:3003/webhook/jenkins/123

## Roadmap

* webhook generator: eventhub的webhook生成器，用于告知其他服务（如Github）如何通知我们
* log system
* Jira
* Oauth2: 很多服务需要Oauth流程（如Trello, google apps）
* Trello
