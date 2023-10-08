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
}
