package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func event_subscribe() {
	client, err := ethclient.Dial("wss://bsc-testnet.nodereal.io/ws/v1/e9a36765eb8a40b9bd12e680a1fd2bc5")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x3986C657C4597064825b4eaeC2FFac4fc501644f")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
