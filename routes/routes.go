package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lokesh1jha/ecommerce-webapp.git/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/user/signup", controllers.Signup())
	incomingRoutes.POST("/user/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())

	incomingRoutes.GET("/user/productview", controllers.SearchProduct())
	incomingRoutes.GET("/user/search", controllers.SearchProductsByQuery())
}
