package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mishankoGO/GophKeeper/config"
	"github.com/mishankoGO/GophKeeper/internal/cli"
	"github.com/mishankoGO/GophKeeper/internal/client"
	"github.com/mishankoGO/GophKeeper/internal/repository/bolt"
	"log"
	"os"
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
	defer client.Close()

	p := tea.NewProgram(cli.InitialModel(client))

	// Run returns the model as a tea.Model.

	_, err = p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	fmt.Printf("Bye!")

	//	if m.Choice == "Login" {
	//		for {
	//			m, err := tea.NewProgram(login.InitialModel(client), tea.WithAltScreen()).Run()
	//			if err != nil {
	//				fmt.Printf("could not start program: %s\n", err)
	//				os.Exit(1)
	//			}
	//			// Assert the final tea.Model to our local model and print the choice.
	//			if m, ok := m.(login.LoginModel); ok && m.Finish {
	//				break
	//			}
	//		}
	//
	//	} else if m.Choice == "Register" {
	//		if _, err := tea.NewProgram(register.InitialModel(client), tea.WithAltScreen()).Run(); err != nil {
	//			fmt.Printf("could not start program: %s\n", err)
	//			os.Exit(1)
	//		}
	//	}
	//}

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
	//
	//// binary files
	//bf := &pb.BinaryFile{Name: "test binary file", File: []byte("new file"), UpdatedAt: timestamppb.New(time.Now())}
	//insertbf, err := client.BinaryFilesClient.Insert(ctx, &pb.InsertBinaryFileRequest{User: logResp.GetUser(), File: bf})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("insert bf: ", insertbf)
	//
	//bf = &pb.BinaryFile{Name: "new test binary file", File: []byte("new file"), UpdatedAt: timestamppb.New(time.Now())}
	//insertbf, err = client.BinaryFilesClient.Insert(ctx, &pb.InsertBinaryFileRequest{User: logResp.GetUser(), File: bf})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("insert bf: ", insertbf)
	//
	//getbf, err := client.BinaryFilesClient.Get(&pb.GetRequest{Name: "test binary file"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("get bf: ", getbf)
	//
	//bf = &pb.BinaryFile{Name: "new test binary file", File: []byte("new new file"), UpdatedAt: timestamppb.New(time.Now())}
	//updatebf, err := client.BinaryFilesClient.Update(ctx, &pb.UpdateBinaryFileRequest{User: logResp.GetUser(), File: bf})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("update bf: ", updatebf)
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
	//logpass := &pb.LogPass{Name: "test logpass", Login: []byte("new login"), Pass: []byte("new pass"), UpdatedAt: timestamppb.New(time.Now())}
	//insertlogpass, err := client.LogPassesClient.Insert(ctx, &pb.InsertLogPassRequest{User: logResp.GetUser(), LogPass: logpass})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("insert logpass: ", insertlogpass)
	//
	//logpass = &pb.LogPass{Name: "new test logpass", Login: []byte("new login"), Pass: []byte("new pass"), UpdatedAt: timestamppb.New(time.Now())}
	//insertlogpass, err = client.LogPassesClient.Insert(ctx, &pb.InsertLogPassRequest{User: logResp.GetUser(), LogPass: logpass})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("insert logpass: ", insertlogpass)
	//
	//getlogpass, err := client.LogPassesClient.Get(ctx, &pb.GetRequest{Name: "test logpass"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("get logpass: ", getlogpass)
	//
	//logpass = &pb.LogPass{Name: "new test logpass", Login: []byte("new new login"), Pass: []byte("new new pass"), UpdatedAt: timestamppb.New(time.Now())}
	//updatelogpass, err := client.LogPassesClient.Update(ctx, &pb.UpdateLogPassRequest{User: logResp.GetUser(), LogPass: logpass})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("update logpass: ", updatelogpass)
}
