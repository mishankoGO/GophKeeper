// Package server contains NewServer function to create new grpc server instance.
// Server struct has Serve method to run grpc server.
package server

import (
	"fmt"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/mishankoGO/GophKeeper/config"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/interceptors"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"github.com/mishankoGO/GophKeeper/internal/server/handlers"
	"github.com/mishankoGO/GophKeeper/internal/server/interfaces"
)

// Server struct to communicate with server
type Server struct {
	server *grpc.Server   // grpc server
	conf   *config.Config // configuration file
}

// NewServer function creates new server instance.
func NewServer(
	repo interfaces.Repository,
	jwtManager *security.JWTManager,
	security *security.Security,
	config *config.Config) *Server {

	// init interceptor
	interceptor := interceptors.NewAuthInterceptor(jwtManager)

	// init grpc server
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(selector.UnaryServerInterceptor(interceptor.Unary(), selector.MatchFunc(interceptors.LoginSkip))),
	)

	// init and register credential server
	credServer := handlers.NewCredentials(repo)
	pb.RegisterCredentialsServer(grpcServer, credServer)

	// init and register authorization server
	authServer := handlers.NewUsers(repo, jwtManager)
	pb.RegisterUsersServer(grpcServer, authServer)

	// init and register binary files server
	bfServer := handlers.NewBinaryFiles(repo, security)
	pb.RegisterBinaryFilesServer(grpcServer, bfServer)

	// init and register card server
	cardServer := handlers.NewCards(repo, security)
	pb.RegisterCardsServer(grpcServer, cardServer)

	// init and register log pass server
	lpServer := handlers.NewLogPasses(repo, security)
	pb.RegisterLogPassesServer(grpcServer, lpServer)

	// init and register text server
	textServer := handlers.NewTexts(repo, security)
	pb.RegisterTextsServer(grpcServer, textServer)

	// activate reflection to test server in evans
	reflection.Register(grpcServer)

	return &Server{server: grpcServer, conf: config}
}

// Serve method runs grpc server.
func (s *Server) Serve() error {
	// init listener
	listener, err := net.Listen("tcp", s.conf.Address)
	if err != nil {
		return fmt.Errorf("error creating listener: %w", err)
	}

	// run server
	err = s.server.Serve(listener)
	if err != nil {
		return fmt.Errorf("error running a server: %w", err)
	}

	return nil
}
