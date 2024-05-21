package routes

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(incomingRoutes *gin.Engine) {
	authRoutes(incomingRoutes)
	productRoutes(incomingRoutes)
	cartRoutes(incomingRoutes)
}

func authRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/auth/signup")
	incomingRoutes.POST("/auth/login")
}

func productRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/products/add-product")
	incomingRoutes.GET("/products/view-product")
	incomingRoutes.GET("/products/search")
}

func cartRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/cart/add-to-cart")
	incomingRoutes.GET("/cart/remove-item")
	incomingRoutes.GET("/cart/checkout")
	incomingRoutes.GET("/cart/instant-buy")
}
