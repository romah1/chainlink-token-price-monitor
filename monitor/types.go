package monitor

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Monitor struct {
	ContractAddress common.Address
	ContractAbi     abi.ABI
	Client          *ethclient.Client
	Filter          ethereum.FilterQuery
}

type AnswerUpdatedEvent struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
}
