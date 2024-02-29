package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/anonutopia/gowaves"
	"github.com/mr-tron/base58"
	"github.com/wavesplatform/gowaves/pkg/client"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

type Monitor struct {
	StartedTime int64
}

func (m *Monitor) start() {
	m.StartedTime = time.Now().Unix() * 1000
	for {
		// todo - make sure that everything is ok with 100 here
		// gowaves.WNC.Host = WavesNodeURL
		// pages, err := gowaves.WNC.TransactionsAddressLimit(wavesAddress, 100)
		// if err != nil {
		// 	log.Println(err)
		// 	logTelegram(err.Error())
		// }

		// if len(pages) > 0 {
		// 	for _, t := range pages[0] {
		// 		m.checkTransaction(&t, TypeWaves)
		// 	}
		// }

		pages, err := anc.TransactionsAddressLimit(anoteAddress, 100)
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
	if talr.Recipient == anoteAddress && talr.AssetID == AnoteId {
		key := blockchain + Sep + talr.ID
		data, err := getData(key)

		t := &Transaction{}
		db.First(t, &Transaction{TxID: talr.ID})

		if err == nil && (data == nil || !data.(bool)) && t.ID == 0 && !t.Processed {
			done := true
			dataTransaction(key, nil, nil, &done)

			t.TxID = talr.ID
			t.Processed = true
			t.Type = blockchain
			db.Save(t)

			if len(talr.Attachment) > 0 && talr.Timestamp > m.StartedTime {
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
		} else if talr.AssetID == AnotePrealphaWavesId {
			assetId = ""
			talr.Amount /= 1000
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

	if talr.Sender == "3AH265emjtkeK3wYLyHSP3HC1sV28zXqMCP" ||
		talr.Sender == "3AJj8zKHiitrfG6FGLzPcBCgeHNQzNASToz" ||
		talr.Sender == "3AT5e7vq8DCby4ooqhf3biw59WBRMFWhB1F" ||
		talr.Sender == "3AQKvpt1cMX7KyL1wLkGryAC5xE9kQyhqUF" ||
		talr.Sender == "3AM9n979R9ttNFvNpENBqesgX13GKGsJ2Cb" ||
		talr.Sender == "3AMKhRz9gy2oAfAiztSFk4LymtmiyKvSUui" ||
		talr.Sender == "3AVUJPDn9vvDxA9JtauyVX312Ykp5g2txsb" ||
		recAddress == "3P4SXwzKohZmj4w8gvwdBab5u9dQothxKXd" ||
		recAddress == "3PMTF844fus4LS8w2TkNRGu3Jf5SadLFcAx" ||
		recAddress == "3PGCBwomcSSHTsugmrQiMqPHSLgab19JNas" ||
		recAddress == "3PHBLpQmx3jrMxqixYsXZ22F3CVGD8T63bj" ||
		recAddress == "3PDYrgoT1sS84PwutQwF4Tkuc7W1acvHoXq" {

		log.Printf("Caught: %d %s %s %s\n", uint64(talr.Amount), assetId, recAddress, talr.Sender)
		logTelegram("Caught a scumbag.")
	} else {
		recAddress = strings.ReplaceAll(strings.ReplaceAll(recAddress, "\n", ""), " ", "")
		if strings.HasPrefix(recAddress, "0x") {
			amount := uint64(talr.Amount - 10000000)
			if amount > 2*SatInBTC {
				amount = 2 * SatInBTC
			}
			addWithdraw(recAddress, amount)
		} else {
			sendAsset(uint64(talr.Amount), assetId, recAddress, talr.Sender)
		}

		if assetId == AintAnoteId || assetId == AintWavesId {
			logTelegram(fmt.Sprintf("Gateway AINT: %s %s %.8f", talr.Sender, recAddress, float64(talr.Amount)/float64(SatInBTC)))
		} else {
			logTelegram(fmt.Sprintf("Gateway: %s %s %.8f", talr.Sender, recAddress, float64(talr.Amount)/float64(SatInBTC)))
		}

		log.Printf("Sent: %d %s %s %s\n", uint64(talr.Amount), assetId, recAddress, talr.Sender)
	}
}

func (m *Monitor) loadTransactions() {
	cl, err := client.NewClient(client.Options{BaseUrl: AnoteNodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}

	// Context to cancel the request execution on timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	sender, err := crypto.NewPublicKeyFromBase58(conf.PublicKey)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}

	addr, err := proto.NewAddressFromPublicKey(55, sender)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}

	data, _, err := cl.Addresses.AddressesData(ctx, addr)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}
	for _, de := range data {
		txid := strings.Split(de.GetKey(), Sep)[1]
		blockchain := strings.Split(de.GetKey(), Sep)[0]
		t := &Transaction{TxID: txid}
		db.FirstOrCreate(t, t)
		t.Type = blockchain
		t.Processed = de.ToProtobuf().GetBoolValue()
		db.Save(t)
	}
}

func initMonitor() *Monitor {
	m := &Monitor{}
	go m.loadTransactions()
	go m.start()

	return m
}
