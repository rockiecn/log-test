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
	"github.com/rockiecn/log-test/query/contracts/erc"
)

var (

	// erc20 token address
	// for test
	//erc20Address = common.HexToAddress("0xD5346Db60E116bd8c98181C9235bC9BAb6d64031")

	// for economic env
	erc20Address = common.HexToAddress("0x9Ca0f1A85795ea800Ae3F8ed53e174297fbe4c6b")

	// for product
	//erc20Address = common.HexToAddress("0xeB8eec5a2dBf6e6f4Cc542ad31CCe706f8f80419")
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
		ToBlock:   big.NewInt(3859716),
		Addresses: []common.Address{
			erc20Address,
		},
	}

	// filter logs
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	println("logs length:", len(logs))

	// get contract abi
	erc20Abi, err := abi.JSON(strings.NewReader(string(erc.ERC20ABI)))
	if err != nil {
		log.Fatal(err)
	}

	totalIssue := new(big.Int)

	// tranvel each log
	fmt.Println("Travelling logs")
	for i, vLog := range logs {
		// fmt.Println("block hash:", vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
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

		//fmt.Println("check log for Transfer event")
		// unpack log data
		res, err := erc20Abi.Unpack("Transfer", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		// only 1 data in Transfer event
		//fmt.Println("data:", res[0])

		//fmt.Println("searching for issu action")
		// find the issue event, aggregate the total issue
		cmp := strings.Compare(topics[1], "0x0000000000000000000000000000000000000000000000000000000000000000")
		if cmp == 0 {
			//fmt.Println("---- log index:", i)
			fmt.Println("block number:", vLog.BlockNumber) // 2490757
			fmt.Println("txHash:", vLog.TxHash)
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
			totalIssue = totalIssue.Add(totalIssue, issu)
		} else {
			continue
		}

		fmt.Println()
	}

	fmt.Println("total issue:", totalIssue)

	// eventSignature := []byte("Transfer(address indexed, address indexed, uint256)")
	// hash := crypto.Keccak256Hash(eventSignature)
	// fmt.Println("event hash:", hash.Hex()) // 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef

}
