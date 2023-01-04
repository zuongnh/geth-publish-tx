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

	store "github.com/zuongnh/publish-tx/contracts" // for demo
)

func main() {
	client, err := ethclient.Dial("wss://bsc-testnet.nodereal.io/ws/v1/e9a36765eb8a40b9bd12e680a1fd2bc5")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("341ff1cc75cc7a421982cc904c075c4819b528d80182f36cfab3530c953f3fb6")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0xa66feE583714f35cbB20b8322F50cF5AD21F5388")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	zombieId := big.NewInt(104)
	targetId := big.NewInt(4)

	// Publish Attack transaction
	tx, err := instance.Attack(auth, zombieId, targetId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	// result, err := instance.Zombies(nil, zombieId)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(result) // "bar"
}
