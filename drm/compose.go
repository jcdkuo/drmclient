package drm

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"time"
)

func composeDiscoveryREQ() []byte {

	rand.Seed(time.Now().UTC().UnixNano())
	buf := new(bytes.Buffer)

	var action uint8 = discoveryREQ
	binary.Write(buf, binary.BigEndian, action)

	var rnd int32
	rnd = rand.Int31()
	binary.Write(buf, binary.BigEndian, rnd)

	return buf.Bytes()
}
