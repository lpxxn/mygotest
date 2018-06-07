package addr

import (
	"net"
	"testing"
	"fmt"
)

func TestExtractor(t *testing.T) {
	testData := []struct {
		addr   string
		expect string
		parse  bool
	}{
		{"127.0.0.1", "127.0.0.1", false},
		{"10.0.0.1", "10.0.0.1", false},
		{"", "", true},
		{"0.0.0.0", "", true},
		{"[::]", "", true},
	}

	for _, d := range testData {
		addr, err := Extract(d.addr)
		fmt.Println("addr ----", addr)
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		if d.parse {
			ip := net.ParseIP(addr)
			fmt.Println(ip)
			if ip == nil {
				t.Error("Unexpected nil IP")
			}

		} else if addr != d.expect {
			t.Errorf("Expected %s got %s", d.expect, addr)
		}
	}

}
