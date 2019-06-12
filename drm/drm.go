package drm

import (
	//"fmt"
	"log"
	"net"
	"time"
)

// Discovery is used for scaning VIVOTEK products
func Discovery(waitChan chan bool, senderIPAddr string, drmListenPort int) {

	addr, err := net.ResolveUDPAddr("udp", senderIPAddr)
	if err != nil {
		log.Println(err.Error())
		//fmt.Println("ResolveUDPAddr")
		waitChan <- true
		return
	}

	udpSock, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Println(err.Error())
		//fmt.Println("ListenUDP")
		waitChan <- true
		return
	}

	//Send discovery REQ packet
	broadcastAddr := net.UDPAddr{IP: net.IPv4bcast, Port: drmListenPort}
	udpSock.WriteToUDP(composeDiscoveryREQ(), &broadcastAddr)

	//start listening
	buf := make([]byte, 1024)

	for {
		udpSock.SetReadDeadline(time.Now().Add(time.Second * 2))
		readSize, _, err := udpSock.ReadFromUDP(buf)
		if err != nil {
			//fmt.Println("ReadFromUDP")
			waitChan <- true
			return
		}

		record := Record{}
		parseDiscoveryACK(readSize, buf, &record)

		checkResult(&record)
	}
}
