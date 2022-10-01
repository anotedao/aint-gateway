package main

import (
	"log"
	"time"

	"github.com/anonutopia/gowaves"
	"github.com/mr-tron/base58"
)

type Monitor struct {
	StartedTime int64
}

func (m *Monitor) start() {
	m.StartedTime = time.Now().Unix() * 1000
	for {
		// todo - make sure that everything is ok with 100 here
		pages, err := gowaves.WNC.TransactionsAddressLimit(wavesAddress, 100)
		if err != nil {
			log.Println(err)
			logTelegram(err.Error())
		}

		if len(pages) > 0 {
			for _, t := range pages[0] {
				m.checkTransaction(&t, TypeWaves)
			}
		}

		pages, err = anc.TransactionsAddressLimit(anoteAddress, 100)
		if err != nil {
			log.Println(err)
			logTelegram(err.Error())
		}

		if len(pages) > 0 {
			for _, t := range pages[0] {
				m.checkTransaction(&t, TypeAnote)
			}
		}

		time.Sleep(time.Second * MonitorTick)
	}
}

func (m *Monitor) checkTransaction(talr *gowaves.TransactionsAddressLimitResponse, blockchain string) {
	if (talr.Recipient == wavesAddress && talr.AssetID == AnoteWavesId) ||
		(talr.Recipient == anoteAddress && talr.AssetID == "") ||
		(talr.Recipient == wavesAddress && talr.AssetID == AintWavesId) ||
		(talr.Recipient == anoteAddress && talr.AssetID == AintAnoteId) {
		key := blockchain + Sep + talr.ID
		data, err := getData(key)

		if err == nil && (data == nil || !data.(bool)) {
			done := true
			dataTransaction(key, nil, nil, &done)

			if len(talr.Attachment) > 0 {
				m.processTransaction(talr, blockchain)
			}
		}
	}
}

func (m *Monitor) processTransaction(talr *gowaves.TransactionsAddressLimitResponse, blockchain string) {
	var assetId string

	recipient, err := base58.Decode(talr.Attachment)
	if err != nil {
		log.Println(err.Error())
		logTelegram(err.Error())
	}
	recAddress := string(recipient[:])

	if blockchain == TypeWaves {
		if talr.AssetID == AnoteWavesId {
			assetId = ""
		} else {
			assetId = AintAnoteId
		}
	} else {
		if talr.AssetID == "" {
			assetId = AnoteWavesId
		} else {
			assetId = AintWavesId
		}
	}

	sendAsset(uint64(talr.Amount), assetId, recAddress, talr.Sender)

	log.Println("Sent.")
}

func initMonitor() {
	m := &Monitor{}
	m.start()
}
