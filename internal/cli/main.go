package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mishankoGO/GophKeeper/internal/cli/index"
	"github.com/mishankoGO/GophKeeper/internal/cli/login"
	"github.com/mishankoGO/GophKeeper/internal/cli/register"
	"os"
)

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

func main() {

	p := tea.NewProgram(index.IndexModel{})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(index.IndexModel); ok && m.Choice != "" {
		fmt.Printf("\n---\nYou chose %s!\n", m.Choice)
		if m.Choice == "Login" {
			for {
				//m := tea.NewProgram(login.InitialModel())
				m, err := tea.NewProgram(login.InitialModel()).Run()
				if err != nil {
					fmt.Printf("could not start program: %s\n", err)
					os.Exit(1)
				}
				// Assert the final tea.Model to our local model and print the choice.
				if m, ok := m.(login.LoginModel); ok && m.Res {
					break
				}
			}

		} else if m.Choice == "Register" {
			if _, err := tea.NewProgram(register.InitialModel()).Run(); err != nil {
				fmt.Printf("could not start program: %s\n", err)
				os.Exit(1)
			}
		}
	}

}
