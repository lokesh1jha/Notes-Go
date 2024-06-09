package database

import (
	"errors"
)

var (
	ErrCantFindProduct        = errors.New("cannot find product")
	ErrCantDecodeProducts     = errors.New("cannot find the product")
	ErrUserIDIsNotValid       = errors.New("user id is not valid")
	ErrCantUpdateUser         = errors.New("cannot update user")
	ErrCantRemoveItemFromCart = errors.New("cannot remove item from cart")
	ErrCantGetItem            = errors.New("cannot get item from the cart")
	ErrCantBuyCartItem        = errors.New("cannot update the purchase")
)

func AddProductToCart() {

}

func RemoveProductFromCart() {
}

func BuyFromCart() {

}

func InstantBuy() {
}

func GetItemFromCart() {
}

func EmptyCart() {
}
