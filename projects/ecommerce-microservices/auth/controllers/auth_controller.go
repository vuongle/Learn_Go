package controllers

import (
	"context"
	"ecommerce-microservices/auth/common"
	"ecommerce-microservices/auth/database"
	"ecommerce-microservices/auth/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")

func HashPassword(password string) string {
	return ""
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	return false, ""
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// 1. Binding request to struc
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)
			return
		}

		// 2. check email exists or not
		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
			return
		}
		if count > 0 {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": "user already existed"},
			)
			return
		}

		// 3. Insert into db
		pwd := HashPassword(*user.Password)

		user.Password = &pwd
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserID = user.ID.Hex()

		token, refreshToken, _ = common.GenerateTokens(*user.Email, *user.FirstName, *user.LastName, user.UserID)
		user.Token = &token
		user.RefreshToken = &refreshToken
		user.UserCart = make([]models.ProductUser, 0)
		user.AddressDetails = make([]models.Address, 0)
		user.OrderStatus = make([]models.Order, 0)

		_, insertErr := UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": insertErr.Error()},
			)
			return
		}

		c.JSON(http.StatusCreated, "Successfully signed up")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// 1. Binding request to struc
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()},
			)
			return
		}

		// 2. Find by email
		var founduser models.User
		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": "email is incorrect"},
			)
			return
		}

		// 3. Verify password after found by email
		PasswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
		defer cancel()
		if !PasswordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			log.Println(msg)
			return
		}

		// 4. Update token
		token, refreshToken, _ := common.GenerateTokens(*founduser.Email, *founduser.FirstName, *founduser.LastName, founduser.UserID)
		defer cancel()
		common.UpdateAllTokens(token, refreshToken, founduser.UserID)
		c.JSON(http.StatusFound, founduser)
	}
}
