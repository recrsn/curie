package controllers

import "github.com/gin-gonic/gin"

type Error struct {
	Message string `json:"message"`
}

func badRequest(c *gin.Context, message string) {
	c.JSON(400, Error{
		Message: message,
	})
}

func internalServerError(c *gin.Context) {
	c.JSON(500, Error{Message: "Internal server error"})
}
