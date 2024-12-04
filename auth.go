package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

const validToken = "my-secure-token"

// AuthInterceptor checks the token in the gRPC metadata.
func AuthInterceptor(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing metadata")
	}
	if tokens, ok := md["authorization"]; ok && len(tokens) > 0 && tokens[0] == validToken {
		return nil
	}
	return fmt.Errorf("invalid token")
}
