package models

import (
	"time"
)

type JdFavoriteProduct struct {
	ProductName string	`json:"product_name"`
	ProductCode string `json:"product_code"`
	FavoritePrice float32	`json:"favorite_price"`
	SendEmailTime time.Time
	SendCount int
}

type JdFavoriteProducts []JdFavoriteProduct

type JdInfo struct {
	Favorite_Products JdFavoriteProducts `json:"jd_favorite_products"`
	PriceUrl string	`json:"price_url"`
}

