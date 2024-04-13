package core

import (
	"context"
	"yumtrip/models"
)

// GetCoupons is a function that returns a list of coupons
func GetCoupons(context context.Context, skip, limit int) ([]models.Coupon, error) {
	coupons, err := models.GetCoupons(context)
	if err != nil {
		return nil, err
	}
	return coupons, nil
}