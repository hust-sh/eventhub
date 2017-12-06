# Eventhub

This is a golang hobby project.

## Get started

        # run server
        docker-compose up -d
        # generate a eventhub.webhook, remain the "webhook" field of the response
        ./tools/gen_hook.sh -s jenkins -u https://oapi.dingtalk.com/robot/send?access_token=4431ce3a5a8ac6d0  
        # 模拟Jenkins触发消息，注意将下面命令中的url 换成上条命令得到的webhook
        curl -H "Content-Type" -X POST --data @tools/test_jenkins.json http://localhost:3003/webhook/jenkins/78baf7af-0a6c-4356-80c2-70f245d46fe9
        

## Oauth

* 原理可[参考](http://www.yangluoshen.site/blog/post/yangluoshen/Oauth2)
* eventhub oauth github ([API docs](https://developer.github.com/apps/building-integrations/setting-up-and-registering-oauth-apps/)):
  * 首先创建[github app](https://github.com/settings/developers)([how to](https://developer.github.com/apps/building-integrations/setting-up-and-registering-oauth-apps/registering-oauth-apps/))
    * `Authorization callback URL`这栏填`http://localhost:3003/callback/github`
    * 创建成功后保存`client_id`和`secret`

  * 将上述的cleing_id和secret设置为环境变量(替换下述的client和secret)

        export $GITHUBKEY=client_id
        export $GITHUB_SECRET=secret

  * docker-compose build --pull && docker-compose up -d
  * 前往浏览器输入`http://localhost:3003/siteentry`


## Roadmap

* webhook generator: eventhub的webhook生成器，用于告知其他服务（如Github）如何通知我们 (done)
* log system
* Jira
* Oauth2: 很多服务需要Oauth流程（如Trello, google apps） (doing 20%)
* Trello
