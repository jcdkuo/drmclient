package main

import (
	"drmclient/drm"
	"drmclient/foo"
	"fmt"
	"log"
	"net"
)

func main() {

	addr, _ := net.InterfaceAddrs()
	fmt.Println(addr)
	interfaces, _ := net.Interfaces()
	fmt.Println(interfaces)

	foo.Hello()
	fmt.Println(drm.Discovery_Req)
	go drm.Drm()

	<-drm.WaitChan
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
