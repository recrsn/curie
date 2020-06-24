package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	limitQueryParam  = "limit"
	offsetQueryParam = "offset"
)

func limitAndOffset(c *gin.Context) (int, int, error) {
	limitParam := c.Query(limitQueryParam)
	offsetParam := c.Query(offsetQueryParam)

	limit, err := toInt(limitParam, 0)

	if err != nil || limit < 0 {
		return -1, -1, err
	}

	offset, err := toInt(offsetParam, 0)

	if err != nil || offset < 0 {
		return -1, -1, err
	}

	return limit, offset, nil
}

func toInt(value string, defaultValue int) (int, error) {
	if value == "" {
		return defaultValue, nil
	}

	return strconv.Atoi(value)
}
