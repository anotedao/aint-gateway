package main

import (
	"fmt"

	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

func initWaves() {
	pk, _ := crypto.NewPublicKeyFromBase58(conf.PublicKey)

	a, _ := proto.NewAddressFromPublicKey(proto.MainNetScheme, pk)

	wavesAddress = a.String()

	fmt.Printf("Waves Address: %s\n", wavesAddress)
}
