package helpers

import (
	//"fmt"
	"strings"
	"sync"

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
				result = checkTronBalance(name, cfg, userAddress)
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


func checkEVMBalance(chainName string, config ChainConfig, address string) ChainBalanceResult {
	// --- REAL WORLD: Initialize ethclient (Placeholder) ---
	// client, err := ethclient.Dial(config.RPCURL)
	// if err != nil {
	// 	return ChainBalanceResult{ChainName: chainName, Error: fmt.Errorf("EVM RPC error: %w", err)}
	// }
	// defer client.Close()
	
	// Convert userAddress to common.Address (Placeholder)
	//userCommonAddress := common.HexToAddress(address) 

	// --- REAL WORLD: Get Balances via Token Contract Calls ---
	
	//balanceUSDC, err := getERC20Balance(client, userCommonAddress, config.TokenAddresses.USDC, 6) // USDC typically has 6 decimals
	// if err != nil { /* Handle error */ }

	// balanceUSDT, err := getERC20Balance(client, userCommonAddress, config.TokenAddresses.USDT, 18) // USDT often has 18 decimals
	// if err != nil { /* Handle error */ }

	// Simulating success
	return ChainBalanceResult{
		ChainName: chainName,
		USDC:      "123.45", 
		USDT:      "500.00",
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