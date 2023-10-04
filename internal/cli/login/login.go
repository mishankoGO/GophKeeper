package login

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mishankoGO/GophKeeper/internal/client"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
	"strings"
	"time"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	focusedButton       = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton       = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	cursorStyle         = focusedStyle.Copy()
)

type LoginModel struct {
	LoginCursorMode cursor.Mode
	LoginFocusIndex int
	LoginInputs     []textinput.Model
	Client          *client.Client
	User            *users.User
	Finish          bool
	Err             error
	Step            string
}

type RegisterModel struct {
	RegisterCursorMode cursor.Mode
	RegisterFocusIndex int
	RegisterInputs     []textinput.Model
	Client             *client.Client
	User               *users.User
	Finish             bool
	Err                error
	Step               string
}

func NewLoginModel(client *client.Client) LoginModel {
	var t textinput.Model
	loginInputs := make([]textinput.Model, 2)
	for i := range loginInputs {
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

		loginInputs[i] = t
	}

	return LoginModel{
		LoginInputs: loginInputs,
		Client:      client,
		Step:        "Login",
	}
}

func NewRegisterModel(client *client.Client) RegisterModel {
	var t textinput.Model

	registerInputs := make([]textinput.Model, 3)
	for i := range registerInputs {
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

		registerInputs[i] = t
	}

	return RegisterModel{
		RegisterInputs: registerInputs,
		Client:         client,
		Step:           "Register",
	}
}

/* REGISTER */
func (m *RegisterModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *RegisterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.Finish = true
			m.Client.UsersClient.Close()
			return m, tea.Quit

		case "ctrl+z":
			m.Step = "index"
			return m, nil

		// Change cursor mode
		case "ctrl+r":
			m.RegisterCursorMode++
			if m.RegisterCursorMode > cursor.CursorHide {
				m.RegisterCursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.RegisterInputs))
			for i := range m.RegisterInputs {
				cmds[i] = m.RegisterInputs[i].Cursor.SetMode(m.RegisterCursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.RegisterFocusIndex == len(m.RegisterInputs) {
				// login
				cred := &pb.Credential{Login: m.RegisterInputs[0].Value(), Password: m.RegisterInputs[1].Value()}
				ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
				defer cancel()

				regResp, err := m.Client.UsersClient.Register(ctx, &pb.RegisterRequest{Cred: cred})
				if err != nil {
					m.Err = err
					m.RegisterFocusIndex = 0
					cmds := make([]tea.Cmd, len(m.RegisterInputs))
					for i := 0; i <= len(m.RegisterInputs)-1; i++ {
						if i == m.RegisterFocusIndex {
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
				return m, nil
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.RegisterFocusIndex--
			} else {
				m.RegisterFocusIndex++
			}

			if m.RegisterFocusIndex > len(m.RegisterInputs) {
				m.RegisterFocusIndex = 0
			} else if m.RegisterFocusIndex < 0 {
				m.RegisterFocusIndex = len(m.RegisterInputs)
			}

			cmds := make([]tea.Cmd, len(m.RegisterInputs))
			for i := 0; i <= len(m.RegisterInputs)-1; i++ {
				if i == m.RegisterFocusIndex {
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
	cmd := updateRegisterInputs(msg, *m)
	return m, cmd
}

func updateRegisterInputs(msg tea.Msg, m RegisterModel) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.RegisterInputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.RegisterInputs {
		m.RegisterInputs[i], cmds[i] = m.RegisterInputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m RegisterModel) View() string {
	var b strings.Builder

	b.WriteString("Enter your login and password\n\n")

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during register: %v\n\n", m.Err))
	}

	for i := range m.RegisterInputs {
		b.WriteString(m.RegisterInputs[i].View())
		if i < len(m.RegisterInputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.RegisterFocusIndex == len(m.RegisterInputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.RegisterCursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}

/* LOGIN */
func (m *LoginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.Finish = true
			m.Client.UsersClient.Close()
			return m, tea.Quit

		case "ctrl+z":
			m.Step = "index"
			return m, nil

		// Change cursor mode
		case "ctrl+r":
			m.LoginCursorMode++
			if m.LoginCursorMode > cursor.CursorHide {
				m.LoginCursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.LoginInputs))
			for i := range m.LoginInputs {
				cmds[i] = m.LoginInputs[i].Cursor.SetMode(m.LoginCursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.LoginFocusIndex == len(m.LoginInputs) {
				// login
				cred := &pb.Credential{Login: m.LoginInputs[0].Value(), Password: m.LoginInputs[1].Value()}
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				logResp, err := m.Client.UsersClient.Login(ctx, &pb.LoginRequest{Cred: cred})
				if err != nil {
					m.Err = err
					m.LoginFocusIndex = 0
					cmds := make([]tea.Cmd, len(m.LoginInputs))
					for i := 0; i <= len(m.LoginInputs)-1; i++ {
						if i == m.LoginFocusIndex {
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
				return m, tea.ClearScreen
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.LoginFocusIndex--
			} else {
				m.LoginFocusIndex++
			}

			if m.LoginFocusIndex > len(m.LoginInputs) {
				m.LoginFocusIndex = 0
			} else if m.LoginFocusIndex < 0 {
				m.LoginFocusIndex = len(m.LoginInputs)
			}

			cmds := make([]tea.Cmd, len(m.LoginInputs))
			for i := 0; i <= len(m.LoginInputs)-1; i++ {
				if i == m.LoginFocusIndex {
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
	cmd := updateLoginInputs(msg, m)
	return m, cmd
}

func updateLoginInputs(msg tea.Msg, m *LoginModel) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.LoginInputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.LoginInputs {
		m.LoginInputs[i], cmds[i] = m.LoginInputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m LoginModel) View() string {
	var b strings.Builder
	b.WriteString("Enter your login and password\n\n")

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during login: %v\n\n", m.Err))
	}

	for i := range m.LoginInputs {
		b.WriteString(m.LoginInputs[i].View())
		if i < len(m.LoginInputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.LoginFocusIndex == len(m.LoginInputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.LoginCursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}
