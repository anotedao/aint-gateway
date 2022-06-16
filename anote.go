package main

import (
	"fmt"
	"log"
	"time"

	"github.com/anonutopia/gowaves"
)

func initAnote() *gowaves.WavesNodeClient {
	anc := &gowaves.WavesNodeClient{
		Host: "http://localhost",
		Port: 6869,
	}

	a, err := anc.Addresses()
	if err != nil {
		log.Println(err.Error())
		for err != nil {
			time.Sleep(time.Second * 10)
			a, err = anc.Addresses()
			if err != nil {
				log.Println(err.Error())
			}
		}
	}

	ar := *a
	anoteAddress = ar[0]

	fmt.Printf("Anote Address: %s\n", anoteAddress)

	return anc
}
