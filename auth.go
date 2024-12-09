package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc/metadata"
)

// AuthInterceptor checks the token in the gRPC metadata.
func AuthInterceptor(ctx context.Context) error {
	validToken := os.Getenv("AUTH_TOKEN")
	if validToken == "" {
		return fmt.Errorf("missing environment variable: AUTH_TOKEN")
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing metadata")
	}
	if tokens, ok := md["authorization"]; ok && len(tokens) > 0 && tokens[0] == validToken {
		return nil
	}
	return fmt.Errorf("invalid token")
}
