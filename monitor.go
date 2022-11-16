package main

import (
	"context"
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

		t := &Transaction{}
		db.First(t, &Transaction{TxID: talr.ID})

		if err == nil && (data == nil || !data.(bool)) && t.ID == 0 && !t.Processed {
			done := true
			dataTransaction(key, nil, nil, &done)

			t.TxID = talr.ID
			t.Processed = true
			t.Type = blockchain
			db.Save(t)

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

	if talr.Sender != "3AJj8zKHiitrfG6FGLzPcBCgeHNQzNASToz" || recAddress != "3P4SXwzKohZmj4w8gvwdBab5u9dQothxKXd" {
		sendAsset(uint64(talr.Amount), assetId, recAddress, talr.Sender)
		log.Printf("Sent: %d %s %s %s\n", uint64(talr.Amount), assetId, recAddress, talr.Sender)
	} else {
		log.Printf("Caught: %d %s %s %s\n", uint64(talr.Amount), assetId, recAddress, talr.Sender)
		logTelegram("Caught a scumbag.")
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

func initMonitor() {
	m := &Monitor{}
	go m.loadTransactions()
	m.start()
}
