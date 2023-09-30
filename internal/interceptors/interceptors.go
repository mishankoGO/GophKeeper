// Package interceptors contains NewAuthInterceptor function to create interceptor.
// AuthInterceptor has Unary method that offer unary server interceptor.
// LoginSkip allows to skip authorization on login and registration handlers.
package interceptors

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/mishankoGO/GophKeeper/internal/security"
)

// AuthInterceptor collects jwt manager.
type AuthInterceptor struct {
	jwtManager *security.JWTManager // jwt manager
}

// NewAuthInterceptor creates authorization manager.
func NewAuthInterceptor(jwtManager *security.JWTManager) *AuthInterceptor {
	return &AuthInterceptor{jwtManager}
}

// LoginSkip function allows to skip login and registration handlers.
func LoginSkip(_ context.Context, c interceptors.CallMeta) bool {
	return c.FullMethod() != "/api.Credentials/Register" && c.FullMethod() != "/api.Users/Login"
}

// Unary method represents unary server interceptor.
func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		err := i.authorize(ctx)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

// authorize method verifies token.
func (i *AuthInterceptor) authorize(ctx context.Context) error {
	// get meta from context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	// get authorization value
	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	// get token
	accessToken := values[0]

	// verify token
	_, err := i.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return nil
}
