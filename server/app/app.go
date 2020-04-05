package app

import (
	"github.com/gin-gonic/gin"
)

type App struct {
}

func NewApp() *App {
	return &App{
	}
}

func (a *App) Mount(r gin.IRoutes) {
	r.GET("/", versionInfo)

	r.GET("/ping", ping)
}

