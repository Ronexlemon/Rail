package helpers


import (
	"errors"
	"strings"
)


type TokenAddresses struct {
	USDC string
	USDT string
}


type ChainConfig struct {
	RPCURL         string
	TokenAddresses TokenAddresses
}


var chainConfigs = map[string]ChainConfig{
	// Ethereum Mainnet
	"ethereum_mainnet": {
		RPCURL: ETHEREUM_MAINNET,
		TokenAddresses: TokenAddresses{
			USDC: USDC_ETHEREUM_MAINNET,
			USDT: USDT_ETHEREUM_MAINNET,
		},
	},
	// Sepolia Testnet
	"sepolia": {
		RPCURL: SEPOLIA,
		TokenAddresses: TokenAddresses{
			USDC: USDC_ETHEREUM_SEPOLIA,
			USDT: USDT_ETHEREUM_SEPOLIA,
		},
	},
	// Celo Mainnet (using "celo" as the key)
	"celo": {
		RPCURL: CELO,
		TokenAddresses: TokenAddresses{
			USDC: USDC_CELO_MAINNET,
			USDT: USDT_CELO_MAINNET,
		},
	},
	// Base Mainnet
	"base": {
		RPCURL: BASE,
		TokenAddresses: TokenAddresses{
			USDC: USDC_BASE_MAINNET,
			USDT: USDT_BASE_MAINNET,
		},
	},
	// Add other chains here...
	"polygon_mainnet": {
		RPCURL: POLYGON_MAINNET,
		TokenAddresses: TokenAddresses{
			USDC: USDC_POLYGON_MAINNET,
			USDT: USDT_POLYGON_MAINNET,
		},
	},
	"arbitrum_mainnet": {
		RPCURL: ARBITRUM_MAINNET,
		TokenAddresses: TokenAddresses{
			USDC: USDC_ARBITRUM_MAINNET,
			USDT: USDT_ARBITRUM_MAINNET,
		},
	},
	"solana": {
		RPCURL: SOLANA_MAINNET,
		TokenAddresses: TokenAddresses{
			USDC: USDC_SOLANA_MAINNET,
			USDT: USDT_SOLANA_MAINNET,
		},
	},

	"tron": {
		RPCURL: TRON_MAINNET,
		TokenAddresses: TokenAddresses{
			USDC: USDC_TRON_MAINNET,
			USDT: USDT_TRON_MAINNET,
		},
	},
}


func GetChainConfig(chainName string) (ChainConfig, error) {
	
	key := strings.ToLower(strings.ReplaceAll(chainName, " ", "_"))
	
	config, ok := chainConfigs[key]
	if !ok {
		return ChainConfig{}, errors.New("chain configuration not found for: " + chainName)
	}

	return config, nil
}



