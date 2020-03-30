package common

import (
	"net"

	st "github.com/clonerOpsTool/pkg/structs"
)

// IsHost checks if the device running the application is part of the scanned
func IsHost(s st.Server) (bool, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return false, err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return false, err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip.String() == s.LocalIP {
				// fmt.Printf("LocalIP match: %s vs %s.\n", ip, s.LocalIP)
				return true, nil
			} else if ip.String() == s.VpnIP {
				// fmt.Printf("VpnIP match: %s vs %s.\n", ip, s.VpnIP)
				return true, nil
			}
		}
	}
	return false, nil
}
