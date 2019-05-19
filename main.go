package main

import (
	"drmclient/drm"
	"log"
)

func main() {

	drm.Usage()
	go drm.Drm()

	<-drm.WaitChan
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
