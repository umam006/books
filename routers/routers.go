package routers

import (
	"golang-books-api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	user := new(controllers.UserController)
	r.GET("/", user.Index)
}
