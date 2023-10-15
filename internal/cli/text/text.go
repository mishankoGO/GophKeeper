// Package text offers interface to work with text tea model.
package text

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/mishankoGO/GophKeeper/internal/client"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/texts"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
)

// fields to fill.
const (
	name = iota
)

// used colors.
const (
	hotPink = lipgloss.Color("#FF06B7")
)

// used styles.
var (
	inputStyle   = lipgloss.NewStyle().Foreground(hotPink)
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	helpStyle    = blurredStyle.Copy()
)

// TextInputs struct contains text and text name.
type TextInputs struct {
	textName []textinput.Model // field for text name
	text     textarea.Model    // input text
}

// TextModel struct for current TextModel state.
type TextModel struct {
	TextInsertInputs TextInputs        // insert inputs
	TextGetInputs    []textinput.Model // get inputs
	TextUpdateInputs TextInputs        // update inputs
	TextDeleteInputs []textinput.Model // delete inputs
	GetResult        string            // get results
	InsertResult     string            // insert results
	UpdateResult     string            // update results
	DeleteResult     string            // delete results
	FocusedText      int               // text field index
	Client           *client.Client    // client instance
	User             *users.User       // user instance
	Finish           bool              // flag if tui is closed
	Step             string            // current step
	Err              error             // occurred error
}

// NewTextModel creates new TextModel instance.
func NewTextModel(client *client.Client) TextModel {
	var textInsertInputs = TextInputs{textName: make([]textinput.Model, 1)}
	var textUpdateInputs = TextInputs{textName: make([]textinput.Model, 1)}
	var textGetInputs = make([]textinput.Model, 1)
	var textDeleteInputs = make([]textinput.Model, 1)

	textGetInputs[name] = textinput.New()
	textGetInputs[name].Placeholder = "Enter name"
	textGetInputs[name].Focus()
	textGetInputs[name].CharLimit = 20
	textGetInputs[name].Width = 30
	textGetInputs[name].Prompt = ""

	textDeleteInputs[name] = textinput.New()
	textDeleteInputs[name].Placeholder = "Enter name"
	textDeleteInputs[name].Focus()
	textDeleteInputs[name].CharLimit = 20
	textDeleteInputs[name].Width = 30
	textDeleteInputs[name].Prompt = ""

	textInsertInputs.textName[name] = textinput.New()
	textInsertInputs.textName[name].Placeholder = "Enter name"
	textInsertInputs.textName[name].Focus()
	textInsertInputs.textName[name].CharLimit = 20
	textInsertInputs.textName[name].Width = 30
	textInsertInputs.textName[name].Prompt = ""

	textInsertInputs.text = textarea.New()
	textInsertInputs.text.Placeholder = "Once upon a time"
	textInsertInputs.text.Focus()

	textUpdateInputs.textName[name] = textinput.New()
	textUpdateInputs.textName[name].Placeholder = "Enter name"
	textUpdateInputs.textName[name].Focus()
	textUpdateInputs.textName[name].CharLimit = 20
	textUpdateInputs.textName[name].Width = 30
	textUpdateInputs.textName[name].Prompt = ""

	textUpdateInputs.text = textarea.New()
	textUpdateInputs.text.Placeholder = "Once upon a time"
	textUpdateInputs.text.Focus()

	textModel := TextModel{
		TextInsertInputs: textInsertInputs,
		TextUpdateInputs: textUpdateInputs,
		TextGetInputs:    textGetInputs,
		TextDeleteInputs: textDeleteInputs,
		FocusedText:      0,
		GetResult:        "",
		UpdateResult:     "",
		InsertResult:     "",
		DeleteResult:     "",
		Client:           client,
	}
	return textModel
}

