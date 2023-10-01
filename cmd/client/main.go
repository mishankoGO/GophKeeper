package main

import (
	"context"
	"github.com/mishankoGO/GophKeeper/config"
	"github.com/mishankoGO/GophKeeper/internal/client"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func main() {
	// init configuration
	conf, err := config.NewConfig("server_config.json")
	if err != nil {
		log.Fatal(err)
	}

	client, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cred := &pb.Credential{Login: "test_user", Password: "test_pass"}
	regResp, err := client.UsersClient.Register(ctx, &pb.RegisterRequest{Cred: cred})
	if err != nil {
		log.Fatal(err)
	}

	card := &pb.Card{Name: "test_card_new", Card: []byte("new card"), UpdatedAt: timestamppb.New(time.Now())}
	insert, err := client.CardsClient.Insert(ctx, &pb.InsertCardRequest{User: regResp.GetUser(), Card: card})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(insert)
}
