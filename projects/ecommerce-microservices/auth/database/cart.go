package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("can not find the product")
	ErrCantDecodeProducts = errors.New("can not find products")
	ErrUserIdNotValid     = errors.New("user is not valid")
	ErrCantUpdateUser     = errors.New("can not add product to cart")
	ErrCantRemoveCartItem = errors.New("can not remove item from cart")
	ErrCantGetItem        = errors.New("can not get item from cart")
	ErrCantBuyCartItem    = errors.New("can not update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuy() {

}
