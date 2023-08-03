package main

import (
	"log"

	"github.com/anonutopia/gowaves"
	"gorm.io/gorm"
)

var conf *Config

var anoteAddress string

var wavesAddress string

var anc *gowaves.WavesNodeClient

var db *gorm.DB

var mon *Monitor

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	db = initDb()

	initWaves()

	anc = initAnote()

	mon = initMonitor()

	initBsc()

	// addWithdraw("0x78Dd02e309196D8673881C81D6c2261CbB8627c3", 10000)
}
