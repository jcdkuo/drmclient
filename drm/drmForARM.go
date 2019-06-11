package drm

import (
	//"fmt"
	"log"
	"net"
	"time"
)

// ScanForARM is used for scaning VIVOTEK products
func ScanForARM(waitChan chan bool, senderIPAddr string, drmListenPort int) {

	udpSock, err := net.ListenUDP("udp", nil)
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
