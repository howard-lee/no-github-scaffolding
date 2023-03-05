package server

import (
	grpcApi "api"
	"fmt"
)

type WalletStore interface {
	GetWallets(name string) []grpcApi.Wallet
}

type BlkWallet struct {
	grpcApi.UnimplementedBlkWalletServer
	//grpcApi.Wallet
	WalletStore
	GenericGetWallet interface{}
	//store []grpcApi.Wallet
}

func NewBlkWallet() (*BlkWallet, error) {
	return &BlkWallet{}, nil
}

func (c *BlkWallet) GetWallets(wid string) (*grpcApi.Wallets, error) {
	fmt.Println("@@@@ ")
	if c.GenericGetWallet == nil {
		fmt.Println(" c.GenericGetWallet  empty")
	} else {
		fmt.Println(" c.GenericGetWallet not empty")

		switch v := c.GenericGetWallet.(type) {
		case WalletStore:
			fmt.Println(" it is walletstore:", v)
		// case float64:
		// 	fmt.Println("float64:", v)
		default:
			fmt.Printf("unknown %v \n", v)
		}
	}
	//walcol := c.WalletStore.GetWallets(wid)

	//result := true
	var wls []*grpcApi.Wallet
	var wellets grpcApi.Wallets
	// if wid != "" {
	// 	for _, v := range walcol {
	// 		//var bl grpcApi.Wallet
	// 		if v.WalletID == wid {
	// 			bl := grpcApi.Wallet{
	// 				WalletID: v.WalletID,
	// 				URL:      v.URL,
	// 				Gaslimit: v.Gaslimit,
	// 				Response: &grpcApi.Response{
	// 					Code:    0,
	// 					Message: "get wallet suessfully complete",
	// 				},
	// 			}
	// 			wls = append(wls, &bl)
	// 			wellets = grpcApi.Wallets{
	// 				Wallets: wls,
	// 			}
	// 		}

	// 		return &wellets, nil
	// 	}
	// }
	r := grpcApi.Wallet{
		WalletID: "",
		URL:      "",
		Gaslimit: 0,
		Response: &grpcApi.Response{
			Code:    404,
			Message: "cannot find wallet with the specific wallet id",
		},
	}
	wls = append(wls, &r)
	wellets = grpcApi.Wallets{
		Wallets: wls,
	}

	return &wellets, nil

}
