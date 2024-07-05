package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v79"
	"github.com/vuongle/stripe-payment/modules/stripe/transport"
)

func main() {

	// Setup configs
	stripe.Key = os.Getenv("STRIPE_SK")
	fmt.Println(stripe.Key)

	// 1. Start a http server
	r := gin.Default()
	r.Use(cors.Default())
	v1 := r.Group("/v1")
	{
		// create a group apis for "stripe"
		stripe := v1.Group("/stripe")
		{
			stripe.POST("/payment-intent", transport.CreatePaymentIntent())
			// todos.GET("", transport.ListTodoItems(db))
			// todos.GET("/:id", transport.GetTodoItem(db))
			// todos.PATCH("/:id", transport.UpdateTodoItemById(db))
			// todos.DELETE("/:id", transport.DeleteTodoItemById(db))
		}
	}

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
