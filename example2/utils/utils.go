package utils

import (
	"net"
	"fmt"
)

func GetIp() string{
	ifaces, _ := net.Interfaces()
	var ip net.IP
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
		}
	}

	return ip.String()
}


func Statistics(ip string, path string) {
	fmt.Println("IP from: ",ip, "Destiny Path: ", path)
}

