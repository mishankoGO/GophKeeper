package main

import (
	"context"
	"fmt"
	"github.com/mishankoGO/GophKeeper/config"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	db "github.com/mishankoGO/GophKeeper/internal/repository/postgres"
	"github.com/mishankoGO/GophKeeper/internal/server/handlers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func main() {
	conf, err := config.NewConfig("server_config.json")
	if err != nil {
		log.Fatal(status.Error(codes.Internal, err.Error()))
	}
	repo, err := db.NewDBRepository(conf)
	if err != nil {
		log.Fatal(err)
	}

	creds := handlers.Credentials{Repo: repo}
	credential := &pb.Credential{Login: "test_user", HashPassword: "test_password"}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	user := creds.Register(ctx, &pb.RegisterRequest{Cred: credential})
	fmt.Println(user)

	users := handlers.Users{Repo: repo}
	aaa := users.Login(ctx, &pb.LoginRequest{Cred: credential})

	fmt.Println(aaa)
}
