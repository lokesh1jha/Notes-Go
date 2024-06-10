package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lokesh1jha/ecommerce-webapp.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {

		user_id := c.Query("id")
		if user_id == "" {
			log.Println("user id is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "invalid user id"})
			c.Abort()
			return
		}

		address, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			log.Panic(err)
			c.IndentedJSON(http.StatusInternalServerError, "Something went wrong")
			c.Abort()
			return
		}

		var addressInfo models.Address

		addressInfo.Address_id = primitive.NewObjectID()

		if err := c.BindJSON(&addressInfo); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		match_filter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: address}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$address"}}}}
		group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$address_id"}, {Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}}}}}

		cursor, err := userCollection.Aggregate(ctx, mongo.Pipeline{match_filter, unwind, group})

		if err != nil {
			log.Panic(err)
			c.IndentedJSON(http.StatusInternalServerError, "Error occured while fetching address from db")
			c.Abort()
			return
		}

		var address_list []bson.M
		if err = cursor.All(ctx, &address_list); err != nil {
			log.Panic(err)
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
			c.Abort()
			return
		}

		var size int32
		for _, address := range address_list {
			count := address["count"]
			size = count.(int32)
		}
		if size > 2 {
			c.IndentedJSON(http.StatusInternalServerError, "Cannot add more than 3 addresses")
			c.Abort()
			return
		}

		filter := bson.D{primitive.E{Key: "_id", Value: address}}
		update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "address", Value: addressInfo}}}}
		_, err = userCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Panic(err)
			c.IndentedJSON(http.StatusInternalServerError, "Something went wrong")
			c.Abort()
			return
		}
		c.IndentedJSON(http.StatusOK, "Added Successfully")

	}
}

func EditAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			log.Println("user id is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "invalid user id"})
			c.Abort()
			return
		}

		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, "Something went wrong")
			c.Abort()
			return
		}

		var address models.Address

		if err := c.BindJSON(&address); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
			return
		}

		if address.House == "" || address.Street == "" || address.City == "" || address.Pincode == "" {
			c.IndentedJSON(http.StatusBadRequest, "invalid address")
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: usert_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address.0.house_name", Value: address.House},
			{Key: "address.0.street_name", Value: address.Street},
			{Key: "address.0.city", Value: address.City},
			{Key: "address.0.pincode", Value: address.Pincode}}}}

		_, err = userCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, "Something went wrong")
			c.Abort()
			return
		}
		c.IndentedJSON(http.StatusOK, "Updated Successfully")
	}
}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			log.Println("user id is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "invalid user id"})
			c.Abort()
			return
		}

		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, "Something went wrong")
			c.Abort()
			return
		}

		var address models.Address

		if err := c.BindJSON(&address); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
			return
		}

		if address.House == "" || address.Street == "" || address.City == "" || address.Pincode == "" {
			c.IndentedJSON(http.StatusBadRequest, "invalid address")
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: usert_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address.0.house_name", Value: address.House},
			{Key: "address.0.street_name", Value: address.Street},
			{Key: "address.0.city", Value: address.City},
			{Key: "address.0.pincode", Value: address.Pincode}}}}

		_, err = userCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, "Something went wrong")
			c.Abort()
			return
		}
		c.IndentedJSON(http.StatusOK, "Updated Successfully")
	}
}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			log.Println("user id is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "invalid user id"})
			c.Abort()
			return
		}

		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, "Something went wrong")
			c.Abort()
			return
		}

		var address models.Address

		if err := c.BindJSON(&address); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
			return
		}

		if address.House == "" || address.Street == "" || address.City == "" || address.Pincode == "" {
			c.IndentedJSON(http.StatusBadRequest, "invalid address")
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: usert_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address.1.house_name", Value: address.House},
			{Key: "address.1.street_name", Value: address.Street},
			{Key: "address.1.city", Value: address.City},
			{Key: "address.1.pincode", Value: address.Pincode}}}}

		_, err = userCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, "Something went wrong")
			c.Abort()
			return
		}
		c.IndentedJSON(http.StatusOK, "Updated Successfully")
	}

}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {

		user_id := c.Query("id")
		if user_id == "" {
			log.Println("user id is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "invalid user id"})
			c.Abort()
			return
		}

		address := make([]models.Address, 0)
		user_id_hex, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			log.Panic(err)
			c.IndentedJSON(http.StatusInternalServerError, "something went wrong")
			c.Abort()
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{{Key: "_id", Value: user_id_hex}}
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "address", Value: address}}}}

		cursor, err := userCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Panic(err)
			c.IndentedJSON(http.StatusInternalServerError, "something went wrong")
			return
		}

		if cursor.MatchedCount < 1 {
			log.Println(cursor, "matched no documents in AddAddress")
			c.Header("Content-Type", "application/json")
			c.IndentedJSON(http.StatusInternalServerError, "Address not found")
			c.Abort()
			return
		}

		ctx.Done()
		c.IndentedJSON(200, "Successfully deleted the address")

	}
}
