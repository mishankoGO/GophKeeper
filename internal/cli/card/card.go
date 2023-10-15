// Package card offers interface to work with card tea Model.
package card

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/mishankoGO/GophKeeper/internal/cli/utils"
	"github.com/mishankoGO/GophKeeper/internal/client"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/cards"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
)

// fields to fill.
const (
	name = iota
	ccn
	exp
	cvv
)

// input colors.
const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

// styles.
var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
	blurredStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	helpStyle     = blurredStyle.Copy()
)

// CardModel is a tui card model instance.
type CardModel struct {
	CardInsertInputs []textinput.Model // insert page
	CardGetInputs    []textinput.Model // get page
	CardUpdateInputs []textinput.Model // update page
	CardDeleteInputs []textinput.Model // delete page
	GetResult        string            // result of get request
	InsertResult     string            // result of insert request
	UpdateResult     string            // result of update request
	DeleteResult     string            // result of delete request
	FocusedCard      int               // index of focused field
	Client           *client.Client    // client
	User             *users.User       // user object
	Finish           bool              // flag if tui is closed
	Step             string            // current step
	Err              error             // occurred error
}

// NewCardModel function creates new CardModel instance.
func NewCardModel(client *client.Client) CardModel {
	var cardInsertInputs = make([]textinput.Model, 4)
	var cardUpdateInputs = make([]textinput.Model, 4)
	var cardGetInputs = make([]textinput.Model, 1)
	var cardDeleteInputs = make([]textinput.Model, 1)

	cardGetInputs[name] = textinput.New()
	cardGetInputs[name].Placeholder = "Enter name"
	cardGetInputs[name].Focus()
	cardGetInputs[name].CharLimit = 20
	cardGetInputs[name].Width = 30
	cardGetInputs[name].Prompt = ""

	cardDeleteInputs[name] = textinput.New()
	cardDeleteInputs[name].Placeholder = "Enter name"
	cardDeleteInputs[name].Focus()
	cardDeleteInputs[name].CharLimit = 20
	cardDeleteInputs[name].Width = 30
	cardDeleteInputs[name].Prompt = ""

	cardInsertInputs[name] = textinput.New()
	cardInsertInputs[name].Placeholder = "Enter name"
	cardInsertInputs[name].Focus()
	cardInsertInputs[name].CharLimit = 20
	cardInsertInputs[name].Width = 30
	cardInsertInputs[name].Prompt = ""

	cardInsertInputs[ccn] = textinput.New()
	cardInsertInputs[ccn].Placeholder = "4505 **** **** 1234"
	cardInsertInputs[ccn].CharLimit = 20
	cardInsertInputs[ccn].Width = 30
	cardInsertInputs[ccn].Prompt = ""
	cardInsertInputs[ccn].Validate = utils.CCNValidator

	cardInsertInputs[exp] = textinput.New()
	cardInsertInputs[exp].Placeholder = "MM/YY "
	cardInsertInputs[exp].CharLimit = 5
	cardInsertInputs[exp].Width = 5
	cardInsertInputs[exp].Prompt = ""
	cardInsertInputs[exp].Validate = utils.EXPValidator

	cardInsertInputs[cvv] = textinput.New()
	cardInsertInputs[cvv].Placeholder = "XXX"
	cardInsertInputs[cvv].CharLimit = 3
	cardInsertInputs[cvv].Width = 5
	cardInsertInputs[cvv].Prompt = ""
	cardInsertInputs[cvv].Validate = utils.CVVValidator

	cardUpdateInputs[name] = textinput.New()
	cardUpdateInputs[name].Placeholder = "Enter name"
	cardUpdateInputs[name].Focus()
	cardUpdateInputs[name].CharLimit = 20
	cardUpdateInputs[name].Width = 30
	cardUpdateInputs[name].Prompt = ""

	cardUpdateInputs[ccn] = textinput.New()
	cardUpdateInputs[ccn].Placeholder = "4505 **** **** 1234"
	cardUpdateInputs[ccn].CharLimit = 20
	cardUpdateInputs[ccn].Width = 30
	cardUpdateInputs[ccn].Prompt = ""
	cardUpdateInputs[ccn].Validate = utils.CCNValidator

	cardUpdateInputs[exp] = textinput.New()
	cardUpdateInputs[exp].Placeholder = "MM/YY "
	cardUpdateInputs[exp].CharLimit = 5
	cardUpdateInputs[exp].Width = 5
	cardUpdateInputs[exp].Prompt = ""
	cardUpdateInputs[exp].Validate = utils.EXPValidator

	cardUpdateInputs[cvv] = textinput.New()
	cardUpdateInputs[cvv].Placeholder = "XXX"
	cardUpdateInputs[cvv].CharLimit = 3
	cardUpdateInputs[cvv].Width = 5
	cardUpdateInputs[cvv].Prompt = ""
	cardUpdateInputs[cvv].Validate = utils.CVVValidator

	cardModel := CardModel{
		CardInsertInputs: cardInsertInputs,
		CardUpdateInputs: cardUpdateInputs,
		CardGetInputs:    cardGetInputs,
		CardDeleteInputs: cardDeleteInputs,
		FocusedCard:      0,
		GetResult:        "",
		UpdateResult:     "",
		InsertResult:     "",
		DeleteResult:     "",
		Client:           client,
	}
	return cardModel
}

