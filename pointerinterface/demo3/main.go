package main



import (
"sync"
"time"
"reflect"
"github.com/gin-gonic/gin"
)

func UpdateCache(ele1,ele2 *int) {
	var lock sync.Mutex
	timer1 := time.NewTicker(time.Second * 5)
	defer timer1.Stop()
	timer2 := time.NewTicker(time.Second * 3)
	defer timer2.Stop()
	for {
		select {
		case <-timer1.C:
			go func(ele *int) {
				lock.Lock()
				defer lock.Unlock()
				*ele += 1
			}(ele1)
		case <-timer2.C:
			go func(ele *int) {
				lock.Lock()
				defer  lock.Unlock()
				*ele += 2
			}(ele2)
		}
	}
}

func cronMiddlewarex(ptr1,ptr2 *int) gin.HandlerFunc {
	return func(c *gin.Context){
		c.Set("ptr1",ptr1)
		c.Set("ptr2",ptr2)
		c.Next()
	}
}

func handler(c *gin.Context) {
	raw_ptr1, _ := c.Get("ptr1")
	prj1 := reflect.ValueOf(raw_ptr1).Elem().Interface()
	raw_ptr2, _ := c.Get("ptr2")
	prj2 := reflect.ValueOf(raw_ptr2).Elem().Interface()
	c.JSON(200,gin.H{"ptr1":prj1,"ptr2":prj2})
}

func main() {
	ele1 := 0
	ele2 := 0

	r := gin.New()
	r.Use(cronMiddlewarex(&ele1,&ele2))
	r.GET("/test",handler)
	go UpdateCache(&ele1,&ele2)
	r.Run(":3333")
}
