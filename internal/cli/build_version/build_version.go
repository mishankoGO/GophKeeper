package build_version

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	Version   = ""
	BuildDate = ""
)

var (
	focusedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	noStyle       = lipgloss.NewStyle()
	helpStyle     = blurredStyle.Copy()
	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
	cursorStyle   = focusedStyle.Copy()
)

type BuildModel struct {
	Finish bool
	Err    error
	Step   string
}

func NewBuildModel() BuildModel {

	return BuildModel{
		Step: "Build",
	}
}

/* Build */
func (m *BuildModel) Init() tea.Cmd {
	return nil
}

func (m *BuildModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.Finish = true
			return m, tea.Quit

		case "ctrl+z":
			m.Step = "index"
			return m, nil
		}
	}
	m.Step = "index"

	return m, nil
}

func (m BuildModel) View() string {
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during build: %v\n\n", m.Err))
	}

	b.WriteString("Your Version:  ")
	b.WriteString(Version)
	b.WriteRune('\n')

	b.WriteString("Build Date:  ")
	b.WriteString(BuildDate)
	b.WriteRune('\n')

	b.WriteString(helpStyle.Render("ctrl+c to quit | ctrl+z to return"))

	return b.String()
}
