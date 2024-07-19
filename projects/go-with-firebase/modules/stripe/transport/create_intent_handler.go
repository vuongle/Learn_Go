package transport

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"

	"github.com/vuongle/go-with-firebase/common"
	"github.com/vuongle/go-with-firebase/modules/stripe/entity"
)

func CreatePaymentIntent() func(*gin.Context) {
	return func(ctx *gin.Context) {

		// Parse request from client
		var data entity.IntentCreationBody
		err := ctx.ShouldBind(&data) // pass pointer of data (not pass data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("invalid request payload"))
			return
		}

		// Create params for create an intent
		params := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(data.Amount),
			Currency: stripe.String(string(data.Currency)),
			AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
				Enabled: stripe.Bool(true),
			},
		}

		// Create a new intent
		result, err := paymentintent.New(params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse{Data: common.IntentResponse{ClientSecret: result.ClientSecret}, Error: nil})
	}
}
