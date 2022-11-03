package monitor

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewMonitor(contractAddress common.Address, contractAbi abi.ABI, client *ethclient.Client) *Monitor {
	return &Monitor{
		ContractAddress: contractAddress,
		ContractAbi:     contractAbi,
		Client:          client,
		Filter: ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
		},
	}
}

func (monitor *Monitor) Start(ctx context.Context, events chan AnswerUpdatedEvent) error {
	logs := make(chan types.Log)
	sub, err := monitor.Client.SubscribeFilterLogs(ctx, monitor.Filter, logs)
	if err != nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-sub.Err():
			return err
		case log := <-logs:
			var answerUpdatedEvent AnswerUpdatedEvent
			err := monitor.ContractAbi.UnpackIntoInterface(&answerUpdatedEvent, "AnswerUpdated", log.Data)
			if err != nil {
				return err
			}
			events <- answerUpdatedEvent
		}
	}
}
