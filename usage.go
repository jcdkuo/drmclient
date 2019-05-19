package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
)

// Argument is used for users customization
type Argument struct {
	IPAddress     string
	DRMListenPort int
	SenderIPAddr  string
}

// Usage is used for user customization
func usage(args *Argument) {
	// Describe the arguments
	flag.StringVar(&args.IPAddress, "ip", "0.0.0.0", "The interface you want to use. Maybe you have both wired and wireless interface.")
	flag.IntVar(&args.DRMListenPort, "port", 10000, "DRM service listen port")
	// Parse and check arguments
	parseArguments(args)
	// Show the arguments
	fmt.Println("----------------------------------------------")
	fmt.Printf("Interface IP: %s, Port: %d\n", args.IPAddress, args.DRMListenPort)
	fmt.Println("----------------------------------------------")
}

func parseArguments(args *Argument) {
	flag.Parse()
	if args.IPAddress == "0.0.0.0" {
		args.IPAddress = getLocalIP()
	}
	args.SenderIPAddr = args.IPAddress + ":" + strconv.Itoa(args.DRMListenPort)
}

func getLocalIP() string {
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
