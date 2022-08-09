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

	erc "github.com/rockiecn/log-test/query/contracts" // for demo
)

func main() {
	client, err := ethclient.Dial("https://testchain.metamemo.one:24180")
	if err != nil {
		log.Fatal(err)
	}

	// erc20 token address
	contractAddress := common.HexToAddress("0x37aC6152B689EBEBdF311710ca541B2413777b7d")
	// create a query
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(8228622),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	// filter logs
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	println("logs length:", len(logs))

	// get contract abi
	contractAbi, err := abi.JSON(strings.NewReader(string(erc.ERC20ABI)))
	if err != nil {
		log.Fatal(err)
	}

	totalInssue := new(big.Int)

	// tranvel each log
	fmt.Println("Travelling logs")
	for i, vLog := range logs {
		// fmt.Println("log index:", i)
		// fmt.Println("block hash:", vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		// fmt.Println("block number:", vLog.BlockNumber)   // 2394201
		// fmt.Println("tx hash:", vLog.TxHash.Hex())       // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		// event := struct {
		// 	Key   [32]byte
		// 	Value [32]byte
		// }{}

		// fmt.Println(string(event.Addr[:]))  // foo
		// fmt.Println(string(event.Index[:])) // bar

		// 3 topics in all, the first is the hash of the event
		topics := make([]string, 3)
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		// unpack log data
		res, err := contractAbi.Unpack("Transfer", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		// only 1 data in Transfer event
		//fmt.Println("data:", res[0])

		// find the issue event, aggregate the total issue
		cmp := strings.Compare(topics[1], "0x0000000000000000000000000000000000000000000000000000000000000000")
		if cmp == 0 {
			fmt.Println("---- log index:", i)
			fmt.Println("block number:", vLog.BlockNumber) // 2490757
			for i := range vLog.Topics {
				fmt.Printf("topic %d: %s\n", i, topics[i])
			}

			// one time issue
			issu, ok := res[0].(*big.Int)
			if !ok {
				log.Fatal("res assertion failed")
			}
			fmt.Println("issue this time:", issu.String())

			// accumulate
			totalInssue = totalInssue.Add(totalInssue, issu)
		} else {
			continue
		}

		fmt.Println()
	}

	fmt.Println("total issue:", totalInssue)

	// eventSignature := []byte("Transfer(address indexed, address indexed, uint256)")
	// hash := crypto.Keccak256Hash(eventSignature)
	// fmt.Println("event hash:", hash.Hex()) // 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
}
