// Package datatype offers interface to work with datatype tea model.
package datatype

import (
	"fmt"
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

// methods to use with data.
var dataTypes = []string{"GET", "UPDATE", "INSERT", "DELETE"}

// DataTypeModel is a struct for current datatype model state.
type DataTypeModel struct {
	Finish bool   // flag if user terminated the process
	Choice string // current choice
	Step   string // current step
	Cursor int    // cursor position
	Tab    string // which tab was chosen
}

// NewDataTypeModel function create new datatype model instance.
func NewDataTypeModel() DataTypeModel {
	return DataTypeModel{}
}

// Init method for tea Model interface.
func (m *DataTypeModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update method updates current datatype model state.
func (m *DataTypeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.Finish = true
			return m, tea.Quit

		case "ctrl+z":
			m.Step = "Tab"

		case "enter":
			// Send the choice on the channel and exit.
			m.Choice = dataTypes[m.Cursor]
			m.Step = fmt.Sprintf("%s_%s", m.Tab, m.Choice)

		case "down", "j":
			m.Cursor++
			if m.Cursor >= len(dataTypes) {
				m.Cursor = 0
			}

		case "up", "k":
			m.Cursor--
			if m.Cursor < 0 {
				m.Cursor = len(dataTypes) - 1
			}
		}
	}
	return m, nil
}

// View method displays datatype model view.
func (m DataTypeModel) View() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("What do you want to do with %s?\n\n", m.Tab))

	for i := 0; i < len(dataTypes); i++ {
		if m.Cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(dataTypes[i])
		s.WriteString("\n")
	}
	s.WriteString(helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"))

	return s.String()
}
