package card

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mishankoGO/GophKeeper/internal/cli/utils"
	"github.com/mishankoGO/GophKeeper/internal/client"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const (
	name = iota
	ccn
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
	Security    security.Security
	FocusedCard int
	Client      *client.Client
	User        *users.User
	Finish      bool
	Step        string
	Err         error
}

func NewCardModel(client *client.Client, security *security.Security) CardModel {
	var cardInputs = make([]textinput.Model, 4)
	cardInputs[name] = textinput.New()
	cardInputs[name].Placeholder = "Enter name"
	cardInputs[name].Focus()
	cardInputs[name].CharLimit = 20
	cardInputs[name].Width = 30
	cardInputs[name].Prompt = ""

	cardInputs[ccn] = textinput.New()
	cardInputs[ccn].Placeholder = "4505 **** **** 1234"
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
		Client:      client,
		Security:    *security,
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
				ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
				defer cancel()

				name := m.CardInputs[name].Value()
				card_number := m.CardInputs[ccn].Value()
				expDate := m.CardInputs[exp].Value()
				cvv := m.CardInputs[cvv].Value()
				cardString := fmt.Sprintf("%s,%s,%s", card_number, expDate, cvv)

				// encrypt data
				var buf bytes.Buffer
				encoder := json.NewEncoder(&buf)
				err := encoder.Encode(cardString)
				if err != nil {
					m.Err = err
				}
				encData := m.Security.EncryptData(buf)

				card := &pb.Card{Name: name, Card: encData, UpdatedAt: timestamppb.New(time.Now())}
				pbUser := converters.UserToPBUser(m.User)

				_, err = m.Client.CardsClient.Insert(ctx, &pb.InsertCardRequest{User: pbUser, Card: card})
				if err != nil {
					m.Err = err
				}

				m.Finish = true
				m.Step = "DataTypes"
				return m, nil
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

 %s
 %s

 %s  %s
 %s  %s

 %s
`,
		inputStyle.Width(30).Render("Card Name"),
		m.CardInputs[name].View(),
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
