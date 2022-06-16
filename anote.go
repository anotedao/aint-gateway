package main

import (
	"fmt"

	"github.com/anonutopia/gowaves"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

func initAnote() *gowaves.WavesNodeClient {
	anc := &gowaves.WavesNodeClient{
		Host: "http://localhost",
		Port: 6869,
	}

	pk, _ := crypto.NewPublicKeyFromBase58(conf.PublicKey)

	a, _ := proto.NewAddressFromPublicKey(55, pk)

	anoteAddress = a.String()

	fmt.Printf("Anote Address: %s\n", anoteAddress)

	return anc
}
