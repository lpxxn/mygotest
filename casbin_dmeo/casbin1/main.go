package main

import (
	"fmt"

	"github.com/casbin/casbin"
)

func main() {
	e, err := casbin.NewEnforcer("rbac_with_domains_model.conf", "rbac_with_domains_policy.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(e)
}
