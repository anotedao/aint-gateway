package main

import (
	"log"

	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

var conf *Config

var nodeAddress string

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	initWaves()

	pk, _ := crypto.NewPublicKeyFromBase58(conf.PublicKey)

	a, _ := proto.NewAddressFromPublicKey(55, pk)

	log.Println(a.String())

	mess := "%s%s__3AP5zrscjAo7Jb8urXd6836XRTGakZf7Z9B__178.62.103.70"

	err := dataTransaction("3ASLefwuE3dz9cW9bhP6ZC9N73pLqs2vPEH", &mess, nil)
	if err != nil {
		log.Println(err)
	}
}
