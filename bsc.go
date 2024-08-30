package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}

func addWithdraw(addr string, amount uint64) {
	ao := common.HexToAddress(addr)
	am := big.NewInt(int64(amount))

	client, err := ethclient.Dial("wss://bsc-mainnet.core.chainstack.com/b01b27656400b69d4a783761ebc0ca80")
	if err != nil {
		log.Fatal(err)
	}

	auth := getAccountAuth(client, conf.EthKey)
	auth.GasPrice = big.NewInt(3000000000)
	auth.GasLimit = 100000

	tokenAddress := common.HexToAddress("0xe7f0f1585bdbd06b18dbb87099b87bd79bbd315b")
	instance, err := NewMain(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// address := common.HexToAddress("0x78Dd02e309196D8673881C81D6c2261CbB8627c3")
	// bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// name, err := instance.Name(&bind.CallOpts{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// symbol, err := instance.Symbol(&bind.CallOpts{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// decimals, err := instance.Decimals(&bind.CallOpts{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	// fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	// fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

	// fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

	_, err = instance.AddWithdraw(auth, ao, am)
	if err != nil {
		log.Println(err)
		logTelegram(err.Error())
	}

	// fbal := new(big.Float)
	// fbal.SetString(bal.String())
	// value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	// fmt.Printf("balance: %f", value) // "balance: 74605500.647409"

	// a := big.NewInt(1000000000)
	// t, err := instance.AddWithdraw(auth, "blablabla", a)
	// log.Println(err)
	// log.Println(prettyPrint(t))

	// ctx := context.Background()
	// err = client.SendTransaction(ctx, t)
	// log.Println(err)

	// // instance.Mint(auth, address, amount)

	// amount := big.NewInt(9000000000000000000)

	// _, err = instance.Approve(auth, address, amount)
	// log.Println(err)
}

func initBsc() {
	client, err := ethclient.Dial("wss://bsc-mainnet.core.chainstack.com/b01b27656400b69d4a783761ebc0ca80")
	if err != nil {
		log.Fatal(err)
		logTelegram(err.Error())
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
		logTelegram(err.Error())
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
			logTelegram(err.Error())
		case header := <-headers:
			// fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByNumber(context.Background(), header.Number)
			if err != nil {
				log.Println(err)
				logTelegram(err.Error())
			} else {
				for _, t := range block.Transactions() {
					ca := common.HexToAddress("0xe7f0f1585bdbd06b18dbb87099b87bd79bbd315b")
					if t.To() != nil && *t.To() == ca {
						contractABI, err := abi.JSON(strings.NewReader(GetLocalABI("./anote.abi")))
						if err != nil {
							log.Fatal(err)
							logTelegram(err.Error())
						}
						log.Println(prettyPrint(t))
						addr, amount := DecodeTransactionInputData(&contractABI, t.Data())
						log.Println(addr)
						log.Println(amount)

						blockchain := "BSC"

						key := blockchain + Sep + t.Hash().String()
						data, err := getData(key)

						tdb := &Transaction{}
						db.First(t, &Transaction{TxID: t.Hash().String()})

						if err == nil && (data == nil || !data.(bool)) && tdb.ID == 0 && !tdb.Processed {
							if block.Time()*1000 > uint64(mon.StartedTime) {
								addr, amount := DecodeTransactionInputData(&contractABI, t.Data())
								// log.Println(block.Time())
								// log.Println(mon.StartedTime)
								if len(addr) > 0 && amount > 0 && strings.HasPrefix(addr, "3A") {
									err := sendAsset(amount, AnoteId, addr, t.Hash().String())
									if err == nil {
										done := true
										dataTransaction(key, nil, nil, &done)

										tdb.TxID = t.Hash().String()
										tdb.Processed = true
										tdb.Type = blockchain
										db.Save(t)
									}

									chainID, err := client.NetworkID(context.Background())
									if err != nil {
										log.Println(err)
										logTelegram(err.Error())
									}

									// m, err := t.AsMessage(types.NewEIP155Signer(chainID))
									// if err != nil {
									// 	log.Println(err)
									// 	logTelegram(err.Error())
									// }
									// sender := m.From().Hex()

									from, err := types.Sender(types.NewLondonSigner(chainID), t)
									if err != nil {
										fmt.Println(err) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
										logTelegram(err.Error())
									}
									logTelegram(fmt.Sprintf("Gateway: %s %s %.8f", from.Hex(), addr, float64(amount)/float64(SatInBTC)))
								}
							}
						}
					}
				}
			}
		}
	}
}

func GetLocalABI(path string) string {
	abiFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer abiFile.Close()

	result, err := io.ReadAll(abiFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(result)
}

func DecodeTransactionInputData(contractABI *abi.ABI, data []byte) (string, uint64) {
	addr := ""
	amount := uint64(0)
	// The first 4 bytes of the t represent the ID of the method in the ABI
	// https://docs.soliditylang.org/en/v0.5.3/abi-spec.html#function-selector
	methodSigData := data[:4]
	method, err := contractABI.MethodById(methodSigData)
	if err != nil {
		log.Fatal(err)
	}

	inputsSigData := data[4:]
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		log.Fatal(err)
	}

	if method.Name == "deposit" {
		addr = inputsMap["to"].(string)
		a := inputsMap["amount"].(*big.Int)
		amount = a.Uint64()
	}

	return addr, amount
}
