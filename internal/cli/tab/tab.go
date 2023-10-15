// Package tab offers interface to work with tabs.
// Tabs consist of Text tab, Card tab, BinaryFile tab and LogPass tab.
package tab

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// used styles.
var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	blurredStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	helpStyle         = blurredStyle.Copy()
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

// tabs and its contents.
var tabs = []string{"Card", "Text", "Binary File", "LogPass"}
var tabContent = []string{"Bank card data", "Text data", "Binary File data", "LogPass data"}

// TabModel struct for current TabModel state.
type TabModel struct {
	Tabs       []string // available tabs
	TabContent []string // tabs content
	ActiveTab  int      // current active tab
	Step       string   // current step
	Finish     bool     // flag if tui is closed
}

// NewTabModel creates new TabModel instance.
func NewTabModel() TabModel {
	return TabModel{
		Tabs:       tabs,
		TabContent: tabContent,
	}
}

// Init method for tab tea model interface.
func (m *TabModel) Init() tea.Cmd {
	return nil
}

// Update method updates tab model state.
func (m *TabModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.Finish = true
			return m, tea.Quit
		case "enter":
			m.Step = "DataTypes"
		case "right", "l", "n", "tab":
			m.ActiveTab = min(m.ActiveTab+1, len(m.Tabs)-1)
			return m, nil
		case "left", "h", "p", "shift+tab":
			m.ActiveTab = max(m.ActiveTab-1, 0)
			return m, nil
		}
	}

	return m, nil
}

// View method displays tab model view.
func (m TabModel) View() string {
	doc := strings.Builder{}

	var renderedTabs []string

	doc.WriteString("Choose your data\n\n")

	for i, t := range m.Tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.Tabs)-1, i == m.ActiveTab
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(windowStyle.Width(lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize()).Render(m.TabContent[m.ActiveTab]))
	doc.WriteString(helpStyle.Render("\nctrl+c to quit\n"))
	return docStyle.Render(doc.String())
}

// max helper function to find max integer.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min helper function to find min integer.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// tabBorderWithBottom function creates borders.
func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}
