package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lokesh1jha/ecommerce-webapp.git/controllers"
	"github.com/lokesh1jha/ecommerce-webapp.git/database"
	"github.com/lokesh1jha/ecommerce-webapp.git/middleware"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	router.UserRoutes(router)
	router.Use(middleware.Authenticate())

	router.GET("/add-to-cart", app.AddToCart())
	router.GET("/remove-item-from-cart", app.RemoveItemFromCart())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantBuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
