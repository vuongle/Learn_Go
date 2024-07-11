package routers

import (
	"github.com/gin-gonic/gin"

	pt "github.com/vuongle/stripe-payment/modules/paypal/transport"
	st "github.com/vuongle/stripe-payment/modules/stripe/transport"
)

func RegisterRouters(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		// create a group apis for "stripe"
		stripe := v1.Group("/stripe")
		{
			stripe.POST("/payment-intent", st.CreatePaymentIntent())
			stripe.GET("/products", st.GetAllProducts())
		}

		// create a group apis for "paypal"
		paypal := v1.Group("/paypal")
		{
			paypal.GET("/keys", pt.GetKeys())
		}
	}
}
