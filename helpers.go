package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/wavesplatform/gowaves/pkg/client"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

func dataTransaction(key string, valueStr *string, valueInt *int64) error {
	// Create sender's public key from BASE58 string
	sender, err := crypto.NewPublicKeyFromBase58(conf.PublicKey)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create sender's private key from BASE58 string
	sk, err := crypto.NewSecretKeyFromBase58(conf.PrivateKey)
	if err != nil {
		log.Println(err)
		return err
	}

	// Current time in milliseconds
	ts := time.Now().Unix() * 1000

	tr := proto.NewUnsignedData(2, sender, Fee, uint64(ts))

	if valueStr == nil && valueInt == nil {
		tr.Entries = append(tr.Entries,
			&proto.DeleteDataEntry{
				Key: key,
			},
		)
	} else if valueInt == nil {
		tr.Entries = append(tr.Entries,
			&proto.StringDataEntry{
				Key:   key,
				Value: *valueStr,
			},
		)
	} else if valueStr == nil {
		tr.Entries = append(tr.Entries,
			&proto.IntegerDataEntry{
				Key:   key,
				Value: *valueInt,
			},
		)
	}

	err = tr.Sign(55, sk)
	if err != nil {
		log.Println(err)
		return err
	}

	// Create new HTTP client to send the transaction to public TestNet nodes
	client, err := client.NewClient(client.Options{BaseUrl: AnoteNodeURL, Client: &http.Client{}})
	if err != nil {
		log.Println(err)
		return err
	}

	// Context to cancel the request execution on timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// // Send the transaction to the network
	_, err = client.Transactions.Broadcast(ctx, tr)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func getData(key string) interface{} {
	dkr, err := anc.AddressesDataKey(anoteAddress, key)
	if err != nil {
		log.Println(err.Error())
	}
	return dkr.Value
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
