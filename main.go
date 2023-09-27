package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mishankoGO/GophKeeper/config"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	db "github.com/mishankoGO/GophKeeper/internal/repository/postgres"
	"github.com/mishankoGO/GophKeeper/internal/security"
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
	sec, err := security.NewSecurity("key phrase phrase key")
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	creds := handlers.Credentials{Repo: repo}
	cred := &pb.Credential{Login: "test_user", Password: "test_password"}
	user := creds.Register(ctx, &pb.RegisterRequest{Cred: cred})
	users := handlers.Users{Repo: repo}
	aaa := users.Login(ctx, &pb.LoginRequest{Cred: cred})
	fmt.Println(aaa)

	meta := map[string]string{"test_key": "test_val"}
	b, err := json.Marshal(meta)
	if err != nil {
		log.Fatal(err)
	}
	bf := &pb.BinaryFile{Name: "test_name", File: []byte("test_file"), UpdatedAt: timestamppb.New(time.Now()), Meta: b}
	bfs := handlers.NewBinaryFiles(repo, sec)

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
	//
	//newBF := &pb.BinaryFile{Name: "test_name", HashFile: "new_file", UpdatedAt: timestamppb.New(time.Now()), Meta: b}
	//respp, err := bfs.Update(ctx, &pb.UpdateBinaryFileRequest{User: user.User, File: newBF})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(respp)
	//
	////resppp, err := bfs.Delete(ctx, &pb.DeleteBinaryFileRequest{User: user.User, Name: "test_name"})
	////if err != nil {
	////	log.Fatal(err)
	////}
	////
	////fmt.Println(resppp)
	//
	//c := &pb.Card{Name: "test_name",
	//	CardNumber: "1234321",
	//	CardHolder: "test_holder",
	//	ExpiryDate: timestamppb.New(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
	//	HashCvv:    "123",
	//	UpdatedAt:  timestamppb.New(time.Now()),
	//	Meta:       b}
	//cs := handlers.Cards{Repo: repo}
	//
	//resc, err := cs.Insert(ctx, &pb.InsertCardRequest{
	//	User: user.User,
	//	Card: c,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(resc)
	//
	//respc, err := cs.Get(ctx, &pb.GetRequest{User: user.User, Name: c.Name})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(respc)
	//
	//newC := &pb.Card{
	//	Name:       "test_name",
	//	CardNumber: "12343211",
	//	CardHolder: "new_test_holder",
	//	ExpiryDate: timestamppb.New(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
	//	HashCvv:    "123",
	//	UpdatedAt:  timestamppb.New(time.Now()),
	//	Meta:       b}
	//resppc, err := cs.Update(ctx, &pb.UpdateCardRequest{User: user.User, Card: newC})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(resppc)
	//
	////respppc, err := cs.Delete(ctx, &pb.DeleteCardRequest{User: user.User, Name: "test_name"})
	////if err != nil {
	////	log.Fatal(err)
	////}
	////
	////fmt.Println(respppc)
	//
	//lp := &pb.LogPass{Name: "test_name",
	//	Login:     "test_login",
	//	Pass:      "test_pass",
	//	UpdatedAt: timestamppb.New(time.Now()),
	//	Meta:      b}
	//lps := handlers.LogPasses{Repo: repo}
	//
	//reslp, err := lps.Insert(ctx, &pb.InsertLogPassRequest{
	//	User:    user.User,
	//	LogPass: lp,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(reslp)
	//
	//resplp, err := lps.Get(ctx, &pb.GetRequest{User: user.User, Name: c.Name})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(resplp)
	//
	//newLP := &pb.LogPass{
	//	Name:      "test_name",
	//	Login:     "New login",
	//	Pass:      "new_pass",
	//	UpdatedAt: timestamppb.New(time.Now()),
	//	Meta:      b}
	//respplp, err := lps.Update(ctx, &pb.UpdateLogPassRequest{User: user.User, LogPass: newLP})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(respplp)
	//
	////resppplp, err := lps.Delete(ctx, &pb.DeleteLogPassRequest{User: user.User, Name: "test_name"})
	////if err != nil {
	////	log.Fatal(err)
	////}
	////
	////fmt.Println(resppplp)
	//
	//t := &pb.Text{Name: "test_name",
	//	HashText:  "test_text",
	//	UpdatedAt: timestamppb.New(time.Now()),
	//	Meta:      b}
	//ts := handlers.Texts{Repo: repo}
	//
	//rest, err := ts.Insert(ctx, &pb.InsertTextRequest{
	//	User: user.User,
	//	Text: t,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(rest)
	//
	//respt, err := ts.Get(ctx, &pb.GetRequest{User: user.User, Name: c.Name})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(respt)
	//
	//newT := &pb.Text{
	//	Name:      "test_name",
	//	HashText:  "New text",
	//	UpdatedAt: timestamppb.New(time.Now()),
	//	Meta:      b}
	//resppt, err := ts.Update(ctx, &pb.UpdateTextRequest{User: user.User, Text: newT})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(resppt)
	//
	////respppt, err := ts.Delete(ctx, &pb.DeleteTextRequest{User: user.User, Name: "test_name"})
	////if err != nil {
	////	log.Fatal(err)
	////}
	////
	////fmt.Println(respppt)
	//
	//key, err := generator.New()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = key.SaveKeys()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
