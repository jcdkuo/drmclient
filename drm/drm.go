package drm

import (
	"bytes"
	"encoding/binary"
	//"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func discovery() []byte {

	rand.Seed(time.Now().UTC().UnixNano())
	buf := new(bytes.Buffer)

	var action uint8 = DISCOVERY_REQ
	binary.Write(buf, binary.BigEndian, action)

	var rnd int32 = rand.Int31()
	binary.Write(buf, binary.BigEndian, rnd)

	return buf.Bytes()
}

func Drm() {

	senderAddr := Args.IP + ":" + strconv.Itoa(Args.Port)
	addr, err := net.ResolveUDPAddr("udp", senderAddr)
	if err != nil {
		log.Println(err.Error())
		//fmt.Println("ResolveUDPAddr")
		WaitChan <- true
		return
	}

	udpSock, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Println(err.Error())
		//fmt.Println("ListenUDP")
		WaitChan <- true
		return
	}

	//send discovery packet
	broadcastAddr := net.UDPAddr{IP: net.IPv4bcast, Port: Args.Port}
	udpSock.WriteToUDP(discovery(), &broadcastAddr)

	//start listening
	buf := make([]byte, 1024)

	for {
		udpSock.SetReadDeadline(time.Now().Add(time.Second * 2))
		readSize, _, err := udpSock.ReadFromUDP(buf)
		if err != nil {
			//fmt.Println("ReadFromUDP")
			WaitChan <- true
			return
		}

		record := Record{}
		parseAtt(readSize, buf, &record)

		checkResult(&record)
	}
}
