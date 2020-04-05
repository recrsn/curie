package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
)

type App struct {
	db *pg.DB
}

func NewApp(db *pg.DB) *App {
	return &App{db: db}
}

func (a *App) Mount(r gin.IRoutes) {
	r.GET("/info", versionInfo)
	r.GET("/ping", ping)
}
