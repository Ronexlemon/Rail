package helpers

import (
	//"fmt"

	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	token "github.com/ronexlemon/rail/micro-services/wallet-service/utils/helpers/tokens/abi"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/ethclient"
)



type ChainBalanceResult struct {
	ChainName string
	USDC      string // Balance as a string (with decimals applied, e.g., "123.45")
	USDT      string
	Error     error // To report any errors for this specific chain
}


func GetAllChainBalances(userAddress string) []ChainBalanceResult {
	results := make(chan ChainBalanceResult, len(chainConfigs))
	var wg sync.WaitGroup

	for chainName, config := range chainConfigs {
		wg.Add(1) 
		go func(name string, cfg ChainConfig) {
			defer wg.Done()
			
			// Determine the chain type and call the appropriate handler
			var result ChainBalanceResult
			
			// Simple type check based on key/naming convention
			if strings.Contains(name, "solana") {
				result = checkSolanaBalance(name, cfg, userAddress)
			} else if strings.Contains(name, "tron") {
				//result = checkTronBalance(name, cfg, userAddress)
				
			} else {
				// Default to EVM for all others (Ethereum, Celo, Base, etc.)
				result = checkEVMBalance(name, cfg, userAddress)
			}

			results <- result
		}(chainName, config)
	}

	wg.Wait()
	close(results)

	var finalResults []ChainBalanceResult
	for res := range results {
		finalResults = append(finalResults, res)
	}

	return finalResults
}

func GetChainBalances(userAddress string, chain string) ChainBalanceResult {
	config, exists := chainConfigs[chain]
	if !exists {
		return ChainBalanceResult{
			ChainName: chain,
			Error:     fmt.Errorf("unsupported chain: %s", chain),
		}
	}

	var result ChainBalanceResult
	result.ChainName = chain

	// Choose the right handler based on the chain type
	if strings.Contains(strings.ToLower(chain), "solana") {
		result = checkSolanaBalance(chain, config, userAddress)
	} else if strings.Contains(strings.ToLower(chain), "tron") {
		// result = checkTronBalance(chain, config, userAddress)
		result.Error = fmt.Errorf("tron balance check not implemented")
	} else {
		// Default for EVM-compatible chains (Ethereum, Celo, Base, etc.)
		result = checkEVMBalance(chain, config, userAddress)
	}

	return result
}


func checkEVMBalance(chainName string, config ChainConfig, address string) ChainBalanceResult {
	// --- REAL WORLD: Initialize ethclient (Placeholder) ---
	client, err := ethclient.Dial(config.RPCURL)
	if err != nil {
		return ChainBalanceResult{ChainName: chainName, Error: fmt.Errorf("EVM RPC error: %w", err)}
	}
	defer client.Close()
	
	
	userCommonAddress := common.HexToAddress(address) 

	
	
	balanceUSDC, err := getERC20Balance(client, userCommonAddress, common.HexToAddress(config.TokenAddresses.USDC)) // USDC typically has 6 decimals
	if err != nil { 
		 return ChainBalanceResult{
			ChainName: chainName,
			Error:     fmt.Errorf("EVM RPC error: %w", err),
		}
	 }

	 balanceUSDT, err := getERC20Balance(client, userCommonAddress, common.HexToAddress(config.TokenAddresses.USDT)) // USDT often has 18 decimals
	if err != nil { 
		return ChainBalanceResult{
			ChainName: chainName,
			Error:     fmt.Errorf("EVM RPC error: %w", err),
		}
	 }

	// Simulating success
	return ChainBalanceResult{
		ChainName: chainName,
		USDC:      balanceUSDC.String(), 
		USDT:      balanceUSDT.String(),
		Error:     nil,
	}
}


func checkSolanaBalance(chainName string, config ChainConfig, userPubkey string) ChainBalanceResult {
	// --- REAL WORLD: Initialize Solana RPC Client (Placeholder) ---
	//client := solanarpc.New(config.RPCURL) 
	
	// Solana requires finding the Associated Token Account (ATA) for each token
	// ataUSDC, err := solana.GetATA(userPubkey, config.TokenAddresses.USDC)

	// --- REAL WORLD: Get Balances from ATA Accounts ---
	// balanceUSDC, err := client.GetTokenAccountBalance(ataUSDC)
	// balanceUSDT, err := client.GetTokenAccountBalance(ataUSDT)
	
	// Simulating success
	return ChainBalanceResult{
		ChainName: chainName,
		USDC:      "42.00", 
		USDT:      "100.50",
		Error:     nil,
	}
}

func checkTronBalance(chainName string, config ChainConfig, address string) ChainBalanceResult {
	// --- REAL WORLD: Initialize Tron RPC Client (Placeholder) ---
	// client := tronclient.New(config.RPCURL)

	// TRON addresses need to be converted to hex format (base58 to hex)
	// hexAddress := tron.Base58ToHex(address)

	// --- REAL WORLD: Call TRC-20 Contract Methods ---
	// TRON client often provides specific calls for TRC-20 tokens:
	// balanceUSDT, err := client.TRC20BalanceOf(hexAddress, config.TokenAddresses.USDT)

	// Simulating success
	return ChainBalanceResult{
		ChainName: chainName,
		USDC:      "0.00", 
		USDT:      "999.99",
		Error:     nil,
	}
}

func getERC20Balance(client *ethclient.Client, tokenAddress, userAddress common.Address) (*big.Float, error) {
	
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create token instance: %v", err)
	}

	
	bal, err := instance.BalanceOf(&bind.CallOpts{}, userAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch balance: %v", err)
	}

	
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch token decimals: %v", err)
	}

	// Convert balance to human-readable value (divide by 10^decimals)
	balanceFloat := new(big.Float).SetInt(bal)
	divisor := new(big.Float).SetFloat64(1)
	for i := uint8(0); i < decimals; i++ {
		divisor.Mul(divisor, big.NewFloat(10))
	}
	balanceInEther := new(big.Float).Quo(balanceFloat, divisor)

	return balanceInEther, nil
}