package drm

import (
	"fmt"
)

func (this *Record) Show() {

	fmt.Printf("%s | %15s:%s | %15s | Firmware: %s | HTTPS:%s\n",
		this.Mac, this.IP, this.HttpPort, this.Model, this.Firmware_version, this.HttpsPort)
}

func checkResult(record *Record) {
	if value, ok := Records[record.Mac]; ok {
		Records[record.Mac] = value
	} else {
		Records[record.Mac] = record.Model
		record.Show()
	}
}
