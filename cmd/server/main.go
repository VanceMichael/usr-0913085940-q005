package main

import (
    "log"
    "github.com/gin-gonic/gin"
)

func main() {
    mux := gin.New()
    mux.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
    log.Println("服务监听 0.0.0.0:8080")
    log.Fatal(mux.Run(":8080"))
}
