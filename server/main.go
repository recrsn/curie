package main

import (
	app "github.com/curie/server/app"
	"github.com/gin-gonic/gin"
)

func main() {
	a := app.NewApp()
	r := gin.Default()
	a.Mount(r)

	err := r.Run()

	if err != nil {
		panic(err)
	}
}
