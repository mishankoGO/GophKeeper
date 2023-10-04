package card

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mishankoGO/GophKeeper/internal/cli/utils"
)

const (
	ccn = iota
	exp
	cvv
)

type (
	errMsg error
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

type CardModel struct {
	CardInputs  []textinput.Model
	FocusedCard int
	Finish      bool
	Step        string
	Err         error
}

func NewCardModel() CardModel {
	var cardInputs = make([]textinput.Model, 3)
	cardInputs[ccn] = textinput.New()
	cardInputs[ccn].Placeholder = "4505 **** **** 1234"
	cardInputs[ccn].Focus()
	cardInputs[ccn].CharLimit = 20
	cardInputs[ccn].Width = 30
	cardInputs[ccn].Prompt = ""
	cardInputs[ccn].Validate = utils.CCNValidator

	cardInputs[exp] = textinput.New()
	cardInputs[exp].Placeholder = "MM/YY "
	cardInputs[exp].CharLimit = 5
	cardInputs[exp].Width = 5
	cardInputs[exp].Prompt = ""
	cardInputs[exp].Validate = utils.EXPValidator

	cardInputs[cvv] = textinput.New()
	cardInputs[cvv].Placeholder = "XXX"
	cardInputs[cvv].CharLimit = 3
	cardInputs[cvv].Width = 5
	cardInputs[cvv].Prompt = ""
	cardInputs[cvv].Validate = utils.CVVValidator

	cardModel := CardModel{
		CardInputs:  cardInputs,
		FocusedCard: 0,
		Step:        "Card_INSERT",
	}
	return cardModel
}

/* CARD insert */

func (m *CardModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *CardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.CardInputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.FocusedCard == len(m.CardInputs)-1 {
				return m, tea.Quit
			}
			m.NextInput()
		case tea.KeyCtrlZ:
			m.Step = "DataTypes"
		case tea.KeyCtrlC, tea.KeyEsc:
			m.Finish = true
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.PrevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.NextInput()
		}
		for i := range m.CardInputs {
			m.CardInputs[i].Blur()
		}
		m.CardInputs[m.FocusedCard].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.Err = msg
		return m, nil
	}

	for i := range m.CardInputs {
		m.CardInputs[i], cmds[i] = m.CardInputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m CardModel) View() string {
	return fmt.Sprintf(
		`Enter new card info:

 %s
 %s

 %s  %s
 %s  %s

 %s
`,
		inputStyle.Width(30).Render("Card Number"),
		m.CardInputs[ccn].View(),
		inputStyle.Width(6).Render("EXP"),
		inputStyle.Width(6).Render("CVV"),
		m.CardInputs[exp].View(),
		m.CardInputs[cvv].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

func (m *CardModel) NextInput() {
	m.FocusedCard = (m.FocusedCard + 1) % len(m.CardInputs)
}

// prevInput focuses the previous input field
func (m *CardModel) PrevInput() {
	m.FocusedCard--
	// Wrap around
	if m.FocusedCard < 0 {
		m.FocusedCard = len(m.CardInputs) - 1
	}
}
