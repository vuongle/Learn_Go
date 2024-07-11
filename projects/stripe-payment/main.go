package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v79"
	"github.com/vuongle/stripe-payment/routers"
)

func main() {

	// Load env and Setup configs
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can not load environment variables")
	}
	stripe.Key = os.Getenv("STRIPE_SK")
	fmt.Println(stripe.Key)
	fmt.Println(os.Getenv("ANDROID_SDK"))

	// 1. Start a http server
	r := gin.Default()
	r.Use(cors.Default())
	routers.RegisterRouters(r)

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
