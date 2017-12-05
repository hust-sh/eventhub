package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "log"
)


func setRouter() *gin.Engine {

    router := gin.New()
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Welcome Eventhub")
    })

    router.POST("/webhook/:site/:access_token", webhookHandler) 

    router.POST("/genhook", genWebhookHandler)

    return router
}


var SiteMapper = map[string]Sites{
    "jenkins": Jenkins{},
}


func webhookHandler(c *gin.Context) {
    site := c.Param("site")
    accessToken := c.Param("access_token")
    siteHandler, err := SiteMapper[site]
    if !err {
        log.Printf("Unknown Site:%s", site)
        c.JSON(http.StatusOK, gin.H{"err": "unknown site"})
        return
    }
    webhook, err2 := GetWebhook(site, accessToken)
    if err2 != nil {
        log.Printf("Get webhook failed: %v", err2)
        c.JSON(http.StatusOK, gin.H{"err": "Invalid accesstoken"})
        return
    }
    
    data := siteHandler.Transform(c)
    siteHandler.SendMsg(webhook, data)
    log.Printf("data: %v", data)
    c.JSON(http.StatusOK, data)
}


type GenHookReq struct {
    SiteName string `json:"site", binding:"required"`
    Url string `json:"url", binding:"required"`
}


func genWebhookHandler(c *gin.Context) {

    var req GenHookReq
    c.BindJSON(&req)

    site := req.SiteName
    url := req.Url

    if !IsValidSiteType(site) {
        log.Printf("Invalid site type:%s", site)
        c.JSON(http.StatusOK, gin.H{"err": "Inavlid site type"})
        return
    }

    redisCli, err := GetRedis()
    if err != nil {
        log.Printf("Dial redis failed. %v", err)
        c.JSON(http.StatusOK, gin.H{"err": "Dial redis Failed"})
        return
    }

    accessToken := GenAccessToken()
    _, err = redisCli.Do("HSET", site, accessToken, url)
    if err != nil {
        log.Printf("Hset redis failed, %v", err)
        c.JSON(http.StatusOK, gin.H{"err": "Redis Exception"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "err": "Success",
        "webhook": GenWebhook(site, accessToken, c.Request),
    })
}
