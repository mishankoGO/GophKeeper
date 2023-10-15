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
	"fmt"
	"github.com/mishankoGO/GophKeeper/internal/client/interfaces"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/pkg/hash"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UsersClient struct {
	token              string
	repo               interfaces.Repository
	usersService       pb.UsersClient
	credentialsService pb.CredentialsClient
	offline            bool
}

func NewUsersClient(cc *grpc.ClientConn, repo interfaces.Repository) *UsersClient {
	if cc != nil {
		usersService := pb.NewUsersClient(cc)
		credentialsService := pb.NewCredentialsClient(cc)
		return &UsersClient{usersService: usersService, credentialsService: credentialsService, repo: repo, offline: false}
	}
	return &UsersClient{offline: true, repo: repo}
}

func (u *UsersClient) GetToken() string {
	return u.token
}

func (u *UsersClient) Close() error {
	return u.repo.Close()
}

// Login method logs in user by its credentials.
func (u *UsersClient) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if u.offline {
		cred, user, err := u.repo.Login(req.GetCred().GetLogin())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error during login: %v", err)
		}

		// hash password
		hashPass := hash.HashPass([]byte(req.GetCred().GetPassword()))

		// check if password is valid
		if cred.Password != hashPass {
			return nil, status.Errorf(codes.NotFound, "incorrect username/password: %v", err)
		}

		pbUser := converters.UserToPBUser(user)
		return &pb.LoginResponse{User: pbUser}, nil
	}
	resp, err := u.usersService.Login(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error during login: %v", err)
	}

	u.token = resp.GetToken()

	return resp, nil
}

// Register method registers new user with its credentials.
func (u *UsersClient) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if u.offline {
		return nil, fmt.Errorf("offline registration is not available")
	}

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

	cred := converters.PBCredentialToCredential(req.GetCred())
	user := converters.PBUserToUser(lResp.GetUser())

	cred.Password = hash.HashPass([]byte(cred.Password))
	err = u.repo.InsertUser(cred, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting user to offline db: %v", err)
	}

	return resp, err
}
