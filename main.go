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

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	db = initDb()

	initWaves()

	anc = initAnote()

	initMonitor()
}
