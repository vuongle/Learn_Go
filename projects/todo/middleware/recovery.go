package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Recovery() func(*gin.Context) {
	return func(ctx *gin.Context) {
		log.Println("Recovery")
		ctx.Next()
	}
}
