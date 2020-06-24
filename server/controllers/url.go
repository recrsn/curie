package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/curie/server/repositories"
)

func GetURLs(repo *repositories.UrlRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, offset, err := limitAndOffset(c)

		if err != nil {
			badRequest(c, "Limit and Offset must be positive integers")
			return
		}

		urls, err := repo.GetAll(limit, offset)

		if err != nil {
			internalServerError(c)
			return
		}

		c.JSON(200, urls)

	}
}
