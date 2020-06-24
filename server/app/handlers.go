package app

import "github.com/gin-gonic/gin"

func ping(c *gin.Context) {
	c.String(200, "pong")

}

func versionInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"version":        Version,
	})
}
