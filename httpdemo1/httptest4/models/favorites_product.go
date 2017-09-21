package models

import "google.golang.org/genproto/googleapis/type/date"

type JdFavoriteProduct struct {
	ProductName string	`json:"product_name"`
	ProductCode string `json:"product_code"`
	FavoritePrice float32	`json:"favorite_price"`
	SendEmailTime date.Date
	SendCount int
}

type JdFavoriteProducts []JdFavoriteProduct

type JdInfo struct {
	Favorite_Products JdFavoriteProducts `json:"jd_favorite_products"`
	PriceUrl string	`json:"price_url"`
}

