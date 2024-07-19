package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v79"
	"github.com/vuongle/stripe-payment/middlewares"
	"github.com/vuongle/stripe-payment/routers"
)

func main() {

	/// Load and set environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can not load environment variables")
	}
	stripe.Key = os.Getenv("STRIPE_SK")

	/// Load and setup firebase
	auth, err := middlewares.NewAuthMiddleware("my-flutter-auth-324dd-firebase-adminsdk.json", nil)
	if err != nil {
		panic(err)
	}

	/// Start a http server: middlewares, routers, ...
	r := gin.Default()

	r.Use(cors.Default())
	r.Use(auth.AuthenticationFunc())

	routers.RegisterRouters(r)

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
