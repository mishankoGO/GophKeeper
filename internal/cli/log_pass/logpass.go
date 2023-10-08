package log_pass

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mishankoGO/GophKeeper/internal/client"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/log_passes"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
	"strings"
	"time"
)

// fields to fill.
const (
	name = iota
	login
	password
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

// LogPassModel is a tui logPass model instance.
type LogPassModel struct {
	LogPassInsertInputs []textinput.Model // insert page
	LogPassGetInputs    []textinput.Model // get page
	LogPassUpdateInputs []textinput.Model // update page
	LogPassDeleteInputs []textinput.Model // delete page
	GetResult           string            // result of get request
	InsertResult        string            // result of insert request
	UpdateResult        string            // result of update request
	DeleteResult        string            // result of delete request
	FocusedLogPass      int               // index of focused field
	Client              *client.Client    // client
	User                *users.User       // user object
	Finish              bool              // flag if tui is closed
	Step                string            // current step
	Err                 error             // occurred error
}

// NewLogPassModel j
func NewLogPassModel(client *client.Client) LogPassModel {
	var logPassInsertInputs = make([]textinput.Model, 3)
	var logPassUpdateInputs = make([]textinput.Model, 3)
	var logPassGetInputs = make([]textinput.Model, 1)
	var logPassDeleteInputs = make([]textinput.Model, 1)

	logPassGetInputs[name] = textinput.New()
	logPassGetInputs[name].Placeholder = "Enter name"
	logPassGetInputs[name].Focus()
	logPassGetInputs[name].CharLimit = 20
	logPassGetInputs[name].Width = 30
	logPassGetInputs[name].Prompt = ""

	logPassDeleteInputs[name] = textinput.New()
	logPassDeleteInputs[name].Placeholder = "Enter name"
	logPassDeleteInputs[name].Focus()
	logPassDeleteInputs[name].CharLimit = 20
	logPassDeleteInputs[name].Width = 30
	logPassDeleteInputs[name].Prompt = ""

	logPassInsertInputs[name] = textinput.New()
	logPassInsertInputs[name].Placeholder = "Enter name"
	logPassInsertInputs[name].Focus()
	logPassInsertInputs[name].CharLimit = 20
	logPassInsertInputs[name].Width = 30
	logPassInsertInputs[name].Prompt = ""

	logPassInsertInputs[login] = textinput.New()
	logPassInsertInputs[login].Placeholder = "your login"
	logPassInsertInputs[login].CharLimit = 20
	logPassInsertInputs[login].Width = 30
	logPassInsertInputs[login].Prompt = ""

	logPassInsertInputs[password] = textinput.New()
	logPassInsertInputs[password].Placeholder = "your password"
	logPassInsertInputs[password].CharLimit = 20
	logPassInsertInputs[password].Width = 20
	logPassInsertInputs[password].Prompt = ""

	logPassUpdateInputs[name] = textinput.New()
	logPassUpdateInputs[name].Placeholder = "Enter name"
	logPassUpdateInputs[name].Focus()
	logPassUpdateInputs[name].CharLimit = 20
	logPassUpdateInputs[name].Width = 30
	logPassUpdateInputs[name].Prompt = ""

	logPassUpdateInputs[login] = textinput.New()
	logPassUpdateInputs[login].Placeholder = "your login"
	logPassUpdateInputs[login].CharLimit = 20
	logPassUpdateInputs[login].Width = 30
	logPassUpdateInputs[login].Prompt = ""

	logPassUpdateInputs[password] = textinput.New()
	logPassUpdateInputs[password].Placeholder = "your password"
	logPassUpdateInputs[password].CharLimit = 20
	logPassUpdateInputs[password].Width = 20
	logPassUpdateInputs[password].Prompt = ""

	logPassModel := LogPassModel{
		LogPassInsertInputs: logPassInsertInputs,
		LogPassUpdateInputs: logPassUpdateInputs,
		LogPassGetInputs:    logPassGetInputs,
		LogPassDeleteInputs: logPassDeleteInputs,
		FocusedLogPass:      0,
		GetResult:           "",
		UpdateResult:        "",
		InsertResult:        "",
		DeleteResult:        "",
		Client:              client,
	}
	return logPassModel
}

func (m *LogPassModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *LogPassModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.Step == "LogPass_INSERT" {
		return updateLogPassInsert(msg, m)
	} else if m.Step == "LogPass_GET" {
		return updateLogPassGet(msg, m)
	} else if m.Step == "LogPass_UPDATE" {
		return updateLogPassUpdate(msg, m)
	} else if m.Step == "LogPass_DELETE" {
		return updateLogPassDelete(msg, m)
	}
	//m.Step = "LogPass_GET"
	return m, nil
}

func updateLogPassGet(msg tea.Msg, m *LogPassModel) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.LogPassGetInputs))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.FocusedLogPass == len(m.LogPassGetInputs)-1 {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				name_ := m.LogPassGetInputs[name].Value()

				pbUser := converters.UserToPBUser(m.User)

				getResp, err := m.Client.LogPassesClient.Get(ctx, &pb.GetRequest{User: pbUser, Name: name_})
				if err != nil {
					m.Err = err
					m.FocusedLogPass = 0
					cmds = make([]tea.Cmd, len(m.LogPassGetInputs))
					for i := 0; i <= len(m.LogPassGetInputs)-1; i++ {
						if i == m.FocusedLogPass {
							cmds[i] = m.LogPassGetInputs[i].Focus()
							m.LogPassGetInputs[i].Reset()
							continue
						}
						m.LogPassGetInputs[i].Blur()
						m.LogPassGetInputs[i].Reset()
					}
					m.GetResult = ""
					return m, tea.Batch(cmds...)
				}

				cmds = make([]tea.Cmd, len(m.LogPassGetInputs))
				for i := 0; i <= len(m.LogPassGetInputs)-1; i++ {
					if i == m.FocusedLogPass {
						cmds[i] = m.LogPassGetInputs[i].Focus()
						m.LogPassGetInputs[i].Reset()
						continue
					}
					m.LogPassGetInputs[i].Blur()
					m.LogPassGetInputs[i].Reset()
				}

				login := getResp.GetLogPass().GetLogin()
				password := getResp.GetLogPass().GetPass()
				m.GetResult = fmt.Sprintf("%s_%s", login, password)
				m.Step = "LogPass_GET"
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

		cmds = make([]tea.Cmd, len(m.LogPassGetInputs))
		for i := 0; i <= len(m.LogPassGetInputs)-1; i++ {
			if i == m.FocusedLogPass {
				cmds[i] = m.LogPassGetInputs[i].Focus()
				continue
			}
			m.LogPassGetInputs[i].Blur()
			m.LogPassGetInputs[i].Reset()
		}
		m.GetResult = ""
	}

	cmds = make([]tea.Cmd, len(m.LogPassGetInputs))
	for i := range m.LogPassGetInputs {
		m.LogPassGetInputs[i], cmds[i] = m.LogPassGetInputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func updateLogPassDelete(msg tea.Msg, m *LogPassModel) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.LogPassDeleteInputs))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.FocusedLogPass == len(m.LogPassDeleteInputs)-1 {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				name_ := m.LogPassDeleteInputs[name].Value()

				pbUser := converters.UserToPBUser(m.User)

				_, err := m.Client.LogPassesClient.Delete(ctx, &pb.DeleteLogPassRequest{User: pbUser, Name: name_})
				if err != nil {
					m.Err = err
					m.FocusedLogPass = 0
					m.DeleteResult = ""

					cmds = make([]tea.Cmd, len(m.LogPassDeleteInputs))
					for i := 0; i <= len(m.LogPassDeleteInputs)-1; i++ {
						if i == m.FocusedLogPass {
							cmds[i] = m.LogPassDeleteInputs[i].Focus()
							m.LogPassDeleteInputs[i].Reset()
							continue
						}
						m.LogPassDeleteInputs[i].Blur()
						m.LogPassDeleteInputs[i].Reset()
					}
					return m, tea.Batch(cmds...)
				}
				cmds = make([]tea.Cmd, len(m.LogPassDeleteInputs))
				for i := 0; i <= len(m.LogPassDeleteInputs)-1; i++ {
					if i == m.FocusedLogPass {
						cmds[i] = m.LogPassDeleteInputs[i].Focus()
						m.LogPassDeleteInputs[i].Reset()
						continue
					}
					m.LogPassDeleteInputs[i].Blur()
					m.LogPassDeleteInputs[i].Reset()
				}
				m.DeleteResult = "LogPass deleted successfully!"
				m.Step = "LogPass_DELETE"
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
		cmds = make([]tea.Cmd, len(m.LogPassDeleteInputs))
		for i := 0; i <= len(m.LogPassDeleteInputs)-1; i++ {
			if i == m.FocusedLogPass {
				cmds[i] = m.LogPassDeleteInputs[i].Focus()
				continue
			}
			m.LogPassDeleteInputs[i].Blur()
			m.LogPassDeleteInputs[i].Reset()
		}
		m.DeleteResult = ""

	}

	cmds = make([]tea.Cmd, len(m.LogPassDeleteInputs))
	for i := range m.LogPassDeleteInputs {
		m.LogPassDeleteInputs[i], cmds[i] = m.LogPassDeleteInputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func updateLogPassInsert(msg tea.Msg, m *LogPassModel) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.LogPassInsertInputs))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.FocusedLogPass == len(m.LogPassInsertInputs)-1 {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				name_ := m.LogPassInsertInputs[name].Value()
				login := m.LogPassInsertInputs[login].Value()
				password := m.LogPassInsertInputs[password].Value()

				pbUser := converters.UserToPBUser(m.User)
				logPass := &log_passes.LogPasses{UserID: m.User.UserID, Name: name_, Login: []byte(login), Password: []byte(password), UpdatedAt: time.Now()}

				pbLogPass, err := converters.LogPassToPBLogPass(logPass)
				if err != nil {
					m.Err = err
					m.FocusedLogPass = 0
					cmds = make([]tea.Cmd, len(m.LogPassInsertInputs))
					for i := 0; i <= len(m.LogPassInsertInputs)-1; i++ {
						if i == m.FocusedLogPass {
							cmds[i] = m.LogPassInsertInputs[i].Focus()
							m.LogPassInsertInputs[i].Reset()
							continue
						}
						m.LogPassInsertInputs[i].Blur()
						m.LogPassInsertInputs[i].Reset()
					}
					return m, tea.Batch(cmds...)
				}

				_, err = m.Client.LogPassesClient.Insert(ctx, &pb.InsertLogPassRequest{User: pbUser, LogPass: pbLogPass})
				if err != nil {
					m.Err = err
					m.FocusedLogPass = 0
					cmds = make([]tea.Cmd, len(m.LogPassInsertInputs))
					for i := 0; i <= len(m.LogPassInsertInputs)-1; i++ {
						if i == m.FocusedLogPass {
							cmds[i] = m.LogPassInsertInputs[i].Focus()
							m.LogPassInsertInputs[i].Reset()
							continue
						}
						m.LogPassInsertInputs[i].Blur()
						m.LogPassInsertInputs[i].Reset()
					}
					return m, tea.Batch(cmds...)
				}

				m.FocusedLogPass = 0
				m.InsertResult = "LogPass Inserted successfully!"
				m.Step = "LogPass_INSERT"

				cmds = make([]tea.Cmd, len(m.LogPassInsertInputs))
				for i := 0; i <= len(m.LogPassInsertInputs)-1; i++ {
					if i == m.FocusedLogPass {
						cmds[i] = m.LogPassInsertInputs[i].Focus()
						m.LogPassInsertInputs[i].Reset()
						continue
					}
					m.LogPassInsertInputs[i].Blur()
					m.LogPassInsertInputs[i].Reset()
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

		cmds = make([]tea.Cmd, len(m.LogPassInsertInputs))
		for i := 0; i <= len(m.LogPassInsertInputs)-1; i++ {
			if i == m.FocusedLogPass {
				cmds[i] = m.LogPassInsertInputs[i].Focus()
				continue
			}
			m.LogPassInsertInputs[i].Blur()
		}
		m.InsertResult = ""
	}

	cmds = make([]tea.Cmd, len(m.LogPassInsertInputs))
	for i := range m.LogPassInsertInputs {
		m.LogPassInsertInputs[i], cmds[i] = m.LogPassInsertInputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func updateLogPassUpdate(msg tea.Msg, m *LogPassModel) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.LogPassUpdateInputs))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.FocusedLogPass == len(m.LogPassUpdateInputs)-1 {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				name_ := m.LogPassUpdateInputs[name].Value()
				login := m.LogPassUpdateInputs[login].Value()
				password := m.LogPassUpdateInputs[password].Value()

				pbUser := converters.UserToPBUser(m.User)
				logPass := &log_passes.LogPasses{UserID: m.User.UserID, Name: name_, Login: []byte(login), Password: []byte(password), UpdatedAt: time.Now()}

				pbLogPass, err := converters.LogPassToPBLogPass(logPass)
				if err != nil {
					m.Err = err
					m.FocusedLogPass = 0
					cmds = make([]tea.Cmd, len(m.LogPassUpdateInputs))
					for i := 0; i <= len(m.LogPassUpdateInputs)-1; i++ {
						if i == m.FocusedLogPass {
							cmds[i] = m.LogPassUpdateInputs[i].Focus()
							m.LogPassUpdateInputs[i].Reset()
							continue
						}
						m.LogPassUpdateInputs[i].Blur()
						m.LogPassUpdateInputs[i].Reset()
					}
					return m, tea.Batch(cmds...)
				}

				_, err = m.Client.LogPassesClient.Update(ctx, &pb.UpdateLogPassRequest{User: pbUser, LogPass: pbLogPass})
				if err != nil {
					m.Err = err
					m.FocusedLogPass = 0
					cmds = make([]tea.Cmd, len(m.LogPassUpdateInputs))
					for i := 0; i <= len(m.LogPassUpdateInputs)-1; i++ {
						if i == m.FocusedLogPass {
							cmds[i] = m.LogPassUpdateInputs[i].Focus()
							m.LogPassUpdateInputs[i].Reset()
							continue
						}
						m.LogPassUpdateInputs[i].Blur()
						m.LogPassUpdateInputs[i].Reset()
					}
					return m, tea.Batch(cmds...)
				}

				m.UpdateResult = "LogPass Updated Successfully!"
				m.Step = "LogPass_UPDATE"
				m.FocusedLogPass = 0

				cmds = make([]tea.Cmd, len(m.LogPassUpdateInputs))
				for i := 0; i <= len(m.LogPassUpdateInputs)-1; i++ {
					if i == m.FocusedLogPass {
						cmds[i] = m.LogPassUpdateInputs[i].Focus()
						m.LogPassUpdateInputs[i].Reset()
						continue
					}
					m.LogPassUpdateInputs[i].Blur()
					m.LogPassUpdateInputs[i].Reset()
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
		cmds = make([]tea.Cmd, len(m.LogPassUpdateInputs))
		for i := 0; i <= len(m.LogPassUpdateInputs)-1; i++ {
			if i == m.FocusedLogPass {
				cmds[i] = m.LogPassUpdateInputs[i].Focus()
				continue
			}
			m.LogPassUpdateInputs[i].Blur()
		}
		m.UpdateResult = ""

	}

	cmds = make([]tea.Cmd, len(m.LogPassUpdateInputs))
	for i := range m.LogPassUpdateInputs {
		m.LogPassUpdateInputs[i], cmds[i] = m.LogPassUpdateInputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m LogPassModel) View() string {
	if m.Step == "LogPass_INSERT" {
		return viewLogPassInsert(m)
	} else if m.Step == "LogPass_UPDATE" {
		return viewLogPassUpdate(m)
	} else if m.Step == "LogPass_DELETE" {
		return viewLogPassDelete(m)
	} else {
		return viewLogPassGet(m)
	}
}

func viewLogPassGet(m LogPassModel) string {
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during logPass retrieval: %v", m.Err))
	}

	view := fmt.Sprintf(
		`Enter logPass name:

%s
%s

%s`,
		inputStyle.Width(30).Render("LogPass Name"),
		m.LogPassGetInputs[name].View(),
		helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
	)

	if m.GetResult != "" {
		logPass := strings.Split(m.GetResult, "_")

		view = fmt.Sprintf(
			`Enter logPass name:

%s
%s

%s
%s

%s`,

			inputStyle.Render("login"),
			logPass[0],
			inputStyle.Width(30).Render("password"),
			logPass[1],
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
		return b.String()
	}
	b.WriteString(view)
	return b.String()
}

func viewLogPassDelete(m LogPassModel) string {
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during logPass deletion: %v", m.Err))
	}

	view := fmt.Sprintf(
		`Enter logPass name:

%s
%s`,
		inputStyle.Width(30).Render("LogPass Name"),
		m.LogPassDeleteInputs[name].View(),
	)

	if m.DeleteResult != "" {
		res := m.DeleteResult

		view = fmt.Sprintf(
			`Enter logPass name:

%s
%v

%s`,
			inputStyle.Render("LogPass Deleted"),
			res,
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
		return b.String()
	}
	b.WriteString(view)
	return b.String()
}

func viewLogPassInsert(m LogPassModel) string {
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during logPass insertion: %v", m.Err))
	}

	if m.InsertResult == "" {
		view := fmt.Sprintf(
			`Enter new logPass info:

 %s
 %s

 %s
 %s

 %s
 %s

 %s

 %s
`,
			inputStyle.Width(30).Render("LogPass Name"),
			m.LogPassInsertInputs[name].View(),
			inputStyle.Width(30).Render("Login"),
			m.LogPassInsertInputs[login].View(),
			inputStyle.Width(30).Render("Password"),
			m.LogPassInsertInputs[password].View(),
			continueStyle.Render("Continue ->"),
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
	} else {
		res := m.InsertResult

		view := fmt.Sprintf(
			`Enter new logPass info:

 %s
 %s

 %v
 
 %s

 %s
`,
			inputStyle.Width(30).Render("LogPass Name"),
			m.LogPassInsertInputs[name].View(),
			res,
			continueStyle.Render("Continue ->"),
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
	}

	return b.String()

}

func viewLogPassUpdate(m LogPassModel) string {
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during logPass updating: %v", m.Err))
	}

	if m.UpdateResult == "" {
		view := fmt.Sprintf(
			`Enter new logPass info:

 %s
 %s

 %s
 %s

 %s
 %s

 %s

 %s
`,
			inputStyle.Width(30).Render("LogPass Name"),
			m.LogPassUpdateInputs[name].View(),
			inputStyle.Width(30).Render("Login"),
			m.LogPassUpdateInputs[login].View(),
			inputStyle.Width(30).Render("Password"),
			m.LogPassUpdateInputs[password].View(),
			continueStyle.Render("Continue ->"),
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
	} else {
		res := m.UpdateResult

		view := fmt.Sprintf(
			`Enter new logPass info:

 %s
 %s

 %v
 
 %s

 %s
`,
			inputStyle.Width(30).Render("LogPass Name"),
			m.LogPassUpdateInputs[name].View(),
			res,
			continueStyle.Render("Continue ->"),
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
	}

	return b.String()
}

func (m *LogPassModel) NextInput() {
	m.FocusedLogPass = (m.FocusedLogPass + 1) % len(m.LogPassInsertInputs)
}

// PrevInput focuses the previous input field
func (m *LogPassModel) PrevInput() {
	m.FocusedLogPass--
	// Wrap around
	if m.FocusedLogPass < 0 {
		m.FocusedLogPass = len(m.LogPassInsertInputs) - 1
	}
}
