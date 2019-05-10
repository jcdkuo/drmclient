package main

import (
	//"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

const Discovery_Req = 0x01
const Discovery_Ack = 0x02

const AttType_FirmwareVersion = 0x01
const AttType_Mac = 0x02
const AttType_IP = 0x03
const AttType_Ext = 0x04
const AttType_RegTimeout = 0x05

const Ext_HTTP = 0x06
const Ext_FTP = 0x07
const Ext_Lang = 0x08
const Ext_MODEL_NAME = 0x09
const Ext_EZ_VER = 0x10
const Ext_HOSTNAME = 0x11

var waitChan chan bool = make(chan bool)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	go drm()

	<-waitChan
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func discovery() []byte {
	buf := new(bytes.Buffer)
	var action uint8 = Discovery_Req
	binary.Write(buf, binary.BigEndian, action)
	var rnd int32 = rand.Int31()
	binary.Write(buf, binary.BigEndian, rnd)

	return buf.Bytes()

}

type Record struct {
	Mac      string
	IP       string
	Model    string
	HttpPort string
}

func (this *Record) Show() {
	fmt.Printf("%s - %16s:%s - %s\n", this.Mac, this.IP, this.HttpPort, this.Model)
}

func drm() {
	udpSock, err := net.ListenUDP("udp", nil)
	if err != nil {
		log.Println(err.Error())
		waitChan <- true
		return
	}

	//send discovery packet
	broadcastAddr := net.UDPAddr{IP: net.IPv4(255, 255, 255, 255), Port: 10000}
	udpSock.WriteToUDP(discovery(), &broadcastAddr)

	//start listening
	buf := make([]byte, 1024)
	for {
		udpSock.SetReadDeadline(time.Now().Add(time.Second * 2))
		readSize, _, err := udpSock.ReadFromUDP(buf)
		if err != nil {
			waitChan <- true
		}

		record := Record{}
		var processedSize int = 0
		if buf[0] == Discovery_Ack {
			r := bytes.NewReader(buf[0:])

			var msgType uint8
			var msgId uint32
			var attType uint8
			var lengthTag uint8
			var contentSize int

			binary.Read(r, binary.BigEndian, &msgType)
			binary.Read(r, binary.BigEndian, &msgId)

			//log.Printf("type:%d id: %d\n", msgType, msgId, attType)
			for processedSize < readSize {
				binary.Read(r, binary.BigEndian, &attType)
				binary.Read(r, binary.BigEndian, &lengthTag)

				processedSize += 6
				tagType := lengthTag >> 7
				processedSize += 1

				if tagType == 0 {
					contentSize = int(lengthTag)
				} else {
					//ignore
					break
				}

				bAtt := make([]byte, contentSize)
				r.Read(bAtt)
				var att string

				switch attType {
				case AttType_FirmwareVersion:
					att = string(bAtt)
				case AttType_Mac:
					att = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", bAtt[0], bAtt[1], bAtt[2], bAtt[3], bAtt[4], bAtt[5])
					att = strings.ToUpper(att)
					record.Mac = att
				case AttType_IP:
					att = fmt.Sprintf("%d.%d.%d.%d", bAtt[0], bAtt[1], bAtt[2], bAtt[3])
					record.IP = att
				case AttType_Ext:
					parseExt(bAtt, &record)
				}

				processedSize += contentSize
			}
		}

		record.Show()
	}
}

func parseExt(buf []byte, record *Record) {
	readSize := len(buf)
	processedSize := 0

	r := bytes.NewReader(buf)

	var attType uint8
	var lengthTag uint8
	var contentSize int

	for processedSize < readSize {
		binary.Read(r, binary.BigEndian, &attType)
		binary.Read(r, binary.BigEndian, &lengthTag)

		processedSize += 2

		tagType := lengthTag >> 7

		if tagType == 0 {
			contentSize = int(lengthTag)
		} else {
			//ignore..
			return
		}

		bAtt := make([]byte, contentSize)
		r.Read(bAtt)
		var att string

		switch attType {
		case Ext_HTTP:
			att = fmt.Sprintf("%d", (bAtt[1]<<8)+bAtt[0])
			record.HttpPort = att
		case Ext_FTP:
			att = fmt.Sprintf("%d", (bAtt[1]<<8)+bAtt[0])
		case Ext_Lang:
			att = string(bAtt)
		case Ext_MODEL_NAME:
			att = string(bAtt)
			record.Model = att
		case Ext_EZ_VER:
			att = fmt.Sprintf("%d.%d.%d.%d", bAtt[0], bAtt[1], bAtt[2], bAtt[3])
		case Ext_HOSTNAME:
			att = string(bAtt)
		}

		processedSize += contentSize
	}
}
