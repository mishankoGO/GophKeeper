package index

// A simple example that shows how to retrieve a value from a Bubble Tea
// program after the Bubble Tea has exited.

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/mishankoGO/GophKeeper/internal/client"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
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
	inputStyle          = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle       = lipgloss.NewStyle().Foreground(darkGray)
	errStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("300"))
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

var tabs = []string{"Card", "Text", "Binary File", "LogPass"}
var tabContent = []string{"Bank card data", "Text data", "Binary File data", "LogPass data"}
var choices = []string{"Login", "Register"}
var dataTypes = []string{"GET", "UPDATE", "INSERT", "DELETE"}

type IndexModel struct {
	cardInputs     []textinput.Model
	focused        int
	Tabs           []string
	TabContent     []string
	activeTab      int
	DataType       string
	cursor         int
	Choice         string
	focusIndex     int
	Finish         bool
	err            error
	Quitting       bool
	Step           string
	User           *users.User
	LoginInputs    []textinput.Model
	RegisterInputs []textinput.Model
	cursorMode     cursor.Mode
	client         *client.Client
}

func InitialModel(client *client.Client) IndexModel {
	var cardInputs = make([]textinput.Model, 3)
	cardInputs[ccn] = textinput.New()
	cardInputs[ccn].Placeholder = "4505 **** **** 1234"
	cardInputs[ccn].Focus()
	cardInputs[ccn].CharLimit = 20
	cardInputs[ccn].Width = 30
	cardInputs[ccn].Prompt = ""
	cardInputs[ccn].Validate = ccnValidator

	cardInputs[exp] = textinput.New()
	cardInputs[exp].Placeholder = "MM/YY "
	cardInputs[exp].CharLimit = 5
	cardInputs[exp].Width = 5
	cardInputs[exp].Prompt = ""
	cardInputs[exp].Validate = expValidator

	cardInputs[cvv] = textinput.New()
	cardInputs[cvv].Placeholder = "XXX"
	cardInputs[cvv].CharLimit = 3
	cardInputs[cvv].Width = 5
	cardInputs[cvv].Prompt = ""
	cardInputs[cvv].Validate = cvvValidator
	m := IndexModel{
		cardInputs:     cardInputs,
		focused:        0,
		err:            nil,
		client:         client,
		Step:           "index",
		RegisterInputs: make([]textinput.Model, 3),
		LoginInputs:    make([]textinput.Model, 2),
		Tabs:           tabs,
		TabContent:     tabContent,
	}
	var t textinput.Model
	for i := range m.LoginInputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Nickname"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}

		m.LoginInputs[i] = t
	}

	for i := range m.RegisterInputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Nickname"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		case 2:
			t.Placeholder = "Confirm password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}

		m.RegisterInputs[i] = t
	}
	return m
}

/* MODEL */
func (m IndexModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m IndexModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.Step == "index" {
		return updateIndex(msg, m)
	} else if m.Step == "Login" {
		return updateLogin(msg, m)
	} else if m.Step == "Register" {
		return updateRegister(msg, m)
	} else if m.Step == "Tab" {
		return updateTab(msg, m)
	} else if m.Step == "DataTypes" {
		return updateDatatype(msg, m)
	} else if m.Step == "Card_INSERT" {
		return updateInsertCard(msg, m)
	}
	return m, nil
}

func (m IndexModel) View() string {
	if m.Step == "Login" {
		return viewLogin(m)
	} else if m.Step == "Register" {
		return viewRegister(m)
	} else if m.Step == "Tab" {
		return viewTab(m)
	} else if m.Step == "DataTypes" {
		return viewDatatype(m)
	} else if m.Step == "Card_INSERT" {
		return viewInsertCard(m)
	}
	return viewIndex(m)
}

