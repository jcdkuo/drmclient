package main

import (
	"drmclient/drm"
	"drmclient/foo"
	"log"
)

func main() {

	foo.Hello()
	go drm.Drm()

	<-drm.WaitChan
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
