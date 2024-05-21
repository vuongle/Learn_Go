package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName      *string            `json:"first_name" validate:"required,min=2,max=2"`
	LastName       *string            `json:"last_name"`
	Password       *string            `json:"password"`
	Email          *string            `json:"email" validate:"email,required"`
	Phone          *string            `json:"phone"`
	Token          *string            `json:"token"`
	RefreshToken   *string            `json:"refresh_token"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	UserID         string             `json:"user_id"`
	UserCart       []ProductUser      `json:"user_cart"`
	AddressDetails []Address          `json:"address" bson:"address"`
	OrderStatus    []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	ProductID   primitive.ObjectID `bson:"_id"`
	ProductName *string            `json:"product_name"`
	Price       *uint64            `json:"price"`
	Image       *string            `json:"image"`
	Rating      *uint8             `json:"rating"`
}

type ProductUser struct {
	ProductID   primitive.ObjectID `bson:"_id"`
	ProductName *string            `json:"product_name" bson:"product_name"`
	Price       *uint64            `json:"price" bson:"price"`
	Image       *string            `json:"image" bson:"image"`
	Rating      *uint8             `json:"rating" bson:"rating"`
}

type Address struct {
	AddressID primitive.ObjectID `bson:"_id"`
	Street    *string            `json:"street" bson:"street"`
	City      *string            `json:"city" bson:"city"`
}

type Order struct {
	OrderID       primitive.ObjectID `bson:"_id"`
	OrderCart     []ProductUser      `json:"order_list" bson:"order_list"`
	OrderedAt     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price         uint64             `json:"total_price" bson:"total_price"`
	Discount      *uint64            `json:"discount" bson:"discount"`
	PaymentMethod Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}
