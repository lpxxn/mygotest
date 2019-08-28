package main

import (
	"testing"
)

type user2 struct {
	Age  int32
	Name string `json:"name"`
}

func (u *user2) BeforeRender() {
	u.Age += 1
}

var _ BeforeRender = &user2{}

type user3 struct {
	Name string `json:"name"`
}

func BenchmarkCustomerUser(b *testing.B) {
	user := &user2{Name: "li"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		customerRender(user)
	}
	b.StopTimer()
}

func BenchmarkCustomerUser2(b *testing.B) {
	user := user3{Name: "li"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		customerRender(&user)
	}
	b.StopTimer()
}

func BenchmarkCustomerUser3(b *testing.B) {
	_ = user3{Name: "li"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//_ = user.Name
	}
	b.StopTimer()
}
