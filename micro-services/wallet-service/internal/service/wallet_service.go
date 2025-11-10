package service

import (
	"fmt"
	"log"

	"github.com/ronexlemon/rail/micro-services/wallet-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/wallet-service/prisma/db"
	"github.com/ronexlemon/rail/micro-services/wallet-service/utils/chains"
)

type WalletService struct {
	repo *repository.WalletRepository
}

func NewWalletService(repo *repository.WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

// CreateWallet will create a Wallet and automatically create WalletAddresses
// for all 3 supported networks (EVM, SOLANA, TRON)
func (s *WalletService) CreateWallet(businessID string, customerID *string, walletType db.WalletType) (*db.WalletModel, error) {
	if businessID == "" && customerID == nil {
		return nil, fmt.Errorf("either businessID or customerID must be provided")
	}

	log.Printf("Creating wallet for businessID=%s customerID=%v", businessID, customerID)

	
	walletInput := repository.CreateWalletInput{
		BusinessID: businessID,
		CustomerID: customerID,
		Type:       walletType,
	}

	wallet, err := s.repo.CreateWallet(walletInput)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}

	
	networks := []db.Network{db.NetworkEvm, db.NetworkSolana, db.NetworkTron}

	for _, network := range networks {
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
			result,err := chains.CreateTronWallet()
			if err != nil {
				log.Printf("Tron wallet creation failed: %v", err)
				continue
			}
			address = result.Address
			privateKey = result.PrivateKey
		}

		
		err = s.repo.CreateWalletAddress(wallet.ID, network, address, privateKey)
		if err != nil {
			log.Printf("Failed to save %s wallet address: %v", network, err)
		} else {
			log.Printf(" %s wallet created: %s", network, address)
		}
	}

	log.Println(" Wallet and all network addresses created successfully")
	return wallet, nil
}
