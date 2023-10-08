// Package handlers contains servers interfaces.
// The list of servers:
//     Users, Credentials, BinaryFiles, Cards, Texts, LogPasses
package handlers

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"github.com/mishankoGO/GophKeeper/internal/server/interfaces"
	"github.com/mishankoGO/GophKeeper/pkg/hash"
)

// Users contains repository and jwt manager.
type Users struct {
	pb.UnimplementedUsersServer
	Repo       interfaces.Repository // repository
	jwtManager *security.JWTManager  // jwt manager
}

// NewUsers function creates new users server.
func NewUsers(repo interfaces.Repository, jwtManager *security.JWTManager) *Users {
	return &Users{Repo: repo, jwtManager: jwtManager}
}

// Login method logins user.
func (u *Users) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// get credentials
	pbCred := req.GetCred()

	// convert proto cred to model cred
	cred := converters.PBCredentialToCredential(pbCred)

	// get login and password
	login := cred.Login
	password := cred.Password

	// get user cred from db
	cred, user, err := u.Repo.Login(ctx, login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting user credentials: %v", err)
	}

	// hash password
	hashPass := hash.HashPass([]byte(password))

	// check if password is valid
	if cred.Password != hashPass {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password: %v", err)
	}
	log.Println("Logged in successfully!")

	// generate jwt token
	token, err := u.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token: %v", err)
	}

	// convert model user to proto user
	pbUser := converters.UserToPBUser(user)

	// create response
	var res = &pb.LoginResponse{User: pbUser, Token: token}

	return res, nil
}
