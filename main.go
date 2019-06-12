// +build linux darwin windows
// +build arm amd64

package main

import (
	"drmclient/drm"
	"log"
	"runtime"
)

func main() {

	args := Argument{}
	usage(&args)

	waitChan := make(chan bool, 1)
	if runtime.GOARCH == "arm" {
		go drm.ScanForARM(waitChan, args.SenderIPAddr, args.DRMListenPort)
	} else {
		go drm.Discovery(waitChan, args.SenderIPAddr, args.DRMListenPort)
	}

	<-waitChan
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
