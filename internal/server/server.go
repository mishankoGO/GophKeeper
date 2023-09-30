package server

import (
	"fmt"
	"github.com/mishankoGO/GophKeeper/config"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"github.com/mishankoGO/GophKeeper/internal/server/handlers"
	"github.com/mishankoGO/GophKeeper/internal/server/interfaces"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	server *grpc.Server
	conf   *config.Config
}

func NewServer(
	repo interfaces.Repository,
	jwtManager *security.JWTManager,
	security *security.Security,
	config *config.Config) *Server {

	grpcServer := grpc.NewServer()

	credServer := handlers.NewCredentials(repo)
	pb.RegisterCredentialsServer(grpcServer, credServer)

	authServer := handlers.NewUsers(repo, jwtManager)
	pb.RegisterUsersServer(grpcServer, authServer)

	bfServer := handlers.NewBinaryFiles(repo, security)
	pb.RegisterBinaryFilesServer(grpcServer, bfServer)

	cardServer := handlers.NewCards(repo, security)
	pb.RegisterCardsServer(grpcServer, cardServer)

	lpServer := handlers.NewLogPasses(repo, security)
	pb.RegisterLogPassesServer(grpcServer, lpServer)

	textServer := handlers.NewTexts(repo, security)
	pb.RegisterTextsServer(grpcServer, textServer)

	reflection.Register(grpcServer)
	return &Server{server: grpcServer, conf: config}
}

func (s *Server) Serve() error {
	listener, err := net.Listen("tcp", s.conf.Address)
	if err != nil {
		return fmt.Errorf("error creating listener: %w", err)
	}

	err = s.server.Serve(listener)
	if err != nil {
		return fmt.Errorf("error running a server: %w", err)
	}

	return nil
}
