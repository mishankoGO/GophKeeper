package main

import (
	"context"
	"encoding/json"
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

	meta := map[string]string{"test_key": "test_val"}
	b, err := json.Marshal(meta)
	if err != nil {
		log.Fatal(err)
	}
	bf := &pb.BinaryFile{Name: "test_name", HashFile: "test_file", UpdatedAt: timestamppb.New(time.Now()), Meta: b}
	bfs := handlers.BinaryFiles{Repo: repo}

	res, err := bfs.Insert(ctx, &pb.InsertBinaryFileRequest{
		User: user.User,
		File: bf,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	resp, err := bfs.Get(ctx, &pb.GetRequest{User: user.User, Name: bf.Name})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

	newBF := &pb.BinaryFile{Name: "test_name", HashFile: "new_file", UpdatedAt: timestamppb.New(time.Now()), Meta: b}
	respp, err := bfs.Update(ctx, &pb.UpdateBinaryFileRequest{User: user.User, File: newBF})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(respp)

	//resppp, err := bfs.Delete(ctx, &pb.DeleteBinaryFileRequest{User: user.User, Name: "test_name"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(resppp)

	c := &pb.Card{Name: "test_name",
		CardNumber: "1234321",
		CardHolder: "test_holder",
		ExpiryDate: timestamppb.New(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		HashCvv:    "123",
		UpdatedAt:  timestamppb.New(time.Now()),
		Meta:       b}
	cs := handlers.Cards{Repo: repo}

	resc, err := cs.Insert(ctx, &pb.InsertCardRequest{
		User: user.User,
		Card: c,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resc)

	respc, err := cs.Get(ctx, &pb.GetRequest{User: user.User, Name: c.Name})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(respc)

	newC := &pb.Card{
		Name:       "test_name",
		CardNumber: "12343211",
		CardHolder: "new_test_holder",
		ExpiryDate: timestamppb.New(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		HashCvv:    "123",
		UpdatedAt:  timestamppb.New(time.Now()),
		Meta:       b}
	resppc, err := cs.Update(ctx, &pb.UpdateCardRequest{User: user.User, Card: newC})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resppc)

	//respppc, err := cs.Delete(ctx, &pb.DeleteCardRequest{User: user.User, Name: "test_name"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(respppc)

	lp := &pb.LogPass{Name: "test_name",
		Login:     "test_login",
		Pass:      "test_pass",
		UpdatedAt: timestamppb.New(time.Now()),
		Meta:      b}
	lps := handlers.LogPasses{Repo: repo}

	reslp, err := lps.Insert(ctx, &pb.InsertLogPassRequest{
		User:    user.User,
		LogPass: lp,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reslp)

	resplp, err := lps.Get(ctx, &pb.GetRequest{User: user.User, Name: c.Name})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resplp)

	newLP := &pb.LogPass{
		Name:      "test_name",
		Login:     "New login",
		Pass:      "new_pass",
		UpdatedAt: timestamppb.New(time.Now()),
		Meta:      b}
	respplp, err := lps.Update(ctx, &pb.UpdateLogPassRequest{User: user.User, LogPass: newLP})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(respplp)

	//resppplp, err := lps.Delete(ctx, &pb.DeleteLogPassRequest{User: user.User, Name: "test_name"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(resppplp)
}
