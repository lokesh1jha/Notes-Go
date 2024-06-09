package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lokesh1jha/ecommerce-webapp.git/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	ProductCollection *mongo.Collection
	UserCollection    *mongo.Collection
}

func NewApplication(productCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		ProductCollection: productCollection,
		UserCollection:    userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {

	return func(c *gin.Context) {
		productQueryId := c.Query("id")

		if productQueryId == "" {
			log.Println("product id is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "invalid product id"})
			c.Abort()
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "invalid user id"})
			c.Abort()
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryId)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, "product id is not valid")
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.AddProductToCart(ctx, app.ProductCollection, app.UserCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, "successfully added to cart")
		defer cancel()

	}
}

func RemoveItemFromCart() gin.HandlerFunc {
}

func GetItemFromCart() gin.HandlerFunc {
}

func BuyFromCart() gin.HandlerFunc {
}

func InstantBuy() gin.HandlerFunc {
}
