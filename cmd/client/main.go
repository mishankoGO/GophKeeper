package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mishankoGO/GophKeeper/config"
	"github.com/mishankoGO/GophKeeper/internal/cli"
	"github.com/mishankoGO/GophKeeper/internal/client"
	"github.com/mishankoGO/GophKeeper/internal/repository/bolt"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"log"
	"os"
)

const (
	keyPhrase = "secret"
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

	// init security
	security, err := security.NewSecurity(keyPhrase)
	if err != nil {
		log.Fatal(err)
	}

	client, err := client.NewClient(conf, repo, security)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	p := tea.NewProgram(cli.InitialModel(client))

	// Run returns the model as a tea.Model.

	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(*cli.Model); ok && m.Finish {
		err = m.Client.Sync(m.GetUser())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Bye!")
	}
	//fmt.Printf("Bye!")

	// register
	//cred := &pb.Credential{Login: "test_user", Password: "test_pass"}
	//regResp, err := client.UsersClient.Register(ctx, &pb.RegisterRequest{Cred: cred})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("register: ", regResp)

	// login
	//logResp, err := client.UsersClient.Login(ctx, &pb.LoginRequest{Cred: cred})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("login: ", logResp)

	// binary file
	//bf := &pb.BinaryFile{Name: "test binary file", File: []byte("new file"), UpdatedAt: timestamppb.New(time.Now())}
	//insertbf, err := client.BinaryFilesClient.Insert(ctx, &pb.InsertBinaryFileRequest{User: logResp.GetUser(), File: bf})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("insert bf: ", insertbf)

	// sync
	//user := converters.PBUserToUser(logResp.GetUser())
	//err = client.Sync(user)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//
	//bf = &pb.BinaryFile{Name: "new test binary file", File: []byte("new file"), UpdatedAt: timestamppb.New(time.Now())}
	//insertbf, err = client.BinaryFilesClient.Insert(ctx, &pb.InsertBinaryFileRequest{User: logResp.GetUser(), File: bf})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("insert bf: ", insertbf)
	//
	//getbf, err := client.BinaryFilesClient.Get(ctx, &pb.GetRequest{User: logResp.GetUser(), Name: "test binary file"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("get bf: ", string(getbf.File.File))

	//deletebf, err := client.BinaryFilesClient.Delete(ctx, &pb.DeleteBinaryFileRequest{User: logResp.GetUser(), Name: "test binary file"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("delete bf: ", deletebf)
	//
	//bf = &pb.BinaryFile{Name: "new test binary file", File: []byte("new new file"), UpdatedAt: timestamppb.New(time.Now())}
	//updatebf, err := client.BinaryFilesClient.Update(ctx, &pb.UpdateBinaryFileRequest{User: logResp.GetUser(), File: bf})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("update bf: ", updatebf)
	//
	//getbf, err = client.BinaryFilesClient.Get(ctx, &pb.GetRequest{User: logResp.GetUser(), Name: "new test binary file"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("get bf: ", string(getbf.File.File))

	//req := &pb.ListBinaryFileRequest{User: logResp.GetUser()}
	//serverbf, clientbf, err := client.BinaryFilesClient.List(ctx, req)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("server: ", serverbf)
	//fmt.Println("client: ", clientbf)
	//
	//// cards
	//card := &pb.Card{Name: "test card", Card: []byte("new card"), UpdatedAt: timestamppb.New(time.Now())}
	//insertcard, err := client.CardsClient.Insert(ctx, &pb.InsertCardRequest{User: logResp.GetUser(), Card: card})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("insert card: ", insertcard)
	//
	//card = &pb.Card{Name: "new test card", Card: []byte("new card"), UpdatedAt: timestamppb.New(time.Now())}
	//insertcard, err = client.CardsClient.Insert(ctx, &pb.InsertCardRequest{User: logResp.GetUser(), Card: card})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("insert card: ", insertcard)
	//
	//getcard, err := client.CardsClient.Get(ctx, &pb.GetRequest{Name: "test card"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("get card: ", getcard)
	//
	//card = &pb.Card{Name: "new test card", Card: []byte("new new card"), UpdatedAt: timestamppb.New(time.Now())}
	//updatecard, err := client.CardsClient.Update(ctx, &pb.UpdateCardRequest{User: logResp.GetUser(), Card: card})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("update card: ", updatecard)
	//
	//// texts
	//text := &pb.Text{Name: "test text", Text: []byte("new text"), UpdatedAt: timestamppb.New(time.Now())}
	//inserttext, err := client.TextsClient.Insert(ctx, &pb.InsertTextRequest{User: logResp.GetUser(), Text: text})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("insert text: ", inserttext)
	//
	//text = &pb.Text{Name: "new test text", Text: []byte("new text"), UpdatedAt: timestamppb.New(time.Now())}
	//inserttext, err = client.TextsClient.Insert(ctx, &pb.InsertTextRequest{User: logResp.GetUser(), Text: text})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("insert text: ", inserttext)
	//
	//gettext, err := client.TextsClient.Get(ctx, &pb.GetRequest{Name: "test text"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("get text: ", gettext)
	//
	//text = &pb.Text{Name: "new test text", Text: []byte("new new text"), UpdatedAt: timestamppb.New(time.Now())}
	//updatetext, err := client.TextsClient.Update(ctx, &pb.UpdateTextRequest{User: logResp.GetUser(), Text: text})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("update text: ", updatetext)
	//
	//// logpasses
	//log_pass := &pb.LogPass{Name: "test log_pass", Login: []byte("new login"), Pass: []byte("new pass"), UpdatedAt: timestamppb.New(time.Now())}
	//insertlogpass, err := client.LogPassesClient.Insert(ctx, &pb.InsertLogPassRequest{User: logResp.GetUser(), LogPass: log_pass})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("insert log_pass: ", insertlogpass)
	//
	//log_pass = &pb.LogPass{Name: "new test log_pass", Login: []byte("new login"), Pass: []byte("new pass"), UpdatedAt: timestamppb.New(time.Now())}
	//insertlogpass, err = client.LogPassesClient.Insert(ctx, &pb.InsertLogPassRequest{User: logResp.GetUser(), LogPass: log_pass})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("insert log_pass: ", insertlogpass)
	//
	//getlogpass, err := client.LogPassesClient.Get(ctx, &pb.GetRequest{Name: "test log_pass"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("get log_pass: ", getlogpass)
	//
	//log_pass = &pb.LogPass{Name: "new test log_pass", Login: []byte("new new login"), Pass: []byte("new new pass"), UpdatedAt: timestamppb.New(time.Now())}
	//updatelogpass, err := client.LogPassesClient.Update(ctx, &pb.UpdateLogPassRequest{User: logResp.GetUser(), LogPass: log_pass})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("update log_pass: ", updatelogpass)
}