/* REGISTER */
func updateRegister(msg tea.Msg, m IndexModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.Finish = true
			m.Quitting = true
			m.client.UsersClient.Close()
			return m, tea.Quit

		case "ctrl+z":
			m.Step = "index"
			return m, nil

		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.RegisterInputs))
			for i := range m.RegisterInputs {
				cmds[i] = m.RegisterInputs[i].Cursor.SetMode(m.cursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.RegisterInputs) {
				// login
				cred := &pb.Credential{Login: m.RegisterInputs[0].Value(), Password: m.RegisterInputs[1].Value()}
				ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
				defer cancel()

				regResp, err := m.client.UsersClient.Register(ctx, &pb.RegisterRequest{Cred: cred})
				if err != nil {
					m.err = err
					m.focusIndex = 0
					cmds := make([]tea.Cmd, len(m.RegisterInputs))
					for i := 0; i <= len(m.RegisterInputs)-1; i++ {
						if i == m.focusIndex {
							// Set focused state
							cmds[i] = m.RegisterInputs[i].Focus()
							m.RegisterInputs[i].PromptStyle = focusedStyle
							m.RegisterInputs[i].TextStyle = focusedStyle
							m.RegisterInputs[i].Reset()
							continue
						}
						// Remove focused state
						m.RegisterInputs[i].Blur()
						m.RegisterInputs[i].PromptStyle = noStyle
						m.RegisterInputs[i].TextStyle = noStyle
						m.RegisterInputs[i].Reset()
					}

					return m, tea.Batch(cmds...)
				}

				user := converters.PBUserToUser(regResp.GetUser())
				m.User = user
				m.Finish = true
				m.Step = "Login"
				//return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.RegisterInputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.RegisterInputs)
			}

			cmds := make([]tea.Cmd, len(m.RegisterInputs))
			for i := 0; i <= len(m.RegisterInputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.RegisterInputs[i].Focus()
					m.RegisterInputs[i].PromptStyle = focusedStyle
					m.RegisterInputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.RegisterInputs[i].Blur()
				m.RegisterInputs[i].PromptStyle = noStyle
				m.RegisterInputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateRegisterInputs(msg)
	return m, cmd
}

func (m *IndexModel) updateRegisterInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.RegisterInputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.RegisterInputs {
		m.RegisterInputs[i], cmds[i] = m.RegisterInputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func viewRegister(m IndexModel) string {
	var b strings.Builder

	b.WriteString("Enter your login and password\n\n")

	if m.err != nil {
		b.WriteString(fmt.Sprintf("Error occured during register: %v\n\n", m.err))
	}

	for i := range m.RegisterInputs {
		b.WriteString(m.RegisterInputs[i].View())
		if i < len(m.RegisterInputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.RegisterInputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}

/* LOGIN */
func updateLogin(msg tea.Msg, m IndexModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.Finish = true
			m.Quitting = true
			m.client.UsersClient.Close()
			return m, tea.Quit

		case "ctrl+z":
			m.Step = "index"
			return m, nil

		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.LoginInputs))
			for i := range m.LoginInputs {
				cmds[i] = m.LoginInputs[i].Cursor.SetMode(m.cursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.LoginInputs) {
				// login
				cred := &pb.Credential{Login: m.LoginInputs[0].Value(), Password: m.LoginInputs[1].Value()}
				ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
				defer cancel()

				logResp, err := m.client.UsersClient.Login(ctx, &pb.LoginRequest{Cred: cred})
				if err != nil {
					m.err = err
					m.focusIndex = 0
					cmds := make([]tea.Cmd, len(m.LoginInputs))
					for i := 0; i <= len(m.LoginInputs)-1; i++ {
						if i == m.focusIndex {
							// Set focused state
							cmds[i] = m.LoginInputs[i].Focus()
							m.LoginInputs[i].PromptStyle = focusedStyle
							m.LoginInputs[i].TextStyle = focusedStyle
							m.LoginInputs[i].Reset()
							continue
						}
						// Remove focused state
						m.LoginInputs[i].Blur()
						m.LoginInputs[i].PromptStyle = noStyle
						m.LoginInputs[i].TextStyle = noStyle
						m.LoginInputs[i].Reset()
					}

					return m, tea.Batch(cmds...)
				}

				user := converters.PBUserToUser(logResp.GetUser())
				m.User = user
				m.Finish = true
				m.Step = "Tab"
				m.cursor = 0
				return m, tea.ClearScreen
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.LoginInputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.LoginInputs)
			}

			cmds := make([]tea.Cmd, len(m.LoginInputs))
			for i := 0; i <= len(m.LoginInputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.LoginInputs[i].Focus()
					m.LoginInputs[i].PromptStyle = focusedStyle
					m.LoginInputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.LoginInputs[i].Blur()
				m.LoginInputs[i].PromptStyle = noStyle
				m.LoginInputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateLoginInputs(msg)
	return m, cmd
}

func (m *IndexModel) updateLoginInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.LoginInputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.LoginInputs {
		m.LoginInputs[i], cmds[i] = m.LoginInputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func viewLogin(m IndexModel) string {
	var b strings.Builder

	b.WriteString("Enter your login and password\n\n")

	if m.err != nil {
		b.WriteString(fmt.Sprintf("Error occured during login: %v\n\n", m.err))
	}

	for i := range m.LoginInputs {
		b.WriteString(m.LoginInputs[i].View())
		if i < len(m.LoginInputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.LoginInputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}

/* INDEX */
func updateIndex(msg tea.Msg, m IndexModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.Quitting = true
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.Choice = choices[m.cursor]
			m.Step = m.Choice

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}
	return m, nil
}

func viewIndex(m IndexModel) string {
	s := strings.Builder{}
	s.WriteString("What do you want to do?\n\n")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

/* TAB */
func updateTab(msg tea.Msg, m IndexModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.Step = "DataTypes"
			return m, tea.ClearScreen
		case "right", "l", "n", "tab":
			m.activeTab = min(m.activeTab+1, len(m.Tabs)-1)
			return m, nil
		case "left", "h", "p", "shift+tab":
			m.activeTab = max(m.activeTab-1, 0)
			return m, nil
		}
	}

	return m, nil
}

func viewTab(m IndexModel) string {
	doc := strings.Builder{}

	var renderedTabs []string

	for i, t := range m.Tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.Tabs)-1, i == m.activeTab
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
	doc.WriteString(windowStyle.Width(lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize()).Render(m.TabContent[m.activeTab]))
	return docStyle.Render(doc.String())
}

/* DATATYPE */
func updateDatatype(msg tea.Msg, m IndexModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.Quitting = true
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.Choice = dataTypes[m.cursor]
			m.Step = fmt.Sprintf("%s_%s", m.Tabs[m.activeTab], m.Choice)

		case "down", "j":
			m.cursor++
			if m.cursor >= len(dataTypes) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(dataTypes) - 1
			}
		}
	}
	return m, nil
}

