package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./html", false)))

	r.GET("/sse", serverPush)
	r.Run(":3030")
}

var counter = 0

func serverPush(c *gin.Context) {
	counter++
	c.SSEvent("message", counter)
}
