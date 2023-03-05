package server

import (
	grpcApi "api"
	"fmt"
)

type AccountStore interface {
	Search(name string) []grpcApi.Account
}

// type BalanceStore interface {
// 	Search(name string) []grpcApi.Balance
// }

type Blk struct {
	grpcApi.Balances

	store []grpcApi.Balance
}

func NewBlk() (*Blk, error) {
	return &Blk{}, nil
}

func (c *Blk) GetBalance(ac string) (*grpcApi.Balance, error) {

	fmt.Println("store side GetBalance invoked. ", c.store)
	//result := true
	if ac != "" {
		for _, v := range c.store {
			var bl grpcApi.Balance
			if v.Account == ac {
				bl = grpcApi.Balance{
					Account: v.Account,
					Balance: v.Balance,
					Response: &grpcApi.Response{
						Code:    0,
						Message: "get suessfully complete",
					},
				}
			}
			return &bl, nil
		}
	}

	//if rslt == false {
	r := grpcApi.Balance{
		Account: "",
		Balance: 0,
		Response: &grpcApi.Response{
			Code:    404,
			Message: "id should be non-zero value",
		},
	}
	return &r, nil
}
