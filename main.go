package main

import (
	"log"

	"github.com/anonutopia/gowaves"
)

var conf *Config

var anoteAddress string

var wavesAddress string

var anc *gowaves.WavesNodeClient

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	initWaves()

	anc = initAnote()

	initMonitor()
}
