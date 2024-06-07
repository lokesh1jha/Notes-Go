package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lokesh1jha/ecommerce-webapp/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/user/signup", controllers.Signup())
	incomingRoutes.POST("/user/login", controllers.Login())
}
