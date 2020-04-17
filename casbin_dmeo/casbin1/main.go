package main

import (
	"fmt"

	"github.com/casbin/casbin"
)

func main() {
	// 能正确读取，但是 r2 p2 g2都不能用，源代码里是写死的
	e, err := casbin.NewEnforcer("rbac_with_domains_model.conf", "rbac_with_domains_policy.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(e)

	res, err := e.Enforce("alice", "domain1", "data1", "read")
	if err != nil {
		panic(err)
	}
	// 源代码里写的太死，所以这个也就是false了, 完蛋
	res, err = e.Enforce("alif", "p2_admin", "data2", "write")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

/*
	testDomainEnforce(t, e, "alice", "domain1", "data1", "read", true)
	testDomainEnforce(t, e, "alice", "domain1", "data1", "write", true)
	testDomainEnforce(t, e, "alice", "domain1", "data2", "read", false)
	testDomainEnforce(t, e, "alice", "domain1", "data2", "write", false)
	testDomainEnforce(t, e, "bob", "domain2", "data1", "read", false)
	testDomainEnforce(t, e, "bob", "domain2", "data1", "write", false)
	testDomainEnforce(t, e, "bob", "domain2", "data2", "read", true)
	testDomainEnforce(t, e, "bob", "domain2", "data2", "write", true)

func testDomainEnforce(t *testing.T, e *Enforcer, sub string, dom string, obj string, act string, res bool) {
	t.Helper()
	if myRes, _ := e.Enforce(sub, dom, obj, act); myRes != res {
		t.Errorf("%s, %s, %s, %s: %t, supposed to be %t", sub, dom, obj, act, myRes, res)
	}
}

*/
