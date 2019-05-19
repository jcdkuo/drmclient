package drm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

func parseAtt(readSize int, buf []byte, record *Record) {

	var processedSize int = 0
	if buf[0] == DISCOVERY_ACK {
		r := bytes.NewReader(buf[0:])

		var msgType uint8
		var msgId uint32
		var attrType uint8
		var lengthTag uint8
		var contentSize int

		binary.Read(r, binary.BigEndian, &msgType)
		binary.Read(r, binary.BigEndian, &msgId)

		//log.Printf("type:%d id: %d\n", msgType, msgId, attType)
		for processedSize < readSize {
			binary.Read(r, binary.BigEndian, &attrType)
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

			switch attrType {
			case ATTR_FIRMEWARE_VERSION:
				att = string(bAtt)
				record.Firmware_version = att
			case ATTR_MAC:
				att = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", bAtt[0], bAtt[1], bAtt[2], bAtt[3], bAtt[4], bAtt[5])
				att = strings.ToUpper(att)
				record.Mac = att
			case ATTR_IP:
				att = fmt.Sprintf("%d.%d.%d.%d", bAtt[0], bAtt[1], bAtt[2], bAtt[3])
				record.IP = att
			case ATTR_EXT:
				parseExt(bAtt, record)
			}

			processedSize += contentSize
		}
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
		case EXT_HTTP:
			att = fmt.Sprintf("%d", (bAtt[1]<<8)+bAtt[0])
			record.HttpPort = att
		case EXT_HTTPS_PORT:
			att = fmt.Sprintf("%d", (bAtt[1]<<8)+bAtt[0])
			record.HttpsPort = att
		case EXT_FTP:
			att = fmt.Sprintf("%d", (bAtt[1]<<8)+bAtt[0])
		case EXT_LANG:
			att = string(bAtt)
		case EXT_MODEL_NAME:
			att = string(bAtt)
			record.Model = att
		case EXT_EZ_VER:
			att = fmt.Sprintf("%d.%d.%d.%d", bAtt[0], bAtt[1], bAtt[2], bAtt[3])
		case EXT_HOSTNAME:
			att = string(bAtt)
		}

		processedSize += contentSize
	}
}