func viewDatatype(m IndexModel) string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("What do you want to do with %s?\n\n", m.Tabs[m.activeTab]))

	for i := 0; i < len(dataTypes); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(dataTypes[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press ctrl+c to quit)\n")

	return s.String()
}

/* CARD insert */
func updateInsertCard(msg tea.Msg, m IndexModel) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.cardInputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.cardInputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()
		case tea.KeyCtrlZ:
			m.Step = "DataTypes"
			return m, tea.ClearScreen
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.cardInputs {
			m.cardInputs[i].Blur()
		}
		m.cardInputs[m.focused].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.cardInputs {
		m.cardInputs[i], cmds[i] = m.cardInputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func viewInsertCard(m IndexModel) string {
	return fmt.Sprintf(
		` Total: $21.50:

 %s
 %s

 %s  %s
 %s  %s

 %s
`,
		inputStyle.Width(30).Render("Card Number"),
		m.cardInputs[ccn].View(),
		inputStyle.Width(6).Render("EXP"),
		inputStyle.Width(6).Render("CVV"),
		m.cardInputs[exp].View(),
		m.cardInputs[cvv].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Validator functions to ensure valid input
func ccnValidator(s string) error {
	// Credit Card Number should a string less than 20 digits
	// It should include 16 integers and 3 spaces
	if len(s) > 16+3 {
		return fmt.Errorf("CCN is too long")
	}

	if len(s) == 0 || len(s)%5 != 0 && (s[len(s)-1] < '0' || s[len(s)-1] > '9') {
		return fmt.Errorf("CCN is invalid")
	}

	// The last digit should be a number unless it is a multiple of 4 in which
	// case it should be a space
	if len(s)%5 == 0 && s[len(s)-1] != ' ' {
		return fmt.Errorf("CCN must separate groups with spaces")
	}

	// The remaining digits should be integers
	c := strings.ReplaceAll(s, " ", "")
	_, err := strconv.ParseInt(c, 10, 64)

	return err
}

func expValidator(s string) error {
	// The 3 character should be a slash (/)
	// The rest should be numbers
	e := strings.ReplaceAll(s, "/", "")
	_, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		return fmt.Errorf("EXP is invalid")
	}

	// There should be only one slash and it should be in the 2nd index (3rd character)
	if len(s) >= 3 && (strings.Index(s, "/") != 2 || strings.LastIndex(s, "/") != 2) {
		return fmt.Errorf("EXP is invalid")
	}

	return nil
}

func cvvValidator(s string) error {
	// The CVV should be a number of 3 digits
	// Since the input will already ensure that the CVV is a string of length 3,
	// All we need to do is check that it is a number
	_, err := strconv.ParseInt(s, 10, 64)
	return err
}

// nextInput focuses the next input field
func (m *IndexModel) nextInput() {
	m.focused = (m.focused + 1) % len(m.cardInputs)
}

// prevInput focuses the previous input field
func (m *IndexModel) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.cardInputs) - 1
	}
}
