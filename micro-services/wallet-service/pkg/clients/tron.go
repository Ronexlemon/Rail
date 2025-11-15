package clients

import (
	"fmt"
	
	"math/big"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
)

type NewGRPCClient struct {
	Client *client.GrpcClient
}

func NewTronClient(trongrpcURL string) *NewGRPCClient {
	return &NewGRPCClient{
		Client: client.NewGrpcClient(trongrpcURL),
	}
}

func (t *NewGRPCClient) ContractAbi(contractAddress string) (*core.SmartContract_ABI, error) {
	contractAbi, err := t.Client.GetContractABI(contractAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get contract ABI: %w", err)
	}
	return contractAbi, nil
}

func (t *NewGRPCClient) TRC20Name(contractAddress string) (string, error) {
	name, err := t.Client.TRC20GetName(contractAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get TRC20 name: %w", err)
	}
	return name, nil
}

func (t *NewGRPCClient) TRC20Decimal(contractAddress string) (*big.Int, error) {
	val, err := t.Client.TRC20GetDecimals(contractAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get TRC20 decimals: %w", err)
	}
	return val, nil
}

func (t *NewGRPCClient) TRC20ContractBalance(addr string, contractAddress string) (*big.Int, error) {
	val, err := t.Client.TRC20ContractBalance(addr, contractAddress)
	
	if err != nil {
		return nil, fmt.Errorf("failed to get TRC20 contract balance: %w", err)
	}
	return val, nil
}

