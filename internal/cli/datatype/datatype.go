package datatype

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	helpStyle    = blurredStyle.Copy()
)

var dataTypes = []string{"GET", "UPDATE", "INSERT", "DELETE"}

type DataTypeModel struct {
	Finish bool
	Choice string
	Step   string
	Cursor int
	Tab    string
}

func NewDataTypeModel() DataTypeModel {
	return DataTypeModel{
		//Step: "DataTypes"
	}
}

/* DATATYPE */

func (m *DataTypeModel) Init() tea.Cmd {
	return textinput.Blink
}
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
