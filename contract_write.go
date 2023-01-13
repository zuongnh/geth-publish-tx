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

func contract_write() {
	// -------------- SET UP --------------
	client, err := ethclient.Dial("wss://bsc-testnet.nodereal.io/ws/v1/e9a36765eb8a40b9bd12e680a1fd2bc5")
	if err != nil {
		log.Fatal(err)
	}

	// privateKey, err := crypto.HexToECDSA("341ff1cc75cc7a421982cc904c075c4819b528d80182f36cfab3530c953f3fb6")
	privateKey, err := crypto.HexToECDSA("57d7a2a1188476d970f71e7dc4084281a66499345e5a1421e0dc15b1e74c631a")
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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice

	// -------------- CREATE SC INSTANCE --------------
	address := common.HexToAddress("0xa66feE583714f35cbB20b8322F50cF5AD21F5388")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// -------------- PUBLISH TX --------------
	publish_open_pack_tx(instance, auth)

	zombieId := big.NewInt(104)
	targetId := big.NewInt(4)
	publish_attack_tx(instance, auth, zombieId, targetId)

	get_zombies_of(instance, fromAddress)
}

// Publish openStarterPack transaction
func publish_open_pack_tx(instance *store.Store, auth *bind.TransactOpts) {
	tx, err := instance.OpenStarterPack(auth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

// Publish Attack transaction
func publish_attack_tx(instance *store.Store, auth *bind.TransactOpts, zombieId *big.Int, targetId *big.Int) {
	tx, err := instance.Attack(auth, zombieId, targetId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

// Publish get all Zombies of an address transaction
func get_zombies_of(instance *store.Store, owner common.Address) {
	res, err := instance.GetZombieOf(nil, owner)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %v", res)
}
