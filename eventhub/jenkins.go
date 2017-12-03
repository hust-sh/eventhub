package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

type HookInfo struct {
    Name string `json:"display_name" binding:"required"`
    Build struct {
        Url string `json:"url" binding:"required"`
        Phase string `json:"phase" binding:"required"`
        Status string `json:"status"`
    } `json:"build" binding:"required"`
}

type Jenkins struct {

}

func (j Jenkins) Transform(c *gin.Context) gin.H {

    var hook HookInfo
    c.BindJSON(&hook)

    return j.PackMsg(hook)
}

func (_ Jenkins) PackMsg(h HookInfo) gin.H {

    phase := h.Build.Phase
    var title string

    if phase == "STARTED" {
        title = "Start Build #" + h.Name 
    } else if phase == "FINALIZED" {
        status := h.Build.Status
        if status == "SUCCESS" {
            title = "Build Success #" + h.Name
        } else {
            title = "Build Failed #" + h.Name
        }
    }

    text := fmt.Sprintf("#### %s \n[%s](%s)", title, "Detail", h.Build.Url)
    return gin.H {
        "msgtype": "markdown",
        "markdown": gin.H {"title": title, "text": text},
    }
}

