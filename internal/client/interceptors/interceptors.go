// Package interceptors contains AuthInterceptor that supports authorization with jwt tokens.
package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/mishankoGO/GophKeeper/internal/client/clients"
)

// AuthInterceptor contains auth client and token.
type AuthInterceptor struct {
	authClient  *clients.UsersClient
	accessToken string
}

// NewAuthInterceptor creates new instance of AuthInterceptor.
func NewAuthInterceptor(authClient *clients.UsersClient) (*AuthInterceptor, error) {
	interceptor := &AuthInterceptor{
		authClient: authClient,
	}

	return interceptor, nil
}

// getToken method returns token.
func (i *AuthInterceptor) getToken() string {
	i.accessToken = i.authClient.GetToken()
	return i.accessToken
}

//func (i *AuthInterceptor) scheduleRefreshToken(refreshDuration time.Duration) error {
//	err := i.refreshToken()
//	if err != nil {
//		return fmt.Errorf("error refreshing token: %w", err)
//	}
//
//	go func() {
//		wait := refreshDuration
//		for {
//			time.Sleep(wait)
//			err := i.refreshToken()
//			if err != nil {
//				wait = time.Second
//			} else {
//				wait = refreshDuration
//			}
//		}
//	}()
//
//	return nil
//}

// Unary method represents unary client interceptor.
func (i *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {

		// get token
		i.getToken()
		return invoker(i.attachToken(ctx), method, req, reply, cc, opts...)
	}
}

// attachToken method attaches token to outgoing context.
func (i *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", i.accessToken)
}
