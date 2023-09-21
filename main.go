package main

import (
	"fmt"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc/registration"
	db "github.com/mishankoGO/GophKeeper/internal/repository/postgres"
	"github.com/mishankoGO/GophKeeper/internal/server"
	"log"
)

func main() {
	repo, err := db.NewDBRepository()
	if err != nil {
		log.Fatal(err)
	}

	creds := server.Credentials{Repo: repo}
	credential := pb.Credential{Login: "test_user", HashPassword: "test_password"}
	user, err := creds.Register(&credential)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
