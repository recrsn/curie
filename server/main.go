package main

import (
	app "github.com/curie/server/app"
	"github.com/curie/server/models"
	"github.com/curie/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
)

func main() {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "example",
		Database: "curie",
	})

	utils.Must(models.Load(db))

	a := app.NewApp(db)
	r := gin.Default()
	a.Mount(r)

	utils.Must(r.Run())
}
