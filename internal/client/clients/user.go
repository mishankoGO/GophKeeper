// Package clients contains clients for all the operations.
// It has:
// CredentialsClient for registration.
// UsersClient for login in.
// CardsClient to operate with bank cards.
// LogPassesClient to operate with log passes.
// TextsClient to operate with texts.
// BinaryFilesClient to operate with binary files.
package clients

import (
	"context"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UsersClient struct {
	token              string
	usersService       pb.UsersClient
	credentialsService pb.CredentialsClient
}

func NewUsersClient(cc *grpc.ClientConn) *UsersClient {
	usersService := pb.NewUsersClient(cc)
	credentialsService := pb.NewCredentialsClient(cc)
	return &UsersClient{usersService: usersService, credentialsService: credentialsService}
}

func (u *UsersClient) GetToken() string {
	return u.token
}

// Login method logs in user by its credentials.
func (u *UsersClient) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	resp, err := u.usersService.Login(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error during login: %v", err)
	}

	u.token = resp.GetToken()

	return resp, nil
}

// Register method registers new user with its credentials.
func (u *UsersClient) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// send request
	resp, err := u.credentialsService.Register(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error during registration: %v", err)
	}

	lResp, err := u.Login(ctx, &pb.LoginRequest{Cred: req.GetCred()})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error login after registration: %v", err)
	}

	u.token = lResp.GetToken()

	return resp, err
}
