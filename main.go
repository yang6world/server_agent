package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	pb "server_agent/module/proto" // Update with actual proto path

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ResourceCheckerServer struct {
	pb.UnimplementedResourceCheckerServer
}

func (s *ResourceCheckerServer) CheckResources(ctx context.Context, req *pb.ResourceRequest) (*pb.ResourceResponse, error) {
	// Auth validation
	if err := AuthInterceptor(ctx); err != nil {
		return nil, err
	}

	// Fetch resources
	data, err := CheckResources()
	if err != nil {
		return nil, err
	}

	return &pb.ResourceResponse{
		Hostname:          data["hostname"].(string),
		Os:                data["os"].(string),
		KernelVersion:     data["kernel_version"].(string),
		CpuUsage:          data["cpu_usage"].(float64),
		MemoryUsage:       data["memory_usage"].(float64),
		SwapUsage:         data["swap_usage"].(float64),
		DiskUsage:         data["disk_usage"].(string),
		LoadAverage:       data["load_average"].(float64),
		WebshellSupported: data["webshell_supported"].(bool),
	}, nil
}

func main() {
	// Load TLS credentials
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("failed to load TLS keys: %v", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certPool,
	})

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterResourceCheckerServer(grpcServer, &ResourceCheckerServer{})
	fmt.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
