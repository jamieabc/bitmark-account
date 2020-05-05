package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/bitmark-inc/bitmark-sdk-go"
	"github.com/bitmark-inc/bitmark-sdk-go/account"
)

var seed string

func init() {
	flag.StringVar(&seed, "s", "", "account seed")
	flag.Parse()
}

func main() {
	sdk.Init(&sdk.Config{Network: sdk.Testnet})

	var acc account.Account
	var err error
	if seed != "" {
		acc, err = account.FromSeed(seed)
		if nil != err {
			fmt.Printf("account recover from seed %s with error: %s\n", seed, err)
			return
		}
	} else {
		acc, err = account.New()
		if nil != err {
			fmt.Printf("new account with error: %s\n", err)
			return
		}
	}

	timestamp := strconv.Itoa(int(time.Now().Unix() * 1000))
	signature := acc.Sign([]byte(timestamp))
	encKey := hex.EncodeToString(acc.(*account.AccountV2).EncrKey.PublicKeyBytes())

	fmt.Println("account: ", acc.AccountNumber())
	fmt.Println("seed: ", acc.Seed())
	fmt.Println("timestamp: ", timestamp)
	fmt.Println("signature: ", hex.EncodeToString(signature))
	fmt.Println("enc key: ", encKey)
}
