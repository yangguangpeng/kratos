package main

import (
	"fmt"
)

func main() {

	h := fmt.Sprintf("%x", [16]byte{123})
	fmt.Println(h)
	return

}
