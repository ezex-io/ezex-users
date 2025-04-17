// package main is the entry point for the application.
package main

import (
	"context"
	"log"
	"time"

	"github.com/ezex-io/ezex-users/api/grpc/proto"
	"github.com/ezex-io/ezex-users/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	conn, err := grpc.Dial(cfg.GRPCServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close connection: %v", err)
		}
	}()

	client := proto.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	saveResp, err := client.SaveSecurityImage(ctx, &proto.SaveSecurityImageRequest{
		UserId:         "test-user",
		SecurityImage:  "test-image-data",
		SecurityPhrase: "test-image-phrase",
	})
	if err != nil {
		log.Printf("SaveSecurityImage test failed (expected): %v", err)
	} else {
		log.Printf("SaveSecurityImage test succeeded: %v", saveResp)
	}

	getResp, err := client.GetSecurityImage(ctx, &proto.GetSecurityImageRequest{
		UserId: "test-user",
	})
	if err != nil {
		log.Printf("GetSecurityImage test failed: %v", err)

		return
	}

	log.Printf("GetSecurityImage test succeeded: %v", getResp)
}
