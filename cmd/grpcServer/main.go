package main

import (
	"fmt"
	"log"
	"net"
	"server"
	"time"
)

func main() {
	fmt.Println("grpcBlkServer invoked ", time.Now())

	port := ":8080"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	log.Printf("Listening on %s", port)
	srv := server.NewGRPCBlkServer()
	// Register reflection service on gRPC server.
	// reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
