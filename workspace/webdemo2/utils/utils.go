package utils

import (
	"fmt"
	"net"
)

func HostLocalIp() string {
	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	var ips []string
	for _, netInterfaceAddress := range netInterfaceAddresses {

		networkIp, ok := netInterfaceAddress.(*net.IPNet)

		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {

			ip := networkIp.IP.String()

			fmt.Println("Resolved Host IP: " + ip)
			ips = append(ips, ip)
			//return ip
		}
	}
	fmt.Println(ips)
	return ips[0]
}
