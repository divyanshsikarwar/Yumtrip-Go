package models

type Coupon struct {
	ID          string `json:"_id"`
	Code        string `json:"code"`
	Discount    string `json:"discount"` // percentage
	IsActive   bool `json:"isActive"`
	TimesUsed   int `json:"timesUsed"`
	MaxUses     int `json:"maxUses"`
	AvailableFrom string `json:"availableFrom"`
	Expiration  string `json:"expiration"`
}

func (c *Coupon) GetID() string {
	return c.ID
}

func (c *Coupon) Create() error {
	// c.create()
	return nil
}

func GetCoupons() ([]Coupon, error) {
	// Get Coupons from MongoDB
	return nil, nil
}