// package repository



// import (
// 	"context"
// 	"fmt"

// 	"github.com/ronexlemon/rail/micro-services/wallet-service/database"
// 	"github.com/ronexlemon/rail/micro-services/wallet-service/prisma/db"
// )

// type WalletRepository struct {
// 	Client  *db.PrismaClient
// 	Context context.Context
// }

// func NewWalletRepository() *WalletRepository {
// 	if database.PrismaDBClient.Client == nil {
// 		panic(" Prisma client not initialized") 
// 	}
// 	return &WalletRepository{
// 		Client:  database.PrismaDBClient.Client,
// 		Context: database.PrismaDBClient.Context,
// 	}
// }

// type CreateWalletInput struct {
// 	BusinessID string
// 	CustomerID *string
// 	Type       db.WalletType
// 	Network    db.Network
// 	Address    string
// 	PrivateKey string
// }

// // ✅ Create wallet for business or customer
// func (r *WalletRepository) CreateWallet(input CreateWalletInput) (*db.WalletModel, error) {
// 	if r.Client == nil {
// 		return nil, fmt.Errorf("Prisma client is nil")
// 	}

// 	wallet, err := r.Client.Wallet.CreateOne(
// 		db.Wallet.BusinessID.Set(input.BusinessID),
// 		db.Wallet.Type.Set(input.Type),
// 		db.Wallet.CustomerID.SetOptional(input.CustomerID),
		
// 	).Exec(r.Context)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create wallet: %w", err)
// 	}
 
// 	_, err = r.Client.WalletAddress.CreateOne(
       
// 		db.WalletAddress.Network.Set(input.Network), 
//         db.WalletAddress.Address.Set(input.Address), 
//         db.WalletAddress.PrivateKey.Set(input.PrivateKey), 
//         db.WalletAddress.Wallet.Link(db.Wallet.ID.Equals(wallet.ID)),
        
//     ).Exec(r.Context)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create wallet address: %w", err)
// 	}

// 	return wallet, nil
// }

package repository

import (
	"context"
	"fmt"

	"github.com/ronexlemon/rail/micro-services/wallet-service/database"
	"github.com/ronexlemon/rail/micro-services/wallet-service/prisma/db"
)

type WalletRepository struct {
	Client  *db.PrismaClient
	Context context.Context
}

func NewWalletRepository() *WalletRepository {
	if database.PrismaDBClient.Client == nil {
		panic("Prisma client not initialized")
	}
	return &WalletRepository{
		Client:  database.PrismaDBClient.Client,
		Context: database.PrismaDBClient.Context,
	}
}

type CreateWalletInput struct {
	BusinessID string
	CustomerID *string
	Type       db.WalletType
}

// ✅ Create the main wallet for a business or customer
func (r *WalletRepository) CreateWallet(input CreateWalletInput) (*db.WalletModel, error) {
	if r.Client == nil {
		return nil, fmt.Errorf("Prisma client is nil")
	}

	wallet, err := r.Client.Wallet.CreateOne(
		db.Wallet.BusinessID.Set(input.BusinessID),
		db.Wallet.Type.Set(input.Type),
		db.Wallet.CustomerID.SetOptional(input.CustomerID),
	).Exec(r.Context)

	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}

	return wallet, nil
}

// ✅ Create wallet address for a specific network
func (r *WalletRepository) CreateWalletAddress(walletID string, network db.Network, address string, privateKey string) error {
	if r.Client == nil {
		return fmt.Errorf("Prisma client is nil")
	}

	_, err := r.Client.WalletAddress.CreateOne(
		db.WalletAddress.Network.Set(network), 
        db.WalletAddress.Address.Set(address), 
        db.WalletAddress.PrivateKey.Set(privateKey), 
       db.WalletAddress.Wallet.Link(db.Wallet.ID.Equals(walletID)),
	).Exec(r.Context)

	if err != nil {
		return fmt.Errorf("failed to create wallet address for %s: %w", network, err)
	}
	return nil
}
