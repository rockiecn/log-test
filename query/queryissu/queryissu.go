package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	// for demo
	"github.com/rockiecn/log-test/query/contracts/issuance"
)

var (
	issuAddress = common.HexToAddress("0xEbb7Cb7b7F60a0B4E58266C30D25BfE00dC8a25C")
	issuTopic   = "0x11b3bd0b1439fef956ad5b4c0b5c1fb4eba94e9c2a8ace3497a5b0712fb4c828"
)

func main() {

	//client, err := ethclient.Dial("https://devchain.metamemo.one:8501")
	//client, err := ethclient.Dial("https://chain.metamemo.one:8501")
	client, err := ethclient.Dial("https://testchain.metamemo.one:24180")

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	bn, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chain's latest block number:", bn)

	// create a query
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(3757169),
		ToBlock:   big.NewInt(3945936),
		Addresses: []common.Address{
			issuAddress,
		},
	}

	// filter logs
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	println("issu logs length:", len(logs))

	// get contract abi
	issuAbi, err := abi.JSON(strings.NewReader(string(issuance.IssuanceABI)))
	if err != nil {
		log.Fatal(err)
	}
	_ = issuAbi

	// tranvel each log
	fmt.Println("Travelling issu logs")
	for i, vLog := range logs {
		fmt.Printf("===== show log %d =====\n", i)

		//event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice);
		// 4 topics in all, the first is the hash of the event
		fmt.Println("parsing topics..")
		topics := make([]string, 4)
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
			//fmt.Printf("topic%d: %s\n", i, topics[i])
		}

		fmt.Println("searching for issu action")
		cmp := strings.Compare(topics[0], issuTopic)
		if cmp == 0 {
			fmt.Println("---- Found an issu event! ----")
			fmt.Println("block number:", vLog.BlockNumber) // 2490757
			fmt.Println("txHash:", vLog.TxHash)

			fmt.Println("topics details:")
			fmt.Println("event Issu(uint256 indexed sprice, uint64 indexed now, uint64 indexed last);")
			for i := range vLog.Topics {
				fmt.Printf("topic%d: %s\n", i, topics[i])
			}
		} else {
			continue
		}

		fmt.Println()
	}
}
