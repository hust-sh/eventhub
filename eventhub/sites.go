package main

import (
    "github.com/gin-gonic/gin"
)

type Sites interface {
    Transform(c *gin.Context) gin.H
    //SendMsg(data gin.H)
}


