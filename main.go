package main

import (
	"drmclient/drm"
	"drmclient/foo"
	"fmt"
	"log"
)

func main() {

	foo.Hello()
	fmt.Println(drm.Discovery_Req)
	go drm.Drm()

	<-drm.WaitChan
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
