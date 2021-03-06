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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	role "github.com/rockiecn/log-test/query/contracts" // for demo
)

func main() {
	client, err := ethclient.Dial("wss://devchain.metamemo.one:6901")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x393ee8b2726E144eB0AB056636f38358719dF17E")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(8622),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	println("logs length:", len(logs))

	contractAbi, err := abi.JSON(strings.NewReader(string(role.IRoleSetterABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		// event := struct {
		// 	Key   [32]byte
		// 	Value [32]byte
		// }{}

		// event := struct {
		// 	Addr  [32]byte
		// 	Index [64]byte
		// }{}

		res, err := contractAbi.Unpack("ReAcc", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(string(event.Addr[:]))  // foo
		// fmt.Println(string(event.Index[:])) // bar

		fmt.Println(res[0])
		fmt.Println(res[1])

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	eventSignature := []byte("ReAcc(address,uint64)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}
