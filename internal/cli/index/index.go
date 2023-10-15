// Package index offers interface to work with index tea model.
package index

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// used styles.
var (
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	helpStyle    = blurredStyle.Copy()
)

// pages to choose.
var choices = []string{"Login", "Register", "Build"}

// IndexModel is a struct for current index model state.
type IndexModel struct {
	Choice string // current choice
	Step   string // current step
	Finish bool   // flag if user terminated the process
	Cursor int    // cursor position
}

// NewIndexModel function creates new IndexModel instance.
func NewIndexModel() IndexModel {
	return IndexModel{}
}

// Update method updates index model state.
func (m *IndexModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.Finish = true
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.Choice = choices[m.Cursor]
			m.Step = m.Choice

		case "down", "j":
			m.Cursor++
			if m.Cursor >= len(choices) {
				m.Cursor = 0
			}

		case "up", "k":
			m.Cursor--
			if m.Cursor < 0 {
				m.Cursor = len(choices) - 1
			}
		}
	}
	return m, nil
}

// View method displays index model view.
func (m IndexModel) View() string {
	s := strings.Builder{}
	s.WriteString("What do you want to do?\n\n")

	for i := 0; i < len(choices); i++ {
		if m.Cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString(helpStyle.Render("\nctrl+c to quit\n"))

	return s.String()
}

// Init method for tea Model interface.
func (m *IndexModel) Init() tea.Cmd {
	return textinput.Blink
}
