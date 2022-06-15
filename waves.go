package main

import (
	"fmt"
	"log"
	"time"

	"github.com/anonutopia/gowaves"
)

func initWaves() {
	gowaves.WNC.Host = "http://localhost"
	gowaves.WNC.Port = 6869

	a, err := gowaves.WNC.Addresses()
	if err != nil {
		log.Println(err.Error())
		for err != nil {
			time.Sleep(time.Second * 10)
			a, err = gowaves.WNC.Addresses()
			if err != nil {
				log.Println(err.Error())
			}
		}
	}

	ar := *a
	nodeAddress = ar[0]
	fmt.Printf("Node Address: %s\n", nodeAddress)
}
