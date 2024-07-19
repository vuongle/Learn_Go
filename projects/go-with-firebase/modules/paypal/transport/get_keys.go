package transport

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vuongle/go-with-firebase/common"
)

func GetKeys() func(*gin.Context) {
	return func(ctx *gin.Context) {
		clientId := os.Getenv("PAYPAL_CLIENT_ID")
		clientSecret := os.Getenv("PAYPAL_SK")
		fmt.Println(clientId)
		fmt.Println(clientSecret)

		if clientId == "" || clientSecret == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "PayPal credentials not found"})
			return
		}

		result := gin.H{"clientId": clientId, "clientSecret": clientSecret}
		ctx.JSON(http.StatusOK, common.SuccessResponse{Data: result, Error: nil})
	}
}
