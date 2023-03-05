package grpcData

import (
	grpcApi "api"
	"fmt"
)

var balanceData = []grpcApi.Balance{

	grpcApi.Balance{
		Account: "1001",
		Balance: 1001,
	},
	grpcApi.Balance{
		Account: "1003",
		Balance: 1003,
	},
	grpcApi.Balance{
		Account: "1006",
		Balance: 1006,
	},
}

var walletData = []grpcApi.Wallet{
	grpcApi.Wallet{
		WalletID: "135246",
		URL:      "grpc://localhost:30031/135246",
	},
	grpcApi.Wallet{
		WalletID: "200301",
		URL:      "grpc://localhost:200301",
	},
	grpcApi.Wallet{
		WalletID: "143467",
		URL:      "grpc://localhost:143467",
	},
}

var accountData = []grpcApi.Account{
	grpcApi.Account{
		AccountID: "135246",
		Nounce:    0,
		Balance:   100,
	},
	grpcApi.Account{
		AccountID: "200301",
		Nounce:    0,
		Balance:   200,
	},
	grpcApi.Account{
		AccountID: "143467",
		Nounce:    0,
		Balance:   300,
	},
}

// memorystore
type CachedBlkStore struct {
	//blkStore data.Blkstore
	//GetBalance(string) []grpcApi.Balance
}

func (c *CachedBlkStore) GetBalance(account string) []grpcApi.Balance {
	// will always return an slice, if there be no data, will return single element of slice of empty data
	var bal []grpcApi.Balance
	found := false
	if account == "" {
		bal = balanceData // if input is empty, return all balances
	} else {
		for _, v := range balanceData {
			if v.Account == account {
				bal = append(bal, v)
				found = true
			}
		}
		if found == false {
			bal = make([]grpcApi.Balance, 1)
			return bal
		}
	}

	return bal
}

type CachedWalletStore struct {
}

func (w *CachedWalletStore) GetWallets(id string) []grpcApi.Wallet {
	// will always return an slice, if there be no data, will return single element of slice of empty data

	var wal []grpcApi.Wallet
	found := false
	if id == "" {
		// return all data in cached aray // special case
		//wal = walletData // if input is empty, return all balances
		for _, v := range walletData {
			res := &grpcApi.Response{
				Code:    0,
				Message: "Wallet data found.",
			}
			v.Response = res
			wal = append(wal, v)
		}
		//fmt.Println(" end of GetWallet, length of wal is ", len(wal))
		return wal

	} else {
		for _, v := range walletData {
			res := &grpcApi.Response{
				Code:    0,
				Message: "Wallet data found.",
			}
			v.Response = res
			if v.WalletID == id {
				wal = append(wal, v)
				found = true
			}
		}
		if found == false {
			wal = make([]grpcApi.Wallet, 1)
			res := &grpcApi.Response{
				Code:    404,
				Message: fmt.Sprintf("Cannot found wallet with %v.\n", id),
			}
			wal[0].Response = res
			return wal
		}
	}
	//fmt.Println(" end of GetWallet, length of wal is ", len(wal))
	return wal
}

func (w *CachedWalletStore) GetAccount(id string) []grpcApi.Account {
	// will always return an slice, if there be no data, will return single element of slice of empty data
	var acc []grpcApi.Account
	found := false
	if id == "" {
		acc = accountData // if input is empty, return all balances
	} else {
		for _, v := range accountData {
			if v.AccountID == id {
				acc = append(acc, v)
				found = true
			}
		}
		if found == false {
			acc = make([]grpcApi.Account, 1)
			return acc
		}
	}

	return acc
}
