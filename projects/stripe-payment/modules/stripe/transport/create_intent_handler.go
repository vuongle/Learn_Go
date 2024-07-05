package transport

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"

	"github.com/vuongle/stripe-payment/modules/stripe/entity"
)

func CreatePaymentIntent() func(*gin.Context) {
	return func(ctx *gin.Context) {

		fmt.Println("CreatePaymentIntent")

		// Parse request from client
		var data entity.IntentCreationBody
		err := ctx.ShouldBind(&data) // pass pointer of data (not pass data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid request payload"))
			return
		}

		// params := &stripe.PaymentIntentParams{
		// 	Amount:   stripe.Int64(50),
		// 	Currency: stripe.String(string(stripe.CurrencyUSD)),
		// 	AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
		// 		Enabled: stripe.Bool(true),
		// 	},
		// }

		params := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(data.Amount),
			Currency: stripe.String(string(data.Currency)),
			AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
				Enabled: stripe.Bool(true),
			},
		}
		result, err := paymentintent.New(params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		type intentResponse struct {
			ClientSecret string `json:"client-secret"`
		}
		type successResponse struct {
			Data  interface{} `json:"data"` //interface{}: means any data type
			Error error       `json:"error"`
		}
		ctx.JSON(http.StatusOK, successResponse{Data: intentResponse{ClientSecret: result.ClientSecret}, Error: nil})
	}
}
