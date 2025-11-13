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
		return nil, fmt.Errorf("prisma client is nil")
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
		return fmt.Errorf("prisma client is nil")
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


func (r *WalletRepository) GetBusinessWallets(businessId string) ([]db.WalletModel, error) {
	if r.Client == nil {
		return nil, fmt.Errorf("prisma client is nil")
	}

	wallets, err := r.Client.Wallet.FindMany(
		db.Wallet.BusinessID.Equals(businessId),
	).With(
		db.Wallet.Addresses.Fetch(),
		db.Wallet.VirtualAccounts.Fetch(),
	).Exec(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to fetch wallets: %w", err)
	}

	return wallets, nil
}

func (r *WalletRepository) GetBusinessCustomerWallets(businessId,customerId string) ([]db.WalletModel, error) {
	if r.Client == nil {
		return nil, fmt.Errorf("prisma client is nil")
	}

	wallets, err := r.Client.Wallet.FindMany(
	db.Wallet.And(
		db.Wallet.BusinessID.Equals(businessId),
		db.Wallet.CustomerID.Equals(customerId),
	),
).With(
	db.Wallet.Addresses.Fetch(),
	db.Wallet.VirtualAccounts.Fetch(),
).Exec(r.Context)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch wallets: %w", err)
	}

	return wallets, nil
}


func (r *WalletRepository) GetWalletAddressByNetwork(businessId string, customerId *string, network db.Network) (*db.WalletAddressModel, error) {
	if r.Client == nil {
		return nil, fmt.Errorf("prisma client is nil")
	}

	if businessId == "" && (customerId == nil || *customerId == "") {
		return nil, fmt.Errorf("either businessId or customerId must be provided")
	}

	var wallet *db.WalletModel
	var err error

	// Fetch the wallet depending on whether customerId is provided
	if customerId != nil && *customerId != "" {
		wallet, err = r.Client.Wallet.FindFirst(
			db.Wallet.BusinessID.Equals(businessId),
			db.Wallet.CustomerID.Equals(*customerId),
		).With(
			db.Wallet.Addresses.Fetch(),
		).Exec(context.Background())
	} else {
		wallet, err = r.Client.Wallet.FindFirst(
			db.Wallet.BusinessID.Equals(businessId),
		).With(
			db.Wallet.Addresses.Fetch(),
		).Exec(context.Background())
	}

	if err != nil {
		return nil, fmt.Errorf("failed to fetch wallet: %w", err)
	}

	if wallet == nil {
		return nil, fmt.Errorf("wallet not found")
	}

	// Filter the wallet address for the requested network
	for _, addr := range wallet.Addresses() {
		if addr.Network == network {
			return &addr, nil
		}
	}

	return nil, fmt.Errorf("wallet address not found for network %s", network)
}

