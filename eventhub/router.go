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

    return router
}


var SiteMapper = map[string]Sites{
    "jenkins": Jenkins{},
}

func webhookHandler(c *gin.Context) {
    site := c.Param("site")
    //accessToken := c.Param("access_token")
    siteHandler, err := SiteMapper[site]
    if !err {
        log.Printf("Unknown Site:%s", site)
        c.JSON(http.StatusOK, gin.H{"err": "unknown site"})
    }
    
    data := siteHandler.Transform(c)
    log.Printf("data: %v", data)
    c.JSON(http.StatusOK, data)
    
    /*
    var hook HookInfo
    c.BindJSON(&hook)
    log.Printf("site:%s, access_token:%s, hook:%v", site, accessToken, hook)
    c.JSON(http.StatusOK, gin.H{"site": site, "access_token": accessToken, "url": hook.Build.Url})
    */

}
