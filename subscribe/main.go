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

func main() {
	fmt.Println("dialing..")
	// dev
	client, err := ethclient.Dial("wss://devchain.metamemo.one:6901")

	if err != nil {
		log.Fatal(err)
	}

	// dev
	contractAddress := common.HexToAddress("0xb069d34acDE5F177ab2b21759f54972a63234F3f")

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	fmt.Println("subscribing..")
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("print..")
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
