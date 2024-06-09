package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	First_Name      *string            `json:"first_name" bson:"first_name" validate:"required,min=2,max=30"`
	Last_Name       *string            `json:"last_name" bson:"last_name" validate:"required,min=2,max=30"`
	Email           *string            `json:"email" bson:"email" validate:"email,required"`
	Password        *string            `json:"password" bson:"password" validate:"required,min=6"`
	Phone           *string            `json:"phone" bson:"phone" validate:"required,min=10,max=10,number"`
	Token           *string            `json:"token" bson:"token"`
	Refresh_Token   *string            `json:"refresh_token" bson:"refresh_token"`
	Created_at      time.Time          `json:"created_at"`
	Updated_at      time.Time          `json:"updated_at"`
	User_ID         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"user_cart" bson:"usercart"`
	Address_Details []Address          `json:"address" bson:"address_details"`
	Order_Status    []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	Product_ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *uint8             `json:"rating"`
	Image        *string            `json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Product_Name string             `json:"product_name" bson:"product_name"`
	Price        uint               `json:"price" bson:"price"`
	Rating       *uint              `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
	Quantity     uint               `json:"quantity" bson:"quantity"`
}

type Address struct {
	Address_id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	House      string             `json:"house" bson:"house"`
	Street     string             `json:"street" bson:"street"`
	City       string             `json:"city" bson:"city"`
	Pincode    string             `json:"pincode" bson:"pincode"`
}

type Order struct {
	Order_ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Order_Cart      []ProductUser      `json:"order_list" bson:"order_list"`
	Ordered_At      time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price           *uint              `json:"total_price" bson:"Price"`
	Discount        *int               `json:"discount" bson:"discount"`
	Delivery_Charge *int               `json:"delivery_charge" bson:"delivery_charge"`
	Payment_Method  Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	Mode    string
}
