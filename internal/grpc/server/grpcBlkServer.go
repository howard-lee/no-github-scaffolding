package server

import (
	grpcApi "api"
	"context"
	"fmt"
	"grpcData"
	"time"

	"google.golang.org/grpc"
)

// var _ api1.Activity = (*grpcServer)(nil)

type grpcBlkServer struct {
	grpcApi.UnimplementedBlkServer
	Blk
}

func NewGRPCBlkServer() *grpc.Server {

	gsrv := grpc.NewServer()

	srv := grpcBlkServer{
		//Activities: acc,
	}
	grpcApi.RegisterBlkServer(gsrv, &srv)
	return gsrv
}

func (s *grpcBlkServer) GetBalance(context context.Context, req *grpcApi.AccountRequest) (*grpcApi.Balance, error) {

	fmt.Println("grpc GetBalance invoked account : ", req.Account, ", time: ", time.Now())
	bal, err := s.Blk.GetBalance(req.Account)

	return bal, err
}

type grpcBlkWalletServer struct {
	grpcApi.UnimplementedBlkWalletServer
	//BlkWallet
	GenericGetWallet interface{}
}

func NewGRPCBlkWalletServer(g interface{}) *grpc.Server {

	gsrv := grpc.NewServer()
	srv := grpcBlkWalletServer{}

	cachedWStore := &grpcData.CachedWalletStore{}
	//h := grpcData.NewGetWallets(cachedWStore)
	srv.GenericGetWallet = cachedWStore
	//srv := BlkWallet{}
	//srv.GenericGetWallet = g

	grpcApi.RegisterBlkWalletServer(gsrv, &srv)
	return gsrv
}

func (s *grpcBlkWalletServer) GetWallets(context context.Context, req *grpcApi.AccountRequest) (*grpcApi.Wallets, error) {

	itc := s.GenericGetWallet

	var bal []grpcApi.Wallet

	//TODO : convert the code below to code that can easily extended to other store dynamically
	//   using reflection
	switch v := itc.(type) {
	case WalletStore:
		h := itc.(WalletStore)
		bal = h.GetWallets(req.Account)
	case grpcData.MySQLStore:
		h := itc.(grpcData.MySQLStore)
		bal = h.GetWallets(req.Account)
	default:
		fmt.Printf("unknown %v \n", v)
	}
	// TODO : end

	//bal, err := s.BlkWallet.GetWallets(req.Account)

	// convert from []grpcApi.Wallet to grpcApi.Wallets
	wal := grpcApi.Wallets{}

	for _, v := range bal {
		if v.WalletID != "" {
			v.Response = &grpcApi.Response{
				Code:    0,
				Message: "data found.",
			}
			wal.Wallets = append(wal.Wallets, &v)
		} else {
			v.Response = &grpcApi.Response{
				Code:    404,
				Message: "data not found.",
			}
			wal.Wallets = append(wal.Wallets, &v)
		}

	}
	return &wal, nil
}
