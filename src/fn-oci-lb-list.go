package main

import (
	"context"
	"encoding/json"
	"io"

	"github.com/Sugi275/fn-oci-lb-list/src/ocilib"
	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(getLBHundler))
}

func getLBHundler(ctx context.Context, in io.Reader, out io.Writer) {
	var lblist = ocilib.GetLBlist()
	msg := struct {
		Msg []string `json:"message"`
	}{
		Msg: lblist,
	}
	json.NewEncoder(out).Encode(&msg)
}
