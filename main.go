package main

import (
	"drmclient/drm"
	"log"
)

func main() {

	args := Argument{}
	usage(&args)

	waitChan := make(chan bool, 1)
	go drm.Drm(waitChan, args.SenderIPAddr, args.DRMListenPort)
	<-waitChan
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
