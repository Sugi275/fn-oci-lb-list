package main

import (
	"fmt"

	"github.com/Sugi275/fn-oci-lb-list/src/ocilib"
)

func main() {
	var lblist = ocilib.GetLBlist()
	fmt.Print(lblist)
}
