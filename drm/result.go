package drm

import (
	"fmt"
)

func (r *Record) show() {

	fmt.Printf("%s | %15s:%s | %15s | Firmware: %s | HTTPS:%s\n",
		r.MacAddress, r.IPAddress, r.HTTPPort, r.ModelName, r.FirmwareVersion, r.HTTPSPort)
}

func checkResult(record *Record) {
	if value, exist := Records[record.MacAddress]; exist {
		Records[record.MacAddress] = value
	} else {
		Records[record.MacAddress] = record.ModelName
		record.show()
	}
}
