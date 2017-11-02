package ordersinfo

import (
	//json2 "encoding/json"

	"github.com/gin-gonic/gin"
)

type OrderByIdParam struct {
	Id   int    `form:"id" json:"id" binding:"required"`
	Desc string `form:"desc" json:"desc"`
}

type Order struct {
	Id int
	No string
}

func OrdersById(c *gin.Context) {
	var param *OrderByIdParam = new(OrderByIdParam)
	err := c.Bind(param)
	if err != nil {
		c.JSON(501, gin.H{"desc": "error json", "errorCode": 123})
		return
	}
	var orders []Order = []Order{Order{Id: 1, No: "11"}, {Id: 2, No: "22"}}

	if param.Id > 2 {
		orders = append(orders, Order{Id: 3, No: "33appand"})
	}
	//var json, _ = json2.Marshal(orders)
	//c.JSON(200, string(json))
	c.JSON(200, orders)
}
