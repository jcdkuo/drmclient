package drm

import (
	"flag"
	"fmt"
	"net"
)

func Usage() {

	flag.StringVar(&Args.IP, "ip", "0.0.0.0", "The interface you want to use. Maybe you have both wired and wireless interface.")
	flag.IntVar(&Args.Port, "port", 10000, "DRM service listen port")
	flag.Parse()

	if Args.IP == "0.0.0.0" {
		Args.IP = GetLocalIP()
	}

	fmt.Println("----------------------------------------------")
	fmt.Printf("Interface IP: %s, Port: %d\n", Args.IP, Args.Port)
	fmt.Println("----------------------------------------------")
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "0.0.0.0"
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "0.0.0.0"
}
