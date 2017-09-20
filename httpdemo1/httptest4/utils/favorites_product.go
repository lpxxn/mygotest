package utils


type FavoriteProduct struct {
	ProductName string	`json:"product_name"`
	ProductCode string `json:"product_code"`
	FavoritePrice float32	`json:"favorite_price"`
}

type FavoriteProducts [] FavoriteProduct