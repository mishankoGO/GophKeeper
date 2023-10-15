package main

import (
	"github.com/mishankoGO/GophKeeper/config"
	"github.com/mishankoGO/GophKeeper/internal/repository/postgres"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"github.com/mishankoGO/GophKeeper/internal/server"
	"log"
	"time"
)

const (
	keyPhrase     = "secret"         // key phrase for data encryption
	secretKey     = "secret"         // secret key for jwt
	tokenDuration = 30 * time.Minute // duration of the token
)

func main() {
	// init configuration
	conf, err := config.NewConfig("server_config.json")
	if err != nil {
		log.Fatal(err)
	}

	// init repository
	repo, err := postgres.NewDBRepository(conf)
	if err != nil {
		log.Fatal(err)
	}

	// init jwt manager
	jwtManager := security.NewJWTManager(secretKey, tokenDuration)

	// init security
	security, err := security.NewSecurity(keyPhrase)
	if err != nil {
		log.Fatal(err)
	}

	// create server
	grpcServer := server.NewServer(repo, jwtManager, security, conf)

	// run server
	log.Println("Running server...")
	err = grpcServer.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
