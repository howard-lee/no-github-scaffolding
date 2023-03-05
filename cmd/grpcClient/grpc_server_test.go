package main

import (
	grpcApi "api"
	"context"
	"fmt"
	"grpcData"
	"log"
	"net"
	"server"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

//var mockStore *grpcData.MockStore
var cachedWStore *grpcData.CachedWalletStore

func setupTestX(d interface{}) *grpcData.GetWallets {
	//mockStore = &grpcData.MockStore{}

	cachedWStore = &grpcData.CachedWalletStore{}

	h := grpcData.NewGetWallets(cachedWStore)

	return h

}

func TestGrpcServerGetWalletsWithValidParameters(t *testing.T) {
	fmt.Println("\n ... TestGrpcServerGetWalletsWithValidParameters invoked")

	// integrated test using mocked data DI into the server

	var testServerURL = "localhost:13436"
	go func(url string) {
		lis, err := net.Listen("tcp", url)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		defer lis.Close()

		log.Printf("Listening on %s", url)

		cachedStore := &grpcData.CachedWalletStore{}

		h := grpcData.NewGetWallets(cachedStore)
		_ = h
		//srv := server.NewGRPCBlkWalletServer(h)
		srv := server.NewGRPCBlkWalletServer(h)

		// Register reflection service on gRPC server.
		// reflection.Register(srv)
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}(testServerURL)

	// --- above server code

	matrix := [][]string{
		{"get with valid input, find result with return code 0", "135246", "0"},   /*  initializers for row indexed by 0 */
		{"get with valid input, find result with return code 404", "1212", "404"}, /*  initializers for row indexed by 1 */
		{"get empty input, find many results with return code 0", "", "0"},        /*  initializers for row indexed by 2 */
	}

	for _, v := range matrix {

		var conn *grpc.ClientConn
		conn, _ = grpc.Dial(testServerURL, grpc.WithInsecure())
		defer conn.Close()

		client := grpcApi.NewBlkWalletClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		ar := grpcApi.AccountRequest{Account: v[1]}
		wals, err := client.GetWallets(ctx, &ar)
		if err != nil {
			er := fmt.Errorf("error encountered calling GetWallets %v \n", err)
			t.Error(er)
		}
		//fmt.Println("wals : ", wals)
		cdi := strconv.FormatUint(uint64(wals.Wallets[0].Response.Code), 10)
		if cdi != v[2] {
			t.Errorf("error in getting wallet,  value got %v, want %v", v[2], cdi)
		} else {
			t.Logf("get wallet operation ok,  got %v, want %v", v[2], cdi)
		}
		//fmt.Println(" wals.Wallets[0] : ", wals.Wallets[0])
	}

}

func TestGetWalletMethodCallsDataStoreWithDifferentQueries(t *testing.T) {
	fmt.Println("\n ... TestGetWalletMethodCallsDataStoreWithDifferentQueries invoked")

	// unit test using mocked data DI into the server

	it1 := setupTestX("dummy now")

	matrix := [][]string{
		{"get with valid input, find result with return code 0", "135246", "0"},   /*  initializers for row indexed by 0 */
		{"get with valid input, find result with return code 404", "1212", "404"}, /*  initializers for row indexed by 1 */
		{"get empty input, find many results with return code 0", "", "0"},        /*  initializers for row indexed by 2 */
	}

	for _, v := range matrix {
		ret := it1.DataStore.GetWallets(v[1])

		cdi := strconv.FormatUint(uint64(ret[0].Response.Code), 10)
		if cdi != v[2] {
			t.Errorf("error in getting wallet,  value got %v, want %v", v[2], cdi)
		} else {
			t.Logf("get wallet operation ok,  got %v, want %v", v[2], cdi)
		}
	}
}

// 		srv := server.NewGRPCBlkWalletServer()
// 		// Register reflection service on gRPC server.
// 		// reflection.Register(srv)
// 		if err := srv.Serve(lis); err != nil {
// 			log.Fatalf("failed to serve: %v", err)
// 		}
// 	}(walletURL)

func setupTestY(t *testing.T, fn func()) (
	client grpcApi.BlkClient,
	teardown func(),
) {
	t.Helper()

	l, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	clientOptions := []grpc.DialOption{grpc.WithInsecure()}
	cc, err := grpc.Dial(l.Addr().String(), clientOptions...)
	require.NoError(t, err)

	//server, err := server. NewGRPCServer(cfg)
	server := server.NewGRPCBlkServer()
	require.NoError(t, err)

	go func() {
		server.Serve(l)
	}()

	client = grpcApi.NewBlkClient(cc)

	return client, func() {
		server.Stop()
		cc.Close()
		l.Close()
		//clog.Remove()
	}
}
