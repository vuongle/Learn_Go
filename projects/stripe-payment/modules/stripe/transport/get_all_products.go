package transport

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/price"
	"github.com/stripe/stripe-go/v79/product"
	"github.com/vuongle/stripe-payment/common"
	"github.com/vuongle/stripe-payment/modules/stripe/entity"
)

func GetAllProducts() func(*gin.Context) {
	return func(ctx *gin.Context) {

		params := &stripe.ProductListParams{}
		params.Limit = stripe.Int64(100)
		products := product.List(params)

		priceParams := &stripe.PriceListParams{}
		priceParams.Limit = stripe.Int64(100)
		//priceParams.Expand = stripe.StringSlice([]string{"data.product"})
		prices := price.List(priceParams)

		// Iterate over products and prices, match them based on common criteria
		var retProducts = []entity.CustomProduct{}

		for products.Next() {
			product := products.Product()

			for prices.Next() {
				price := prices.Price()

				// Check if the product and price have a common criteria
				// For example, if they have the same product ID
				if product.ID == price.Product.ID {
					// Create a custom product/price struct with the desired fields
					customPrice := entity.CustomPrice{
						ID:         price.ID,
						UnitAmount: price.UnitAmount,
						Currency:   price.Currency,
					}
					customProduct := entity.CustomProduct{
						ID:          product.ID,
						Name:        product.Name,
						Description: product.Description,
						Image:       product.Images[0],
						Price:       customPrice,
					}
					retProducts = append(retProducts, customProduct)

					break
				}
			}

			fmt.Println(product)
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse{Data: retProducts, Error: nil})

	}
}
