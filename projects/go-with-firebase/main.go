package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v79"
	ginfirebaseauth "github.com/vuongle/go-with-firebase/middlewares"
	"github.com/vuongle/go-with-firebase/routers"
)

func main() {

	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can not load environment variables")
	}
	stripe.Key = os.Getenv("STRIPE_SK")

	// 1. Start a http server
	r := gin.Default()
	r.StaticFS("/static", http.Dir("static")) // Serving static files
	r.Use(cors.Default())
	routers.RegisterRouters(r)

	// Set up firebase app and Auth client
	middleware, err := ginfirebaseauth.New("my-flutter-auth-324dd-firebase-adminsdk.json", nil)
	if err != nil {
		panic(err)
	}
	auth := r.Group("/premium")
	auth.Use(middleware.MiddlewareFunc())
	auth.GET("/", func(c *gin.Context) {
		claims := ginfirebaseauth.ExtractClaims(c)
		fmt.Println(claims)
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8000/static/premium.html")
	})

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
