package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDnc, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDnc))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDnc, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDnc))
}
