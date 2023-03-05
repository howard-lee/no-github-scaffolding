package grpcData

// // for grpc items di mock stores
import (
	grpcApi "api"
)

type BlkStore interface {
	GetBalance(name string) []grpcApi.Balance
}
type WalletStore interface {
	//GetAccounts(name string) []grpcApi.Account
	GetWallets(name string) []grpcApi.Wallet
}

type GetBalance struct {
	DataStore BlkStore
}

func NewGetBalance(dataStore BlkStore) *GetBalance {
	return &GetBalance{
		DataStore: dataStore,
	}
}

type GetWallets struct {
	DataStore WalletStore
}

func NewGetWallets(dataStore WalletStore) *GetWallets {
	return &GetWallets{
		DataStore: dataStore,
	}
}
