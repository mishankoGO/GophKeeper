package main

import (
	"context"
	"github.com/mishankoGO/GophKeeper/config"
	"github.com/mishankoGO/GophKeeper/internal/client"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/repository/bolt"
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
	repo, err := bolt.NewDBRepository(conf)
	if err != nil {
		log.Fatal(err)
	}

	client, err := client.NewClient(conf, repo)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cred := &pb.Credential{Login: "test_user", Password: "test_pass"}
	//regResp, err := client.UsersClient.Register(ctx, &pb.RegisterRequest{Cred: cred})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("register: ", regResp)

	logResp, err := client.UsersClient.Login(ctx, &pb.LoginRequest{Cred: cred})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("login: ", logResp)

	bf := &pb.BinaryFile{Name: "new test binary file", File: []byte("new file"), UpdatedAt: timestamppb.New(time.Now())}
	insert, err := client.BinaryFilesClient.Insert(ctx, &pb.InsertBinaryFileRequest{User: logResp.GetUser(), File: bf})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("insert bf: ", insert)

	get, err := client.BinaryFilesClient.Get(&pb.GetRequest{Name: bf.GetName()})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("get bf: ", get)
	//card := &pb.Card{Name: "test_card_new", Card: []byte("new card"), UpdatedAt: timestamppb.New(time.Now())}
	//insert, err := client.CardsClient.Insert(ctx, &pb.InsertCardRequest{User: regResp.GetUser(), Card: card})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("insert card: ", insert)
	//
	//get, err := client.CardsClient.Get(ctx, &pb.GetRequest{Name: "test_card_new", User: regResp.GetUser()})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("get card: ", get)
	//
	//card = &pb.Card{Name: "test_card_new", Card: []byte("new new card"), UpdatedAt: timestamppb.New(time.Now())}
	//update, err := client.CardsClient.Update(ctx, &pb.UpdateCardRequest{Name: "test_card_new", User: regResp.GetUser(), Card: card})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("update card: ", update)
	//
	//delete, err := client.CardsClient.Delete(ctx, &pb.DeleteCardRequest{Name: "test_card_new", User: regResp.GetUser()})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("delete card: ", delete)
	//
	//text := &pb.Text{Name: "test_text_new", Text: []byte("new text"), UpdatedAt: timestamppb.New(time.Now())}
	//insert, err = client.TextsClient.Insert(ctx, &pb.InsertTextRequest{User: regResp.GetUser(), Text: text})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("insert text: ", insert)
	//
	//gettext, err := client.TextsClient.Get(ctx, &pb.GetRequest{Name: "test_text_new", User: regResp.GetUser()})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("get text: ", gettext)
	//
	//text = &pb.Text{Name: "test_text_new", Text: []byte("new new text"), UpdatedAt: timestamppb.New(time.Now())}
	//updatetext, err := client.TextsClient.Update(ctx, &pb.UpdateTextRequest{Name: "test_text_new", User: regResp.GetUser(), Text: text})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("update text: ", updatetext)
	//
	//delete, err = client.TextsClient.Delete(ctx, &pb.DeleteTextRequest{Name: "test_text_new", User: regResp.GetUser()})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println("delete text: ", delete)
}
