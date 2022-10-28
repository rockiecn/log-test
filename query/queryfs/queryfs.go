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
	"github.com/rockiecn/log-test/query/contracts/fs"
)

var (

	// fs address
	// for product
	fsAddress     = common.HexToAddress("0x9db4E0eF558865369a620C9B7953C965e0703D5f")
	addOrderTopic = "0x9ad187def591246eb8fff84534926005f62988636593745516c0ac08fa086314"
)

func main() {

	//client, err := ethclient.Dial("https://devchain.metamemo.one:8501")
	client, err := ethclient.Dial("https://chain.metamemo.one:8501")
	//client, err := ethclient.Dial("https://testchain.metamemo.one:24180")

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
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(8228622),
		Addresses: []common.Address{
			fsAddress,
		},
	}

	// filter logs
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	println("logs length:", len(logs))

	// get contract abi
	fsAbi, err := abi.JSON(strings.NewReader(string(fs.FileSysABI)))
	if err != nil {
		log.Fatal(err)
	}
	_ = fsAbi

	totalIssue := new(big.Int)

	// tranvel each log
	fmt.Println("Travelling logs")
	for i, vLog := range logs {
		fmt.Printf("===== show log %d =====\n", i)
		// fmt.Println("block hash:", vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		// fmt.Println("tx hash:", vLog.TxHash.Hex())       // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		// event := struct {
		// 	Key   [32]byte
		// 	Value [32]byte
		// }{}

		// fmt.Println(string(event.Addr[:]))  // foo
		// fmt.Println(string(event.Index[:])) // bar

		//event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice);
		// 4 topics in all, the first is the hash of the event
		topics := make([]string, 4)
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
			fmt.Printf("topic%d: %s\n", i, topics[i])
		}

		fmt.Println("check log for AddOrder event")
		// unpack log data
		// logData, err := fsAbi.Unpack("AddOrder", vLog.Data)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// only 1 data in Transfer event
		//fmt.Println("data:", res[0])

		fmt.Println("searching for addorder action")
		// find the addorder event, aggregate the total issue
		cmp := strings.Compare(topics[0], addOrderTopic)
		if cmp == 0 {
			fmt.Println("---- Found an  addorder event! ----")
			fmt.Println("block number:", vLog.BlockNumber) // 2490757
			fmt.Println("txHash:", vLog.TxHash)

			// start
			// start, ok := logData[0].(*big.Int)
			// if !ok {
			// 	log.Fatal("log data assertion failed")
			// }
			// fmt.Println("start:", start.String())

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
