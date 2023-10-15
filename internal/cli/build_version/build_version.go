// Package build_version offers an interface to work with build version tea Model.
package build_version

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// version and build date values
var (
	Version   = ""
	BuildDate = ""
)

// used styles
var (
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	helpStyle    = blurredStyle.Copy()
)

// BuildModel is a struct for current build version model state.
type BuildModel struct {
	Finish bool   // flag if user terminated the process
	Err    error  // occurred error
	Step   string // current step
}

// NewBuildModel function creates new BuildModel instance.
func NewBuildModel() BuildModel {

	return BuildModel{
		Step: "Build",
	}
}

// Init method for tea Model interface.
func (m *BuildModel) Init() tea.Cmd {
	return nil
}

// Update method updates BuildModel state.
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

// View method displays build model view.
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
