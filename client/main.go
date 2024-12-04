package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	pb "server_agent/module/proto" // Update with actual proto path

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	// Load TLS credentials
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		RootCAs: certPool,
	})

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewResourceCheckerClient(conn)

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "my-secure-token")
	res, err := client.CheckResources(ctx, &pb.ResourceRequest{})
	if err != nil {
		log.Fatalf("failed to fetch resources: %v", err)
	}
	fmt.Printf("Resource status: %+v\n", res)
}
