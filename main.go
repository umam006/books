package main

import (
	"golang-books-api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	routers.Routes(app)
	app.Run() // listen and serve on 0.0.0.0:8080
}
