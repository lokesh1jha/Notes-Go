package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lokesh1jha/ecommerce-webapp.git/database"
	"github.com/lokesh1jha/ecommerce-webapp.git/models"
	"go.mongodb.org/mongo-driver/bson"
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

	}
}

func (app *Application) RemoveItemFromCart() gin.HandlerFunc {
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

		err = database.RemoveCartItem(ctx, app.ProductCollection, app.UserCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, "successfully removed from cart")

	}
}

func (app *Application) GetItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			log.Println("user id is empty")
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid user id"})
			return
		}

		userQueryID, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid user id"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var filledCart models.User

		err = userCollection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: userQueryID}}).Decode(&filledCart)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		match := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: userQueryID}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$userCart"}}}}
		group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{primitive.E{Key: "$sum", Value: "$cartItems.price"}}}}}}
		pointercursor, err := userCollection.Aggregate(ctx, mongo.Pipeline{match, unwind, group})

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "something went wrong")
			return
		}
		var listing []bson.M

		if err = pointercursor.All(ctx, &listing); err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, "something went wrong")
			return
		}

		for _, json := range listing {
			c.IndentedJSON(http.StatusOK, json["total"])
			c.IndentedJSON(http.StatusOK, filledCart.UserCart)
		}
		ctx.Done()

	}
}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("userID")

		if userQueryID == "" {
			log.Println("user id is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "invalid user id"})
			c.Abort()
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		err := database.BuyItemFromCart(ctx, app.UserCollection, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, "successfully placed the order")

	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
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

		err = database.InstantBuyer(ctx, app.ProductCollection, app.UserCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, "successfully placed the order")

	}
}
