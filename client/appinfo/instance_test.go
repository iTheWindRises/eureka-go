package appinfo

import (
	"fmt"
	"net"
	"testing"
)

func TestDefaultInstanceConfig_IpAddress(t *testing.T) {
	var ip net.IP

	nets, _ := net.Interfaces()
	for _, nt := range nets {
		if nt.Flags != net.FlagLoopback {
			addrs, err := nt.Addrs()
			if err != nil {
				fmt.Println(err)
			}
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				if ip == nil || ip.IsLoopback() {
					continue
				}
				ip = ip.To4()
				if ip == nil {
					continue // not an ipv4 address
				}
			}
		}
	}
}
