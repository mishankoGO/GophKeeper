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
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
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
	cred := &pb.Credential{Login: "test_user", HashPassword: "test_password"}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	user := creds.Register(ctx, &pb.RegisterRequest{Cred: cred})
	fmt.Println(user)

	users := handlers.Users{Repo: repo}
	aaa := users.Login(ctx, &pb.LoginRequest{Cred: cred})

	fmt.Println(aaa)

	bf := &pb.BinaryFile{Name: "test_name", HashFile: "test_file", UpdatedAt: timestamppb.New(time.Now())}
	bfs := handlers.BinaryFiles{Repo: repo}

	res, err := bfs.Insert(ctx, &pb.InsertBinaryFileRequest{
		User: user.User,
		File: bf,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
