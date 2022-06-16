package main

import (
	"log"
	"time"

	"github.com/anonutopia/gowaves"
)

type Monitor struct {
	StartedTime int64
}

func (wm *Monitor) start() {
	wm.StartedTime = time.Now().Unix() * 1000
	for {
		// todo - make sure that everything is ok with 100 here
		pages, err := gowaves.WNC.TransactionsAddressLimit(wavesAddress, 100)
		if err != nil {
			log.Println(err)
		}

		if len(pages) > 0 {
			for _, t := range pages[0] {
				wm.checkTransactionWaves(&t)
			}
		}

		pages, err = anc.TransactionsAddressLimit(anoteAddress, 100)
		if err != nil {
			log.Println(err)
		}

		if len(pages) > 0 {
			for _, t := range pages[0] {
				wm.checkTransactionWaves(&t)
			}
		}

		time.Sleep(time.Second * MonitorTick)
	}
}

func (m *Monitor) checkTransactionWaves(talr *gowaves.TransactionsAddressLimitResponse) {
	// log.Printf("%+v\n", talr)
	log.Println(prettyPrint(talr))
}

func (m *Monitor) checkTransactionAnote(talr *gowaves.TransactionsAddressLimitResponse) {
	// log.Printf("%+v\n", talr)
	log.Println(prettyPrint(talr))
}

func initMonitor() {
	m := &Monitor{}
	m.start()
}
