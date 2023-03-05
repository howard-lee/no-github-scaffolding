package main

import (
	grpcApi "api"

	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultURL = "localhost:8080"

type BlkClient struct {
	client grpcApi.BlkClient
}

func NewBlkClient(URL string) BlkClient {
	conn, err := grpc.Dial(URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := grpcApi.NewBlkClient(conn)
	return BlkClient{client: client}
}

func block_01() {
	blkClient := NewBlkClient(defaultURL)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ar := grpcApi.AccountRequest{Account: "135246"}

	a, err := blkClient.client.GetBalance(ctx, &ar)
	if err != nil {
		er := fmt.Errorf("error encountered calling GetBalance %v \n", err)
		fmt.Println(er)
	}

	fmt.Printf("Get Balance completed successfully, account = %v, balance = %v , response code = %v, message=%v \n", a.Account, a.Balance, a.Response.Code, a.Response.Message)
}

func main() {
	fmt.Println("grpcBlkClient main invoked ", time.Now())

	block_01()

}
