package main

import (
	"errors"
	"fmt"
)

type Cart struct {
	Total float64 // in dollars
}

func ApplyCoupon(cart *Cart, couponCode string) error {
	if cart.Total <= 400 {
		return errors.New("coupon cannot be applied: cart total must be above $400")
	}

	// For demo purposes: flat 10% discount if code is "DISCOUNT10"
	if couponCode == "DISCOUNT10" {
		discount := cart.Total * 0.10
		cart.Total = cart.Total - discount
		return nil
	}

	return errors.New("invalid coupon code")
}

func main() {
	cart := &Cart{Total: 1050} // $450 cart

	err := ApplyCoupon(cart, "DISCOUNT10")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Final cart total:", cart.Total)
}
