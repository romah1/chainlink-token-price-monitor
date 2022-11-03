package main

import (
	"context"
	"flag"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/romah1/chainlink-token-price-monitor/monitor"
	"strings"
)

func main() {
	var wssEndpoint string
	var contractAddress string
	var contractAbiString string

	flag.StringVar(&wssEndpoint, "wss-endpoint", "", "wss endpoint url")
	flag.StringVar(&contractAbiString, "contract-address", "", "contract address")
	flag.StringVar(&contractAbiString, "contract-abi", "", "contract abi string")
	flag.Parse()
	
	client, err := ethclient.Dial(wssEndpoint)
	if err != nil {
		panic(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(contractAbiString))

	priceMonitor := monitor.NewMonitor(common.HexToAddress(contractAddress), contractAbi, client)
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	events := make(chan monitor.AnswerUpdatedEvent)
	err = priceMonitor.Start(ctx, events)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case event := <-events:
			println(event.Current)
		}
	}
}