// Init method for tea model interface.
func (m *TextModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update method updated TextModel state.
func (m *TextModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.Step == "Text_INSERT" {
		return updateTextInsert(msg, m)
	} else if m.Step == "Text_GET" {
		return updateTextGet(msg, m)
	} else if m.Step == "Text_UPDATE" {
		return updateTextUpdate(msg, m)
	} else if m.Step == "Text_DELETE" {
		return updateTextDelete(msg, m)
	}

	return m, nil
}

// updateTextGet updates get page state.
func updateTextGet(msg tea.Msg, m *TextModel) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.TextGetInputs))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.FocusedText == len(m.TextGetInputs)-1 {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				name_ := m.TextGetInputs[name].Value()

				pbUser := converters.UserToPBUser(m.User)

				getResp, err := m.Client.TextsClient.Get(ctx, &pb.GetRequest{User: pbUser, Name: name_})
				if err != nil {
					m.Err = err
					m.FocusedText = 0
					cmds = make([]tea.Cmd, len(m.TextGetInputs))
					for i := 0; i <= len(m.TextGetInputs)-1; i++ {
						if i == m.FocusedText {
							cmds[i] = m.TextGetInputs[i].Focus()
							m.TextGetInputs[i].Reset()
							continue
						}
						m.TextGetInputs[i].Blur()
						m.TextGetInputs[i].Reset()
					}
					m.GetResult = ""
					return m, tea.Batch(cmds...)
				}

				cmds = make([]tea.Cmd, len(m.TextGetInputs))
				for i := 0; i <= len(m.TextGetInputs)-1; i++ {
					if i == m.FocusedText {
						cmds[i] = m.TextGetInputs[i].Focus()
						m.TextGetInputs[i].Reset()
						continue
					}
					m.TextGetInputs[i].Blur()
					m.TextGetInputs[i].Reset()
				}

				m.GetResult = string(getResp.Text.Text)
				m.Step = "Text_GET"
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

		cmds = make([]tea.Cmd, len(m.TextGetInputs))
		for i := 0; i <= len(m.TextGetInputs)-1; i++ {
			if i == m.FocusedText {
				cmds[i] = m.TextGetInputs[i].Focus()
				continue
			}
			m.TextGetInputs[i].Blur()
			m.TextGetInputs[i].Reset()
		}

		m.GetResult = ""
	}
	cmds = make([]tea.Cmd, len(m.TextGetInputs))
	for i := range m.TextGetInputs {
		m.TextGetInputs[i], cmds[i] = m.TextGetInputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

// updateTextDelete updates delete page state.
func updateTextDelete(msg tea.Msg, m *TextModel) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.TextDeleteInputs))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.FocusedText == len(m.TextDeleteInputs)-1 {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				name_ := m.TextDeleteInputs[name].Value()

				pbUser := converters.UserToPBUser(m.User)

				_, err := m.Client.TextsClient.Delete(ctx, &pb.DeleteTextRequest{
					User: pbUser,
					Name: name_,
				})
				if err != nil {
					m.Err = err
					m.FocusedText = 0
					m.DeleteResult = ""

					cmds = make([]tea.Cmd, len(m.TextDeleteInputs))
					for i := 0; i <= len(m.TextDeleteInputs)-1; i++ {
						if i == m.FocusedText {
							cmds[i] = m.TextDeleteInputs[i].Focus()
							m.TextDeleteInputs[i].Reset()
							continue
						}
						m.TextDeleteInputs[i].Blur()
						m.TextDeleteInputs[i].Reset()
					}
					return m, tea.Batch(cmds...)
				}
				cmds = make([]tea.Cmd, len(m.TextDeleteInputs))
				for i := 0; i <= len(m.TextDeleteInputs)-1; i++ {
					if i == m.FocusedText {
						cmds[i] = m.TextDeleteInputs[i].Focus()
						m.TextDeleteInputs[i].Reset()
						continue
					}
					m.TextDeleteInputs[i].Blur()
					m.TextDeleteInputs[i].Reset()
				}
				m.DeleteResult = "Text deleted successfully!"
				m.Step = "Text_DELETE"
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
		cmds = make([]tea.Cmd, len(m.TextDeleteInputs))
		for i := 0; i <= len(m.TextDeleteInputs)-1; i++ {
			if i == m.FocusedText {
				cmds[i] = m.TextDeleteInputs[i].Focus()
				continue
			}
			m.TextDeleteInputs[i].Blur()
			m.TextDeleteInputs[i].Reset()
		}
		m.DeleteResult = ""
	}

	cmds = make([]tea.Cmd, len(m.TextDeleteInputs))
	for i := range m.TextDeleteInputs {
		m.TextDeleteInputs[i], cmds[i] = m.TextDeleteInputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

// updateTextInsert updates insert page state.
func updateTextInsert(msg tea.Msg, m *TextModel) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.TextInsertInputs.textName))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlQ:
			if m.FocusedText == len(m.TextInsertInputs.textName) {
				switch msg.Type {
				case tea.KeyCtrlQ:
					ctx, cancel := context.WithCancel(context.Background())
					defer cancel()

					name_ := m.TextInsertInputs.textName[name].Value()
					text := m.TextInsertInputs.text.Value()

					pbUser := converters.UserToPBUser(m.User)
					t := &texts.Texts{UserID: pbUser.GetUserId(), Name: name_, Text: []byte(text), UpdatedAt: time.Now()}

					pbText, err := converters.TextToPBText(t)
					if err != nil {
						m.Err = err
						m.FocusedText = 0
						cmds = make([]tea.Cmd, len(m.TextInsertInputs.textName))
						for i := 0; i <= len(m.TextInsertInputs.textName)-1; i++ {
							if i == m.FocusedText {
								cmds[i] = m.TextInsertInputs.textName[i].Focus()
								m.TextInsertInputs.textName[i].Reset()
								continue
							}
							m.TextInsertInputs.textName[i].Blur()
							m.TextInsertInputs.textName[i].Reset()
						}
						m.TextInsertInputs.text.Blur()
						m.TextInsertInputs.text.Reset()
						return m, tea.Batch(cmds...)
					}

					_, err = m.Client.TextsClient.Insert(ctx, &pb.InsertTextRequest{User: pbUser, Text: pbText})
					if err != nil {
						m.Err = err
						m.FocusedText = 0
						cmds = make([]tea.Cmd, len(m.TextInsertInputs.textName))
						for i := 0; i <= len(m.TextInsertInputs.textName)-1; i++ {
							if i == m.FocusedText {
								cmds[i] = m.TextInsertInputs.textName[i].Focus()
								m.TextInsertInputs.textName[i].Reset()
								continue
							}
							m.TextInsertInputs.textName[i].Blur()
							m.TextInsertInputs.textName[i].Reset()

						}
						m.TextInsertInputs.text.Blur()
						m.TextInsertInputs.text.Reset()
						return m, tea.Batch(cmds...)
					}

					m.FocusedText = 0
					m.InsertResult = "Text inserted successfully!"
					m.Step = "Text_INSERT"

					cmds = make([]tea.Cmd, len(m.TextInsertInputs.textName))
					for i := 0; i <= len(m.TextInsertInputs.textName)-1; i++ {
						if i == m.FocusedText {
							cmds[i] = m.TextInsertInputs.text.Focus()
							m.TextInsertInputs.textName[i].Reset()
							continue
						}
						m.TextInsertInputs.textName[i].Blur()
						m.TextInsertInputs.textName[i].Reset()
					}

					return m, tea.Batch(cmds...)
				}
			}
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
		cmds = make([]tea.Cmd, len(m.TextInsertInputs.textName))
		for i := 0; i <= len(m.TextInsertInputs.textName)-1; i++ {
			if i == m.FocusedText {
				cmds[i] = m.TextInsertInputs.textName[i].Focus()
				continue
			}
			m.TextInsertInputs.textName[i].Blur()
		}
		m.InsertResult = ""

		if m.FocusedText == 1 {
			var cmd tea.Cmd
			m.TextInsertInputs.text, cmd = m.TextInsertInputs.text.Update(msg)
			return m, cmd
		}

	}

	cmds = make([]tea.Cmd, len(m.TextInsertInputs.textName))
	for i := range m.TextInsertInputs.textName {
		m.TextInsertInputs.textName[i], cmds[i] = m.TextInsertInputs.textName[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

// updateTextUpdate updates update page state.
func updateTextUpdate(msg tea.Msg, m *TextModel) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.TextUpdateInputs.textName))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlQ:
			if m.FocusedText == len(m.TextUpdateInputs.textName) {
				switch msg.Type {
				case tea.KeyCtrlQ:
					ctx, cancel := context.WithCancel(context.Background())
					defer cancel()

					name_ := m.TextUpdateInputs.textName[name].Value()
					text := m.TextUpdateInputs.text.Value()

					pbUser := converters.UserToPBUser(m.User)
					t := &texts.Texts{UserID: pbUser.GetUserId(), Name: name_, Text: []byte(text), UpdatedAt: time.Now()}

					pbText, err := converters.TextToPBText(t)
					if err != nil {
						m.Err = err
						m.FocusedText = 0
						cmds = make([]tea.Cmd, len(m.TextUpdateInputs.textName))
						for i := 0; i <= len(m.TextUpdateInputs.textName)-1; i++ {
							if i == m.FocusedText {
								cmds[i] = m.TextUpdateInputs.textName[i].Focus()
								m.TextUpdateInputs.textName[i].Reset()
								continue
							}
							m.TextUpdateInputs.textName[i].Blur()
							m.TextUpdateInputs.textName[i].Reset()
						}
						m.TextUpdateInputs.text.Blur()
						m.TextUpdateInputs.text.Reset()
						return m, tea.Batch(cmds...)
					}

					_, err = m.Client.TextsClient.Update(ctx, &pb.UpdateTextRequest{User: pbUser, Text: pbText})
					if err != nil {
						m.Err = err
						m.FocusedText = 0
						cmds = make([]tea.Cmd, len(m.TextUpdateInputs.textName))
						for i := 0; i <= len(m.TextUpdateInputs.textName)-1; i++ {
							if i == m.FocusedText {
								cmds[i] = m.TextUpdateInputs.textName[i].Focus()
								m.TextUpdateInputs.textName[i].Reset()
								continue
							}
							m.TextUpdateInputs.textName[i].Blur()
							m.TextUpdateInputs.textName[i].Reset()

						}
						m.TextUpdateInputs.text.Blur()
						m.TextUpdateInputs.text.Reset()
						return m, tea.Batch(cmds...)
					}

					m.FocusedText = 0
					m.UpdateResult = "Text updated successfully!"
					m.Step = "Text_UPDATE"

					cmds = make([]tea.Cmd, len(m.TextUpdateInputs.textName))
					for i := 0; i <= len(m.TextUpdateInputs.textName)-1; i++ {
						if i == m.FocusedText {
							cmds[i] = m.TextUpdateInputs.text.Focus()
							m.TextUpdateInputs.textName[i].Reset()
							continue
						}
						m.TextUpdateInputs.textName[i].Blur()
						m.TextUpdateInputs.textName[i].Reset()
					}

					return m, tea.Batch(cmds...)
				}
			}
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
		cmds = make([]tea.Cmd, len(m.TextUpdateInputs.textName))
		for i := 0; i <= len(m.TextUpdateInputs.textName)-1; i++ {
			if i == m.FocusedText {
				cmds[i] = m.TextUpdateInputs.textName[i].Focus()
				continue
			}
			m.TextUpdateInputs.textName[i].Blur()
		}
		m.UpdateResult = ""

		if m.FocusedText == 1 {
			var cmd tea.Cmd
			m.TextUpdateInputs.text, cmd = m.TextUpdateInputs.text.Update(msg)
			return m, cmd
		}

	}

	cmds = make([]tea.Cmd, len(m.TextUpdateInputs.textName))
	for i := range m.TextUpdateInputs.textName {
		m.TextUpdateInputs.textName[i], cmds[i] = m.TextUpdateInputs.textName[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

// View method displays TextModel view.
func (m TextModel) View() string {
	if m.Step == "Text_INSERT" {
		return viewTextInsert(m)
	} else if m.Step == "Text_UPDATE" {
		return viewTextUpdate(m)
	} else if m.Step == "Text_DELETE" {
		return viewTextDelete(m)
	} else {
		return viewTextGet(m)
	}
}

// viewTextGet displays get page.
func viewTextGet(m TextModel) string {
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during text retrieval: %v", m.Err))
	}

	view := fmt.Sprintf(
		`Enter text name:

%s
%s

%s`,
		inputStyle.Width(30).Render("Text Name"),
		m.TextGetInputs[name].View(),
		helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
	)

	if m.GetResult != "" {
		res := m.GetResult
		view = fmt.Sprintf(
			`Enter text name:

%s
%s

%s`,
			inputStyle.Render("Text"),
			res,
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
		return b.String()
	}
	b.WriteString(view)
	return b.String()
}

// viewTextInsert displays insert page.
func viewTextInsert(m TextModel) string {
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during text insertion: %v", m.Err))
	}

	if m.InsertResult == "" {
		view := fmt.Sprintf(
			`Enter new text info:

%s
%s

%s

%s
%s`,
			inputStyle.Width(30).Render("Text Name"),
			m.TextInsertInputs.textName[name].View(),
			inputStyle.Width(30).Render("Text"),
			m.TextInsertInputs.text.View(),
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return | ctrl+q to send | tab to enter text\n"),
		) + "\n"
		b.WriteString(view)
	} else {
		res := m.InsertResult

		view := fmt.Sprintf(
			`Enter new text info:
%s
%s

%s

%s
`,
			inputStyle.Width(30).Render("Text Name"),
			m.TextInsertInputs.textName[name].View(),
			res,
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return | ctrl+q to send\n"),
		) + "\n"
		b.WriteString(view)
	}
	return b.String()
}

// viewTextUpdate displays update page.
func viewTextUpdate(m TextModel) string {
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during text updating: %v", m.Err))
	}

	if m.UpdateResult == "" {
		view := fmt.Sprintf(
			`Enter new text info:

%s
%s

%s

%s
%s`,
			inputStyle.Width(30).Render("Text Name"),
			m.TextUpdateInputs.textName[name].View(),
			inputStyle.Width(30).Render("Text"),
			m.TextUpdateInputs.text.View(),
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return | ctrl+q to send | tab to enter text\n"),
		) + "\n"
		b.WriteString(view)
	} else {
		res := m.UpdateResult

		view := fmt.Sprintf(
			`Enter new text info:
%s
%s

%s

%s
`,
			inputStyle.Width(30).Render("Text Name"),
			m.TextUpdateInputs.textName[name].View(),
			res,
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return | ctrl+q to send\n"),
		) + "\n"
		b.WriteString(view)
	}
	return b.String()
}

// viewTextDelete displays delete page.
func viewTextDelete(m TextModel) string {
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during text deletion: %v", m.Err))
	}

	view := fmt.Sprintf(
		`Enter text name:

%s
%s

%s`,
		inputStyle.Width(30).Render("Text Name"),
		m.TextDeleteInputs[name].View(),
		helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
	)

	if m.DeleteResult != "" {
		res := m.DeleteResult
		view = fmt.Sprintf(
			`Enter text name:

%s

%s`,
			res,
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
		return b.String()
	}
	b.WriteString(view)
	return b.String()
}

// NextInput is a helper function to increase focus cursor.
func (m *TextModel) NextInput() {
	m.FocusedText++
	if m.FocusedText == 2 {
		m.FocusedText = 0
	}
}

// PrevInput focuses the previous input field
func (m *TextModel) PrevInput() {
	m.FocusedText--
	// Wrap around
	if m.FocusedText < 0 {
		m.FocusedText = 0
	}
}
