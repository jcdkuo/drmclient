package drm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

func parseDiscoveryACK(readSize int, buf []byte, record *Record) {

	if buf[0] == discoveryACK {
		r := bytes.NewReader(buf[0:])

		var msgType uint8
		var msgID uint32

		binary.Read(r, binary.BigEndian, &msgType)
		binary.Read(r, binary.BigEndian, &msgID)

		var attrType uint8
		var lengthTag uint8
		var contentSize int
		processedSize := 0

		for processedSize < readSize {
			binary.Read(r, binary.BigEndian, &attrType)
			binary.Read(r, binary.BigEndian, &lengthTag)

			processedSize += 6
			tagType := lengthTag >> 7
			processedSize++

			if tagType == 0 {
				contentSize = int(lengthTag)
			} else {
				//ignore
				break
			}

			bAttr := make([]byte, contentSize)
			r.Read(bAttr)

			var attStr string

			switch attrType {
			case attrFirmwareVersion:
				attStr = string(bAttr)
				record.FirmwareVersion = attStr
			case attrMacAddress:
				attStr = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", bAttr[0], bAttr[1], bAttr[2], bAttr[3], bAttr[4], bAttr[5])
				record.MacAddress = strings.ToUpper(attStr)
			case attrIPAddress:
				record.IPAddress = fmt.Sprintf("%d.%d.%d.%d", bAttr[0], bAttr[1], bAttr[2], bAttr[3])
			case attrExtension:
				parseExtension(bAttr, record)
			}
			processedSize += contentSize
		}
	}
}

func parseExtension(buf []byte, record *Record) {

	readSize := len(buf)
	processedSize := 0

	r := bytes.NewReader(buf)

	var extType uint8
	var lengthTag uint8
	var contentSize int

	for processedSize < readSize {
		binary.Read(r, binary.BigEndian, &extType)
		binary.Read(r, binary.BigEndian, &lengthTag)

		processedSize += 2

		tagType := lengthTag >> 7

		if tagType == 0 {
			contentSize = int(lengthTag)
		} else {
			//ignore..
			return
		}

		bExt := make([]byte, contentSize)
		r.Read(bExt)
		var extStr string

		switch extType {
		case extHTTP:
			extStr = fmt.Sprintf("%d", binary.LittleEndian.Uint16(bExt))
			record.HTTPPort = extStr
		case extHTTPSPort:
			record.HTTPSPort = fmt.Sprintf("%d", binary.LittleEndian.Uint16(bExt))
		case extFTP:
			extStr = fmt.Sprintf("%d", binary.LittleEndian.Uint16(bExt))
		case extLang:
			extStr = string(bExt)
		case extModelName:
			record.ModelName = string(bExt)
		case extEzVersion:
			extStr = fmt.Sprintf("%d.%d.%d.%d", bExt[0], bExt[1], bExt[2], bExt[3])
		case extHostname:
			extStr = string(bExt)
		}
		processedSize += contentSize
	}
}