// Init method for tea model interface.
func (m *CardModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update method updates model instance according to step.
func (m *CardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.Step == "Card_INSERT" {
		return updateCardInsert(msg, m)
	} else if m.Step == "Card_GET" {
		return updateCardGet(msg, m)
	} else if m.Step == "Card_UPDATE" {
		return updateCardUpdate(msg, m)
	} else if m.Step == "Card_DELETE" {
		return updateCardDelete(msg, m)
	}
	return updateCardGet(msg, m)
}

// updateCardGet function updates get view.
func updateCardGet(msg tea.Msg, m *CardModel) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.CardGetInputs))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.FocusedCard == len(m.CardGetInputs)-1 {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				// get name
				name_ := m.CardGetInputs[name].Value()

				// convert user to proto user
				pbUser := converters.UserToPBUser(m.User)

				// get card
				getResp, err := m.Client.CardsClient.Get(ctx, &pb.GetRequest{User: pbUser, Name: name_})
				if err != nil {
					m.Err = err
					m.FocusedCard = 0
					cmds = make([]tea.Cmd, len(m.CardGetInputs))
					for i := 0; i <= len(m.CardGetInputs)-1; i++ {
						if i == m.FocusedCard {
							cmds[i] = m.CardGetInputs[i].Focus()
							m.CardGetInputs[i].Reset()
							continue
						}
						m.CardGetInputs[i].Blur()
						m.CardGetInputs[i].Reset()
					}
					m.GetResult = ""
					return m, tea.Batch(cmds...)
				}

				// update inputs
				cmds = make([]tea.Cmd, len(m.CardGetInputs))
				for i := 0; i <= len(m.CardGetInputs)-1; i++ {
					if i == m.FocusedCard {
						cmds[i] = m.CardGetInputs[i].Focus()
						m.CardGetInputs[i].Reset()
						continue
					}
					m.CardGetInputs[i].Blur()
					m.CardGetInputs[i].Reset()
				}
				m.GetResult = string(getResp.GetCard().Card)
				m.Step = "Card_GET"
				m.Err = nil
				return m, nil
			}
			m.NextInput()
		case tea.KeyCtrlZ:
			m.Step = "DataTypes"
		case tea.KeyCtrlC:
			m.Finish = true
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.PrevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.NextInput()
		}

		cmds = make([]tea.Cmd, len(m.CardGetInputs))
		for i := 0; i <= len(m.CardGetInputs)-1; i++ {
			if i == m.FocusedCard {
				cmds[i] = m.CardGetInputs[i].Focus()
				continue
			}
			m.CardGetInputs[i].Blur()
			m.CardGetInputs[i].Reset()
		}
		m.GetResult = ""
	}

	cmds = make([]tea.Cmd, len(m.CardGetInputs))
	for i := range m.CardGetInputs {
		m.CardGetInputs[i], cmds[i] = m.CardGetInputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

// updateCardDelete function updates delete view.
func updateCardDelete(msg tea.Msg, m *CardModel) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.CardDeleteInputs))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.FocusedCard == len(m.CardDeleteInputs)-1 {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				name_ := m.CardDeleteInputs[name].Value()

				pbUser := converters.UserToPBUser(m.User)

				_, err := m.Client.CardsClient.Delete(ctx, &pb.DeleteCardRequest{User: pbUser, Name: name_})
				if err != nil {
					m.Err = err
					m.FocusedCard = 0
					m.DeleteResult = ""

					cmds = make([]tea.Cmd, len(m.CardDeleteInputs))
					for i := 0; i <= len(m.CardDeleteInputs)-1; i++ {
						if i == m.FocusedCard {
							cmds[i] = m.CardDeleteInputs[i].Focus()
							m.CardDeleteInputs[i].Reset()
							continue
						}
						m.CardDeleteInputs[i].Blur()
						m.CardDeleteInputs[i].Reset()
					}
					return m, tea.Batch(cmds...)
				}
				cmds = make([]tea.Cmd, len(m.CardDeleteInputs))
				for i := 0; i <= len(m.CardDeleteInputs)-1; i++ {
					if i == m.FocusedCard {
						cmds[i] = m.CardDeleteInputs[i].Focus()
						m.CardDeleteInputs[i].Reset()
						continue
					}
					m.CardDeleteInputs[i].Blur()
					m.CardDeleteInputs[i].Reset()
				}
				m.DeleteResult = "Card deleted successfully!"
				m.Step = "Card_DELETE"
				return m, nil
			}
			m.NextInput()
		case tea.KeyCtrlZ:
			m.Step = "DataTypes"
		case tea.KeyCtrlC:
			m.Finish = true
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.PrevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.NextInput()
		}
		cmds = make([]tea.Cmd, len(m.CardDeleteInputs))
		for i := 0; i <= len(m.CardDeleteInputs)-1; i++ {
			if i == m.FocusedCard {
				cmds[i] = m.CardDeleteInputs[i].Focus()
				//m.CardDeleteInputs[i].Reset()
				continue
			}
			m.CardDeleteInputs[i].Blur()
			m.CardDeleteInputs[i].Reset()
		}
		m.DeleteResult = ""

	}

	cmds = make([]tea.Cmd, len(m.CardDeleteInputs))
	for i := range m.CardDeleteInputs {
		m.CardDeleteInputs[i], cmds[i] = m.CardDeleteInputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

// updateCardInsert function updates insert view.
func updateCardInsert(msg tea.Msg, m *CardModel) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.CardInsertInputs))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.FocusedCard == len(m.CardInsertInputs)-1 {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				name_ := m.CardInsertInputs[name].Value()
				cardNumber := m.CardInsertInputs[ccn].Value()
				expDate := m.CardInsertInputs[exp].Value()
				cvv_ := m.CardInsertInputs[cvv].Value()
				cardString := fmt.Sprintf("%s,%s,%s", cardNumber, expDate, cvv_)

				pbUser := converters.UserToPBUser(m.User)
				card := &cards.Cards{UserID: m.User.UserID, Name: name_, Card: []byte(cardString), UpdatedAt: time.Now()}

				pbCard, err := converters.CardToPBCard(card)
				if err != nil {
					m.Err = err
					m.FocusedCard = 0
					cmds = make([]tea.Cmd, len(m.CardInsertInputs))
					for i := 0; i <= len(m.CardInsertInputs)-1; i++ {
						if i == m.FocusedCard {
							cmds[i] = m.CardInsertInputs[i].Focus()
							m.CardInsertInputs[i].Reset()
							continue
						}
						m.CardInsertInputs[i].Blur()
						m.CardInsertInputs[i].Reset()
					}
					return m, tea.Batch(cmds...)
				}

				_, err = m.Client.CardsClient.Insert(ctx, &pb.InsertCardRequest{User: pbUser, Card: pbCard})
				if err != nil {
					m.Err = err
					m.FocusedCard = 0
					cmds = make([]tea.Cmd, len(m.CardInsertInputs))
					for i := 0; i <= len(m.CardInsertInputs)-1; i++ {
						if i == m.FocusedCard {
							cmds[i] = m.CardInsertInputs[i].Focus()
							m.CardInsertInputs[i].Reset()
							continue
						}
						m.CardInsertInputs[i].Blur()
						m.CardInsertInputs[i].Reset()
					}
					return m, tea.Batch(cmds...)
				}

				m.FocusedCard = 0
				m.InsertResult = "Card Inserted successfully!"
				m.Step = "Card_INSERT"

				cmds = make([]tea.Cmd, len(m.CardInsertInputs))
				for i := 0; i <= len(m.CardInsertInputs)-1; i++ {
					if i == m.FocusedCard {
						cmds[i] = m.CardInsertInputs[i].Focus()
						m.CardInsertInputs[i].Reset()
						continue
					}
					m.CardInsertInputs[i].Blur()
					m.CardInsertInputs[i].Reset()
				}

				return m, tea.Batch(cmds...)
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

		cmds = make([]tea.Cmd, len(m.CardInsertInputs))
		for i := 0; i <= len(m.CardInsertInputs)-1; i++ {
			if i == m.FocusedCard {
				cmds[i] = m.CardInsertInputs[i].Focus()
				continue
			}
			m.CardInsertInputs[i].Blur()
		}
		m.InsertResult = ""
	}

	cmds = make([]tea.Cmd, len(m.CardInsertInputs))
	for i := range m.CardInsertInputs {
		m.CardInsertInputs[i], cmds[i] = m.CardInsertInputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

// updateCardUpdate function updates update view.
func updateCardUpdate(msg tea.Msg, m *CardModel) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.CardUpdateInputs))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.FocusedCard == len(m.CardUpdateInputs)-1 {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				name_ := m.CardUpdateInputs[name].Value()
				cardNumber := m.CardUpdateInputs[ccn].Value()
				expDate := m.CardUpdateInputs[exp].Value()
				cvv_ := m.CardUpdateInputs[cvv].Value()
				cardString := fmt.Sprintf("%s,%s,%s", cardNumber, expDate, cvv_)

				pbUser := converters.UserToPBUser(m.User)
				card := &cards.Cards{UserID: m.User.UserID, Name: name_, Card: []byte(cardString), UpdatedAt: time.Now()}

				pbCard, err := converters.CardToPBCard(card)
				if err != nil {
					m.Err = err
					m.FocusedCard = 0
					cmds = make([]tea.Cmd, len(m.CardUpdateInputs))
					for i := 0; i <= len(m.CardUpdateInputs)-1; i++ {
						if i == m.FocusedCard {
							cmds[i] = m.CardUpdateInputs[i].Focus()
							m.CardUpdateInputs[i].Reset()
							continue
						}
						m.CardUpdateInputs[i].Blur()
						m.CardUpdateInputs[i].Reset()
					}
					return m, tea.Batch(cmds...)
				}

				_, err = m.Client.CardsClient.Update(ctx, &pb.UpdateCardRequest{User: pbUser, Card: pbCard})
				if err != nil {
					m.Err = err
					m.FocusedCard = 0
					cmds = make([]tea.Cmd, len(m.CardUpdateInputs))
					for i := 0; i <= len(m.CardUpdateInputs)-1; i++ {
						if i == m.FocusedCard {
							cmds[i] = m.CardUpdateInputs[i].Focus()
							m.CardUpdateInputs[i].Reset()
							continue
						}
						m.CardUpdateInputs[i].Blur()
						m.CardUpdateInputs[i].Reset()
					}
					return m, tea.Batch(cmds...)
				}

				m.UpdateResult = "Card Updated Successfully!"
				m.Step = "Card_UPDATE"
				m.FocusedCard = 0

				cmds = make([]tea.Cmd, len(m.CardUpdateInputs))
				for i := 0; i <= len(m.CardUpdateInputs)-1; i++ {
					if i == m.FocusedCard {
						cmds[i] = m.CardUpdateInputs[i].Focus()
						m.CardUpdateInputs[i].Reset()
						continue
					}
					m.CardUpdateInputs[i].Blur()
					m.CardUpdateInputs[i].Reset()
				}
				return m, tea.Batch(cmds...)
			}
			m.NextInput()
		case tea.KeyCtrlZ:
			m.Step = "DataTypes"
		case tea.KeyCtrlC:
			m.Finish = true
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.PrevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.NextInput()
		}
		cmds = make([]tea.Cmd, len(m.CardUpdateInputs))
		for i := 0; i <= len(m.CardUpdateInputs)-1; i++ {
			if i == m.FocusedCard {
				cmds[i] = m.CardUpdateInputs[i].Focus()
				continue
			}
			m.CardUpdateInputs[i].Blur()
		}
		m.UpdateResult = ""

	}

	cmds = make([]tea.Cmd, len(m.CardUpdateInputs))
	for i := range m.CardUpdateInputs {
		m.CardUpdateInputs[i], cmds[i] = m.CardUpdateInputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

// View method displays CardModel view.
func (m CardModel) View() string {
	if m.Step == "Card_INSERT" {
		return viewCardInsert(m)
	} else if m.Step == "Card_UPDATE" {
		return viewCardUpdate(m)
	} else if m.Step == "Card_DELETE" {
		return viewCardDelete(m)
	} else {
		return viewCardGet(m)
	}
}

// viewCardGet displays get view.
func viewCardGet(m CardModel) string {
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during card retrieval: %v", m.Err))
	}

	view := fmt.Sprintf(
		`Enter card name:

%s
%s

%s`,
		inputStyle.Width(30).Render("Card Name"),
		m.CardGetInputs[name].View(),
		helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
	)

	if m.GetResult != "" {
		card := strings.Split(m.GetResult, ",")

		view = fmt.Sprintf(
			`Enter card name:

%s
%s

%s %s
%s %s

%s`,

			inputStyle.Render("Card Number"),
			card[0],
			inputStyle.Width(6).Render("EXP"),
			card[1],
			inputStyle.Width(6).Render("CVV"),
			card[2],
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
		return b.String()
	}
	b.WriteString(view)
	return b.String()
}

// viewCardDelete displays delete view.
func viewCardDelete(m CardModel) string {
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during card deletion: %v", m.Err))
	}

	view := fmt.Sprintf(
		`Enter card name:

%s
%s`,
		inputStyle.Width(30).Render("Card Name"),
		m.CardDeleteInputs[name].View(),
	)

	if m.DeleteResult != "" {
		res := m.DeleteResult

		view = fmt.Sprintf(
			`Enter card name:

%s
%v

%s`,
			inputStyle.Render("Card Deleted"),
			res,
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
		return b.String()
	}
	b.WriteString(view)
	return b.String()
}

// viewCardInsert displays insert view.
func viewCardInsert(m CardModel) string {
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during card insertion: %v", m.Err))
	}

	if m.InsertResult == "" {
		view := fmt.Sprintf(
			`Enter new card info:

 %s
 %s

 %s
 %s

 %s  %s
 %s  %s

 %s

 %s
`,
			inputStyle.Width(30).Render("Card Name"),
			m.CardInsertInputs[name].View(),
			inputStyle.Width(30).Render("Card Number"),
			m.CardInsertInputs[ccn].View(),
			inputStyle.Width(6).Render("EXP"),
			inputStyle.Width(6).Render("CVV"),
			m.CardInsertInputs[exp].View(),
			m.CardInsertInputs[cvv].View(),
			continueStyle.Render("Continue ->"),
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
	} else {
		res := m.InsertResult

		view := fmt.Sprintf(
			`Enter new card info:

 %s
 %s

 %v
 
 %s

 %s
`,
			inputStyle.Width(30).Render("Card Name"),
			m.CardInsertInputs[name].View(),
			res,
			continueStyle.Render("Continue ->"),
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
	}

	return b.String()

}

// viewCardUpdate displays update view.
func viewCardUpdate(m CardModel) string {
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during card updating: %v", m.Err))
	}

	if m.UpdateResult == "" {
		view := fmt.Sprintf(
			`Enter new card info:

 %s
 %s

 %s
 %s

 %s  %s
 %s  %s

 %s

 %s
`,
			inputStyle.Width(30).Render("Card Name"),
			m.CardUpdateInputs[name].View(),
			inputStyle.Width(30).Render("Card Number"),
			m.CardUpdateInputs[ccn].View(),
			inputStyle.Width(6).Render("EXP"),
			inputStyle.Width(6).Render("CVV"),
			m.CardUpdateInputs[exp].View(),
			m.CardUpdateInputs[cvv].View(),
			continueStyle.Render("Continue ->"),
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
	} else {
		res := m.UpdateResult

		view := fmt.Sprintf(
			`Enter new card info:

 %s
 %s

 %v
 
 %s

 %s
`,
			inputStyle.Width(30).Render("Card Name"),
			m.CardUpdateInputs[name].View(),
			res,
			continueStyle.Render("Continue ->"),
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
	}

	return b.String()
}

// NextInput is a helper function to increase focus card.
func (m *CardModel) NextInput() {
	m.FocusedCard = (m.FocusedCard + 1) % len(m.CardInsertInputs)
}

// PrevInput focuses the previous input field.
func (m *CardModel) PrevInput() {
	m.FocusedCard--
	// Wrap around
	if m.FocusedCard < 0 {
		m.FocusedCard = len(m.CardInsertInputs) - 1
	}
}
