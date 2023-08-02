package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	// ao := common.HexToAddress(addr)
	// am := big.NewInt(int64(amount))

	client, err := ethclient.Dial("https://bsc.meowrpc.com")
	if err != nil {
		log.Fatal(err)
	}

	auth := getAccountAuth(client, conf.EthKey)

	tokenAddress := common.HexToAddress("0xbad04e33cc88bbcccc1b7adb8319f7d36f5bc472")
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

	// _, err = instance.AddWithdraw(auth, ao, am)
	// log.Println(err)

	// fbal := new(big.Float)
	// fbal.SetString(bal.String())
	// value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	// fmt.Printf("balance: %f", value) // "balance: 74605500.647409"

	a := big.NewInt(1000000000)
	t, err := instance.Deposit(auth, "blablabla", a)
	log.Println(err)
	log.Println(prettyPrint(t))

	ctx := context.Background()
	err = client.SendTransaction(ctx, t)
	log.Println(err)

	// // instance.Mint(auth, address, amount)

	// amount := big.NewInt(9000000000000000000)

	// _, err = instance.Approve(auth, address, amount)
	// log.Println(err)
}
