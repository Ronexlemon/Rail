package service

import (
	"context"
	"fmt"
	"log"

	"github.com/ronexlemon/rail/micro-services/wallet-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/wallet-service/prisma/db"
	"github.com/ronexlemon/rail/micro-services/wallet-service/utils/chains"
	"github.com/ronexlemon/rail/micro-services/wallet-service/utils/helpers"
)

type WalletService struct {
	repo *repository.WalletRepository
}

func NewWalletService(repo *repository.WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

// CreateWallet will create a Wallet and automatically create WalletAddresses
// for all 3 supported networks (EVM, SOLANA, TRON)
func (s *WalletService) CreateWallet(ctx context.Context, businessID string, customerID *string, walletType db.WalletType) (*db.WalletModel, error) {
	if businessID == "" && (customerID == nil || *customerID == "") {
		return nil, fmt.Errorf("either businessID or customerID must be provided")
	}

	
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	log.Printf("Creating wallet for businessID=%s customerID=%v", businessID, customerID)

	// Prepare input
	walletInput := repository.CreateWalletInput{
		BusinessID: businessID,
		CustomerID: customerID,
		Type:       walletType,
	}

	// Create wallet in DB
	wallet, err := s.repo.CreateWallet(walletInput)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}

	// Supported networks
	networks := []db.Network{db.NetworkEvm, db.NetworkSolana, db.NetworkTron}

	for _, network := range networks {
		// Check context before each network operation
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		var address, privateKey string

		switch network {
		case db.NetworkEvm:
			addr, priv := chains.CreateEVMWallet()
			address = addr.Hex()
			privateKey = priv

		case db.NetworkSolana:
			result, err := chains.CreateSolanaWallet()
			if err != nil {
				log.Printf("Solana wallet creation failed: %v", err)
				continue
			}
			address = result.Address
			privateKey = result.PrivateKey

		case db.NetworkTron:
			result, err := chains.CreateTronWallet()
			if err != nil {
				log.Printf("Tron wallet creation failed: %v", err)
				continue
			}
			address = result.Address
			privateKey = result.PrivateKey
		}

		// Save wallet address
		err = s.repo.CreateWalletAddress(wallet.ID, network, address, privateKey)
		if err != nil {
			log.Printf("Failed to save %s wallet address: %v", network, err)
		} else {
			log.Printf("%s wallet created: %s", network, address)
		}
	}

	log.Println("Wallet and all network addresses created successfully")
	return wallet, nil
}


func (s *WalletService) BusinessWallet(ctx context.Context, businessID string, customerID *string) ([]db.WalletModel, error) {
	if businessID == "" && (customerID == nil || *customerID == "") {
		return nil, fmt.Errorf("either businessID or customerID must be provided")
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	log.Printf("Fetching wallets for businessID=%s, customerID=%v", businessID, customerID)

	var (
		wallets []db.WalletModel
		err     error
	)

	// Fetch depending on whether customerID is provided
	if customerID != nil && *customerID != "" {
		wallets, err = s.repo.GetBusinessCustomerWallets(businessID, *customerID)
	} else {
		wallets, err = s.repo.GetBusinessWallets(businessID)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to fetch wallets: %w", err)
	}

	log.Printf("Fetched %d wallet(s) for businessID=%s", len(wallets), businessID)
	return wallets, nil
}


func (s *WalletService) WalletAddresses(ctx context.Context, businessID string, customerID *string, network db.Network) (map[string][]helpers.ChainBalanceResult, error) {
	if businessID == "" && (customerID == nil || *customerID == "") {
		return nil, fmt.Errorf("either businessID or customerID must be provided")
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// Fetch wallet address for the specified network
	addr, err := s.repo.GetWalletAddressByNetwork(businessID, customerID, network)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch wallet address for network %s: %w", network, err)
	}

	// Fetch balances for the address
	balances := helpers.GetAllChainBalances(addr.Address)

	// Return as a map[address] -> balances
	balancesMap := map[string][]helpers.ChainBalanceResult{
		addr.Address: balances,
	}

	return balancesMap, nil
}

