package binary_file

import (
	"context"
	"errors"
	"fmt"
	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mishankoGO/GophKeeper/internal/client"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/binary_files"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	name = iota
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
	blurredStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	helpStyle     = blurredStyle.Copy()
)

type BinaryFileModel struct {
	FileInsertInputs []textinput.Model // insert page
	filepicker       filepicker.Model
	selectedFile     string
	Finish           bool
	Step             string
	User             *users.User
	InsertResult     string
	Client           *client.Client
	Err              error
}

func NewBinaryFileModel(client *client.Client) BinaryFileModel {
	var fileInsertInputs = make([]textinput.Model, 1)
	fileInsertInputs[0] = textinput.New()
	fileInsertInputs[0].Placeholder = "Enter name"
	fileInsertInputs[0].Focus()
	fileInsertInputs[0].CharLimit = 20
	fileInsertInputs[0].Width = 30
	fileInsertInputs[0].Prompt = ""

	fp := filepicker.New()
	fp.AllowedTypes = []string{".mod", ".sum", ".go", ".txt", ".md"}
	d, _ := os.UserHomeDir()
	fp.CurrentDirectory = d
	fp.ShowHidden = true

	m := BinaryFileModel{
		filepicker:       fp,
		FileInsertInputs: fileInsertInputs,
		Client:           client,
	}
	return m
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m *BinaryFileModel) Init() tea.Cmd {
	cmds := []tea.Cmd{m.filepicker.Init(), textinput.Blink}
	return tea.Batch(cmds...)
}

func (m *BinaryFileModel) View() string {
	if m.Finish {
		return ""
	}
	var s strings.Builder
	s.WriteString("\n  ")
	if m.Err != nil {
		s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.Err.Error()))
	} else if m.selectedFile == "" {
		s.WriteString("Pick a file:")
	} else {
		return viewFilePicker(*m)

	}
	s.WriteString("\n\n" + m.filepicker.View() + "\n")
	return s.String()
}

func viewFilePicker(m BinaryFileModel) string {
	view := fmt.Sprintf(`
Selected file: %s

%s
%s

%s
`,
		m.filepicker.Styles.Selected.Render(m.selectedFile),
		inputStyle.Width(30).Render("Card Name:"),
		m.FileInsertInputs[0].View(),
		helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
	)
	return view
}

func updateBinaryFileInsert(m *BinaryFileModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC:
			m.Step = "Binary File_INSERT"
			return m, nil
		}
	}

	m.FileInsertInputs[0], cmd = m.FileInsertInputs[0].Update(msg)
	return m, cmd
}

func (m *BinaryFileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.Finish = true
			return m, tea.Quit
		case "enter":
			if m.selectedFile != "" {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				name_ := m.FileInsertInputs[0].Value()

				file, err := m.readFile(m.selectedFile)
				if err != nil {
					m.Err = err
					//cmds := make([]tea.Cmd, len(m.BinaryFileInsertInputs.binaryFileName))
					//for i := 0; i <= len(m.BinaryFileInsertInputs.binaryFileName)-1; i++ {
					//	if i == m.FocusedBinaryFile {
					//		cmds[i] = m.BinaryFileInsertInputs.binaryFileName[i].Focus()
					//		m.BinaryFileInsertInputs.binaryFileName[i].Reset()
					//		continue
					//	}
					//	m.BinaryFileInsertInputs.binaryFileName[i].Blur()
					//	m.BinaryFileInsertInputs.binaryFileName[i].Reset()
					//}
					//return m, tea.Batch(cmds...)
				}

				pbUser := converters.UserToPBUser(m.User)
				f := &binary_files.Files{UserID: pbUser.GetUserId(), Name: name_, File: file, UpdatedAt: time.Now()}

				pbBinaryFile, err := converters.BinaryFileToPBBinaryFile(f)
				if err != nil {
					m.Err = err
					//m.FocusedBinaryFile = 0
					//cmds := make([]tea.Cmd, len(m.BinaryFileInsertInputs.binaryFileName))
					//for i := 0; i <= len(m.BinaryFileInsertInputs.binaryFileName)-1; i++ {
					//	if i == m.FocusedBinaryFile {
					//		cmds[i] = m.BinaryFileInsertInputs.binaryFileName[i].Focus()
					//		m.BinaryFileInsertInputs.binaryFileName[i].Reset()
					//		continue
					//	}
					//	m.BinaryFileInsertInputs.binaryFileName[i].Blur()
					//	m.BinaryFileInsertInputs.binaryFileName[i].Reset()
					//}
					//return m, tea.Batch(cmds...)
				}

				_, err = m.Client.BinaryFilesClient.Insert(ctx, &pb.InsertBinaryFileRequest{User: pbUser, File: pbBinaryFile})
				if err != nil {
					m.Err = err
					//m.FocusedBinaryFile = 0
					//cmds := make([]tea.Cmd, len(m.BinaryFileInsertInputs.binaryFileName))
					//for i := 0; i <= len(m.BinaryFileInsertInputs.binaryFileName)-1; i++ {
					//	if i == m.FocusedBinaryFile {
					//		cmds[i] = m.BinaryFileInsertInputs.binaryFileName[i].Focus()
					//		m.BinaryFileInsertInputs.binaryFileName[i].Reset()
					//		continue
					//	}
					//	m.BinaryFileInsertInputs.binaryFileName[i].Blur()
					//	m.BinaryFileInsertInputs.binaryFileName[i].Reset()
					//
					//}
					//
					//return m, tea.Batch(cmds...)
				}

				m.InsertResult = "BinaryFile inserted successfully!"
				m.Step = "Binary File_INSERT"
			}
		}
	}

	if m.selectedFile != "" {
		return updateBinaryFileInsert(m, msg)
	}

	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)

	// Did the user select a file?
	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		m.selectedFile = path
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		m.Err = errors.New(path + " is not valid.")
		m.selectedFile = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return m, cmd
}

//func (m *BinaryFileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//	if m.Step == "Binary File_INSERT" {
//		return updateBinaryFileInsert(msg, m)
//	} else if m.Step == "Binary File_GET" {
//		return updateBinaryFileGet(msg, m)
//	} else if m.Step == "Binary File_UPDATE" {
//		return updateBinaryFileUpdate(msg, m)
//	} else if m.Step == "Binary File_DELETE" {
//		return updateBinaryFileDelete(msg, m)
//	}
//
//	return m, nil
//}
//
//func updateBinaryFileGet(msg tea.Msg, m *BinaryFileModel) (tea.Model, tea.Cmd) {
//	var cmds = make([]tea.Cmd, len(m.BinaryFileGetInputs))
//	switch msg := msg.(type) {
//	case tea.KeyMsg:
//		switch msg.Type {
//		case tea.KeyEnter:
//			if m.FocusedBinaryFile == len(m.BinaryFileGetInputs)-1 {
//				ctx, cancel := context.WithCancel(context.Background())
//				defer cancel()
//
//				name_ := m.BinaryFileGetInputs[name].Value()
//
//				pbUser := converters.UserToPBUser(m.User)
//
//				getResp, err := m.Client.BinaryFilesClient.Get(ctx, &pb.GetRequest{User: pbUser, Name: name_})
//				if err != nil {
//					m.Err = err
//					m.FocusedBinaryFile = 0
//					cmds = make([]tea.Cmd, len(m.BinaryFileGetInputs))
//					for i := 0; i <= len(m.BinaryFileGetInputs)-1; i++ {
//						if i == m.FocusedBinaryFile {
//							cmds[i] = m.BinaryFileGetInputs[i].Focus()
//							m.BinaryFileGetInputs[i].Reset()
//							continue
//						}
//						m.BinaryFileGetInputs[i].Blur()
//						m.BinaryFileGetInputs[i].Reset()
//					}
//					m.GetResult = ""
//					return m, tea.Batch(cmds...)
//				}
//
//				cmds = make([]tea.Cmd, len(m.BinaryFileGetInputs))
//				for i := 0; i <= len(m.BinaryFileGetInputs)-1; i++ {
//					if i == m.FocusedBinaryFile {
//						cmds[i] = m.BinaryFileGetInputs[i].Focus()
//						m.BinaryFileGetInputs[i].Reset()
//						continue
//					}
//					m.BinaryFileGetInputs[i].Blur()
//					m.BinaryFileGetInputs[i].Reset()
//				}
//
//				m.GetResult = string(getResp.GetFile().GetFile())
//				m.Step = "Binary File_GET"
//				m.Err = nil
//				return m, nil
//			}
//			m.NextInput()
//		case tea.KeyCtrlZ:
//			m.Step = "DataTypes"
//		case tea.KeyCtrlC:
//			m.Finish = true
//			return m, tea.Quit
//		case tea.KeyShiftTab, tea.KeyCtrlP:
//			m.PrevInput()
//		case tea.KeyTab, tea.KeyCtrlN:
//			m.NextInput()
//		}
//
//		cmds = make([]tea.Cmd, len(m.BinaryFileGetInputs))
//		for i := 0; i <= len(m.BinaryFileGetInputs)-1; i++ {
//			if i == m.FocusedBinaryFile {
//				cmds[i] = m.BinaryFileGetInputs[i].Focus()
//				continue
//			}
//			m.BinaryFileGetInputs[i].Blur()
//			m.BinaryFileGetInputs[i].Reset()
//		}
//
//		m.GetResult = ""
//	}
//	cmds = make([]tea.Cmd, len(m.BinaryFileGetInputs))
//	for i := range m.BinaryFileGetInputs {
//		m.BinaryFileGetInputs[i], cmds[i] = m.BinaryFileGetInputs[i].Update(msg)
//	}
//	return m, tea.Batch(cmds...)
//}
//
//func updateBinaryFileDelete(msg tea.Msg, m *BinaryFileModel) (tea.Model, tea.Cmd) {
//	var cmds = make([]tea.Cmd, len(m.BinaryFileDeleteInputs))
//	switch msg := msg.(type) {
//	case tea.KeyMsg:
//		switch msg.Type {
//		case tea.KeyEnter:
//			if m.FocusedBinaryFile == len(m.BinaryFileDeleteInputs)-1 {
//				ctx, cancel := context.WithCancel(context.Background())
//				defer cancel()
//
//				name_ := m.BinaryFileDeleteInputs[name].Value()
//
//				pbUser := converters.UserToPBUser(m.User)
//
//				_, err := m.Client.BinaryFilesClient.Delete(ctx, &pb.DeleteBinaryFileRequest{
//					User: pbUser,
//					Name: name_,
//				})
//				if err != nil {
//					m.Err = err
//					m.FocusedBinaryFile = 0
//					m.DeleteResult = ""
//
//					cmds = make([]tea.Cmd, len(m.BinaryFileDeleteInputs))
//					for i := 0; i <= len(m.BinaryFileDeleteInputs)-1; i++ {
//						if i == m.FocusedBinaryFile {
//							cmds[i] = m.BinaryFileDeleteInputs[i].Focus()
//							m.BinaryFileDeleteInputs[i].Reset()
//							continue
//						}
//						m.BinaryFileDeleteInputs[i].Blur()
//						m.BinaryFileDeleteInputs[i].Reset()
//					}
//					return m, tea.Batch(cmds...)
//				}
//				cmds = make([]tea.Cmd, len(m.BinaryFileDeleteInputs))
//				for i := 0; i <= len(m.BinaryFileDeleteInputs)-1; i++ {
//					if i == m.FocusedBinaryFile {
//						cmds[i] = m.BinaryFileDeleteInputs[i].Focus()
//						m.BinaryFileDeleteInputs[i].Reset()
//						continue
//					}
//					m.BinaryFileDeleteInputs[i].Blur()
//					m.BinaryFileDeleteInputs[i].Reset()
//				}
//				m.DeleteResult = "BinaryFile deleted successfully!"
//				m.Step = "Binary File_DELETE"
//				return m, nil
//			}
//			m.NextInput()
//		case tea.KeyCtrlZ:
//			m.Step = "DataTypes"
//		case tea.KeyCtrlC:
//			m.Finish = true
//			return m, tea.Quit
//		case tea.KeyShiftTab, tea.KeyCtrlP:
//			m.PrevInput()
//		case tea.KeyTab, tea.KeyCtrlN:
//			m.NextInput()
//		}
//		cmds = make([]tea.Cmd, len(m.BinaryFileDeleteInputs))
//		for i := 0; i <= len(m.BinaryFileDeleteInputs)-1; i++ {
//			if i == m.FocusedBinaryFile {
//				cmds[i] = m.BinaryFileDeleteInputs[i].Focus()
//				continue
//			}
//			m.BinaryFileDeleteInputs[i].Blur()
//			m.BinaryFileDeleteInputs[i].Reset()
//		}
//		m.DeleteResult = ""
//	}
//
//	cmds = make([]tea.Cmd, len(m.BinaryFileDeleteInputs))
//	for i := range m.BinaryFileDeleteInputs {
//		m.BinaryFileDeleteInputs[i], cmds[i] = m.BinaryFileDeleteInputs[i].Update(msg)
//	}
//
//	return m, tea.Batch(cmds...)
//}
//
//func updateBinaryFileInsert(msg tea.Msg, m *BinaryFileModel) (tea.Model, tea.Cmd) {
//	var cmds = make([]tea.Cmd, len(m.BinaryFileInsertInputs.binaryFileName))
//	switch msg := msg.(type) {
//	case tea.KeyMsg:
//		switch msg.Type {
//		case tea.KeyEnter:
//			m.choose = true
//			return updateBinaryFileInsertFile(msg, m)
//		case tea.KeyCtrlZ:
//			m.Step = "DataTypes"
//		case tea.KeyCtrlC, tea.KeyEsc:
//			m.Finish = true
//			return m, tea.Quit
//		case tea.KeyShiftTab, tea.KeyCtrlP:
//			m.PrevInput()
//		case tea.KeyTab, tea.KeyCtrlN:
//			m.NextInput()
//		}
//		cmds = make([]tea.Cmd, len(m.BinaryFileInsertInputs.binaryFileName))
//		for i := 0; i <= len(m.BinaryFileInsertInputs.binaryFileName)-1; i++ {
//			if i == m.FocusedBinaryFile {
//				cmds[i] = m.BinaryFileInsertInputs.binaryFileName[i].Focus()
//				continue
//			}
//			m.BinaryFileInsertInputs.binaryFileName[i].Blur()
//		}
//		m.InsertResult = ""
//	}
//
//	cmds = make([]tea.Cmd, len(m.BinaryFileInsertInputs.binaryFileName))
//	for i := range m.BinaryFileInsertInputs.binaryFileName {
//		m.BinaryFileInsertInputs.binaryFileName[i], cmds[i] = m.BinaryFileInsertInputs.binaryFileName[i].Update(msg)
//	}
//	return m, tea.Batch(cmds...)
//}
//
//func updateBinaryFileInsertFile(msg tea.Msg, m *BinaryFileModel) (tea.Model, tea.Cmd) {
//	m.BinaryFileInsertInputs.binaryFile.Init()
//	switch msg := msg.(type) {
//	case tea.KeyMsg:
//		switch msg.String() {
//		case "ctrl+c", "q":
//			m.Finish = true
//			return m, tea.Quit
//		}
//	}
//
//	var cmd tea.Cmd
//	m.BinaryFileInsertInputs.binaryFile, cmd = m.BinaryFileInsertInputs.binaryFile.Update(msg)
//
//	// Did the user select a file?
//	if didSelect, path := m.BinaryFileInsertInputs.binaryFile.DidSelectFile(msg); didSelect {
//		// Get the path of the selected file.
//		m.selectedFile = path
//		ctx, cancel := context.WithCancel(context.Background())
//		defer cancel()
//
//		name_ := m.BinaryFileInsertInputs.binaryFileName[name].Value()
//
//		file, err := m.readFile(m.selectedFile)
//		if err != nil {
//			m.Err = err
//			m.FocusedBinaryFile = 0
//			cmds := make([]tea.Cmd, len(m.BinaryFileInsertInputs.binaryFileName))
//			for i := 0; i <= len(m.BinaryFileInsertInputs.binaryFileName)-1; i++ {
//				if i == m.FocusedBinaryFile {
//					cmds[i] = m.BinaryFileInsertInputs.binaryFileName[i].Focus()
//					m.BinaryFileInsertInputs.binaryFileName[i].Reset()
//					continue
//				}
//				m.BinaryFileInsertInputs.binaryFileName[i].Blur()
//				m.BinaryFileInsertInputs.binaryFileName[i].Reset()
//			}
//			return m, tea.Batch(cmds...)
//		}
//
//		pbUser := converters.UserToPBUser(m.User)
//		f := &binary_files.Files{UserID: pbUser.GetUserId(), Name: name_, File: file, UpdatedAt: time.Now()}
//
//		pbBinaryFile, err := converters.BinaryFileToPBBinaryFile(f)
//		if err != nil {
//			m.Err = err
//			m.FocusedBinaryFile = 0
//			cmds := make([]tea.Cmd, len(m.BinaryFileInsertInputs.binaryFileName))
//			for i := 0; i <= len(m.BinaryFileInsertInputs.binaryFileName)-1; i++ {
//				if i == m.FocusedBinaryFile {
//					cmds[i] = m.BinaryFileInsertInputs.binaryFileName[i].Focus()
//					m.BinaryFileInsertInputs.binaryFileName[i].Reset()
//					continue
//				}
//				m.BinaryFileInsertInputs.binaryFileName[i].Blur()
//				m.BinaryFileInsertInputs.binaryFileName[i].Reset()
//			}
//			return m, tea.Batch(cmds...)
//		}
//
//		_, err = m.Client.BinaryFilesClient.Insert(ctx, &pb.InsertBinaryFileRequest{User: pbUser, File: pbBinaryFile})
//		if err != nil {
//			m.Err = err
//			m.FocusedBinaryFile = 0
//			cmds := make([]tea.Cmd, len(m.BinaryFileInsertInputs.binaryFileName))
//			for i := 0; i <= len(m.BinaryFileInsertInputs.binaryFileName)-1; i++ {
//				if i == m.FocusedBinaryFile {
//					cmds[i] = m.BinaryFileInsertInputs.binaryFileName[i].Focus()
//					m.BinaryFileInsertInputs.binaryFileName[i].Reset()
//					continue
//				}
//				m.BinaryFileInsertInputs.binaryFileName[i].Blur()
//				m.BinaryFileInsertInputs.binaryFileName[i].Reset()
//
//			}
//
//			return m, tea.Batch(cmds...)
//		}
//
//		m.FocusedBinaryFile = 0
//		m.InsertResult = "BinaryFile inserted successfully!"
//		m.Step = "Binary File_INSERT"
//
//		cmds := make([]tea.Cmd, len(m.BinaryFileInsertInputs.binaryFileName))
//		for i := 0; i <= len(m.BinaryFileInsertInputs.binaryFileName)-1; i++ {
//			if i == m.FocusedBinaryFile {
//				//cmds[i] = m.BinaryFileInsertInputs.binaryFile.Init()
//				m.BinaryFileInsertInputs.binaryFileName[i].Reset()
//				continue
//			}
//		}
//
//		return m, tea.Batch(cmds...)
//	}
//
//	// Did the user select a disabled file?
//	// This is only necessary to display an error to the user.
//	if didSelect, path := m.BinaryFileInsertInputs.binaryFile.DidSelectDisabledFile(msg); didSelect {
//		// Let's clear the selectedFile and display an error.
//		m.Err = errors.New(path + " is not valid.")
//		m.selectedFile = ""
//		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
//	}
//	return m, cmd
//}
//
//type clearErrorMsg struct{}
//
//func clearErrorAfter(t time.Duration) tea.Cmd {
//	return tea.Tick(t, func(_ time.Time) tea.Msg {
//		return clearErrorMsg{}
//	})
//}
//func updateBinaryFileUpdate(msg tea.Msg, m *BinaryFileModel) (tea.Model, tea.Cmd) {
//	//var cmds = make([]tea.Cmd, len(m.BinaryFileUpdateInputs.textName))
//	//switch msg := msg.(type) {
//	//case tea.KeyMsg:
//	//	switch msg.Type {
//	//	case tea.KeyEnter, tea.KeyCtrlQ:
//	//		if m.FocusedBinaryFile == len(m.BinaryFileUpdateInputs.textName) {
//	//			switch msg.Type {
//	//			case tea.KeyCtrlQ:
//	//				ctx, cancel := context.WithCancel(context.Background())
//	//				defer cancel()
//	//
//	//				name_ := m.BinaryFileUpdateInputs.textName[name].Value()
//	//				text := m.BinaryFileUpdateInputs.text.Value()
//	//
//	//				pbUser := converters.UserToPBUser(m.User)
//	//				t := &texts.BinaryFiles{UserID: pbUser.GetUserId(), Name: name_, BinaryFile: []byte(text), UpdatedAt: time.Now()}
//	//
//	//				pbBinaryFile, err := converters.BinaryFileToPBBinaryFile(t)
//	//				if err != nil {
//	//					m.Err = err
//	//					m.FocusedBinaryFile = 0
//	//					cmds = make([]tea.Cmd, len(m.BinaryFileUpdateInputs.textName))
//	//					for i := 0; i <= len(m.BinaryFileUpdateInputs.textName)-1; i++ {
//	//						if i == m.FocusedBinaryFile {
//	//							cmds[i] = m.BinaryFileUpdateInputs.textName[i].Focus()
//	//							m.BinaryFileUpdateInputs.textName[i].Reset()
//	//							continue
//	//						}
//	//						m.BinaryFileUpdateInputs.textName[i].Blur()
//	//						m.BinaryFileUpdateInputs.textName[i].Reset()
//	//					}
//	//					m.BinaryFileUpdateInputs.text.Blur()
//	//					m.BinaryFileUpdateInputs.text.Reset()
//	//					return m, tea.Batch(cmds...)
//	//				}
//	//
//	//				_, err = m.Client.BinaryFilesClient.Update(ctx, &pb.UpdateBinaryFileRequest{User: pbUser, BinaryFile: pbBinaryFile})
//	//				if err != nil {
//	//					m.Err = err
//	//					m.FocusedBinaryFile = 0
//	//					cmds = make([]tea.Cmd, len(m.BinaryFileUpdateInputs.textName))
//	//					for i := 0; i <= len(m.BinaryFileUpdateInputs.textName)-1; i++ {
//	//						if i == m.FocusedBinaryFile {
//	//							cmds[i] = m.BinaryFileUpdateInputs.textName[i].Focus()
//	//							m.BinaryFileUpdateInputs.textName[i].Reset()
//	//							continue
//	//						}
//	//						m.BinaryFileUpdateInputs.textName[i].Blur()
//	//						m.BinaryFileUpdateInputs.textName[i].Reset()
//	//
//	//					}
//	//					m.BinaryFileUpdateInputs.text.Blur()
//	//					m.BinaryFileUpdateInputs.text.Reset()
//	//					return m, tea.Batch(cmds...)
//	//				}
//	//
//	//				m.FocusedBinaryFile = 0
//	//				m.UpdateResult = "BinaryFile updated successfully!"
//	//				m.Step = "BinaryFile_UPDATE"
//	//
//	//				cmds = make([]tea.Cmd, len(m.BinaryFileUpdateInputs.textName))
//	//				for i := 0; i <= len(m.BinaryFileUpdateInputs.textName)-1; i++ {
//	//					if i == m.FocusedBinaryFile {
//	//						cmds[i] = m.BinaryFileUpdateInputs.text.Focus()
//	//						m.BinaryFileUpdateInputs.textName[i].Reset()
//	//						continue
//	//					}
//	//					m.BinaryFileUpdateInputs.textName[i].Blur()
//	//					m.BinaryFileUpdateInputs.textName[i].Reset()
//	//				}
//	//
//	//				return m, tea.Batch(cmds...)
//	//			}
//	//		}
//	//	case tea.KeyCtrlZ:
//	//		m.Step = "DataTypes"
//	//	case tea.KeyCtrlC, tea.KeyEsc:
//	//		m.Finish = true
//	//		return m, tea.Quit
//	//	case tea.KeyShiftTab, tea.KeyCtrlP:
//	//		m.PrevInput()
//	//	case tea.KeyTab, tea.KeyCtrlN:
//	//		m.NextInput()
//	//	}
//	//	cmds = make([]tea.Cmd, len(m.BinaryFileUpdateInputs.textName))
//	//	for i := 0; i <= len(m.BinaryFileUpdateInputs.textName)-1; i++ {
//	//		if i == m.FocusedBinaryFile {
//	//			cmds[i] = m.BinaryFileUpdateInputs.textName[i].Focus()
//	//			continue
//	//		}
//	//		m.BinaryFileUpdateInputs.textName[i].Blur()
//	//	}
//	//	m.UpdateResult = ""
//	//
//	//	if m.FocusedBinaryFile == 1 {
//	//		var cmd tea.Cmd
//	//		m.BinaryFileUpdateInputs.text, cmd = m.BinaryFileUpdateInputs.text.Update(msg)
//	//		return m, cmd
//	//	}
//	//
//	//}
//	//
//	//cmds = make([]tea.Cmd, len(m.BinaryFileUpdateInputs.textName))
//	//for i := range m.BinaryFileUpdateInputs.textName {
//	//	m.BinaryFileUpdateInputs.textName[i], cmds[i] = m.BinaryFileUpdateInputs.textName[i].Update(msg)
//	//}
//	//return m, tea.Batch(cmds...)
//	return m, nil
//}
//
//func (m BinaryFileModel) View() string {
//	if m.Step == "Binary File_INSERT" {
//		return viewBinaryFileInsert(m)
//	} else if m.Step == "Binary File_UPDATE" {
//		return viewBinaryFileUpdate(m)
//	} else if m.Step == "Binary File_DELETE" {
//		return viewBinaryFileDelete(m)
//	} else {
//		return viewBinaryFileGet(m)
//	}
//}
//
//func viewBinaryFileGet(m BinaryFileModel) string {
//	var b strings.Builder
//
//	if m.Err != nil {
//		b.WriteString(fmt.Sprintf("Error occured during file retrieval: %v", m.Err))
//	}
//
//	view := fmt.Sprintf(
//		`Enter file name:
//
//%s
//%s
//
//%s`,
//		inputStyle.Width(30).Render("BinaryFile Name"),
//		m.BinaryFileGetInputs[name].View(),
//		helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
//	)
//
//	if m.GetResult != "" {
//		res := m.GetResult
//		view = fmt.Sprintf(
//			`Enter file name:
//
//%s
//%s
//
//%s`,
//			inputStyle.Render("BinaryFile"),
//			res,
//			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
//		) + "\n"
//		b.WriteString(view)
//		return b.String()
//	}
//	b.WriteString(view)
//	return b.String()
//}
//
//func viewBinaryFileInsertFile(m BinaryFileModel) string {
//	//if m.Finish {
//	//	return ""
//	//}
//	var s strings.Builder
//	s.WriteString("\n  ")
//	if m.Err != nil {
//		s.WriteString(m.BinaryFileInsertInputs.binaryFile.Styles.DisabledFile.Render(m.Err.Error()))
//	} else if m.selectedFile == "" {
//		s.WriteString("Pick a file:")
//	} else {
//		s.WriteString("Selected file: " + m.BinaryFileInsertInputs.binaryFile.Styles.Selected.Render(m.selectedFile))
//	}
//	s.WriteString("\n\n" + m.BinaryFileInsertInputs.binaryFile.View() + "\n")
//	return s.String()
//}
//
//func viewBinaryFileInsert(m BinaryFileModel) string {
//	var b strings.Builder
//
//	if m.Err != nil {
//		b.WriteString(fmt.Sprintf("Error occured during file insertion: %v", m.Err))
//	}
//	return viewBinaryFileInsertFile(m)
//
//	//	if m.choose {
//	//		return viewBinaryFileInsertFile(m)
//	//	} else {
//	//		view := fmt.Sprintf(
//	//			`Enter new file info:
//	//%s
//	//%s
//	//
//	//%s
//	//`,
//	//			inputStyle.Width(30).Render("BinaryFile Name"),
//	//			m.BinaryFileInsertInputs.binaryFileName[name].View(),
//	//
//	//			helpStyle.Render("\nctrl+c to quit | ctrl+z to return | ctrl+s to send\n"),
//	//		) + "\n"
//	//		b.WriteString(view)
//	//	}
//	//	return b.String()
//}
//
//func viewBinaryFileUpdate(m BinaryFileModel) string {
//	//	var b strings.Builder
//	//
//	//	if m.Err != nil {
//	//		b.WriteString(fmt.Sprintf("Error occured during text updating: %v", m.Err))
//	//	}
//	//
//	//	if m.UpdateResult == "" {
//	//		view := fmt.Sprintf(
//	//			`Enter new text info:
//	//
//	//%s
//	//%s
//	//
//	//%s
//	//
//	//%s
//	//%s`,
//	//			inputStyle.Width(30).Render("BinaryFile Name"),
//	//			m.BinaryFileUpdateInputs.textName[name].View(),
//	//			inputStyle.Width(30).Render("BinaryFile"),
//	//			m.BinaryFileUpdateInputs.text.View(),
//	//			helpStyle.Render("\nctrl+c to quit | ctrl+z to return | ctrl+s to send | tab to enter text\n"),
//	//		) + "\n"
//	//		b.WriteString(view)
//	//	} else {
//	//		res := m.UpdateResult
//	//
//	//		view := fmt.Sprintf(
//	//			`Enter new text info:
//	//%s
//	//%s
//	//
//	//%s
//	//
//	//%s
//	//`,
//	//			inputStyle.Width(30).Render("BinaryFile Name"),
//	//			m.BinaryFileUpdateInputs.textName[name].View(),
//	//			res,
//	//			helpStyle.Render("\nctrl+c to quit | ctrl+z to return | ctrl+s to send\n"),
//	//		) + "\n"
//	//		b.WriteString(view)
//	//	}
//	//	return b.String()
//	return ""
//}
//
//func viewBinaryFileDelete(m BinaryFileModel) string {
//	var b strings.Builder
//
//	if m.Err != nil {
//		b.WriteString(fmt.Sprintf("Error occured during file deletion: %v", m.Err))
//	}
//
//	view := fmt.Sprintf(
//		`Enter file name:
//
//%s
//%s
//
//%s`,
//		inputStyle.Width(30).Render("BinaryFile Name"),
//		m.BinaryFileDeleteInputs[name].View(),
//		helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
//	)
//
//	if m.DeleteResult != "" {
//		res := m.DeleteResult
//		view = fmt.Sprintf(
//			`Enter file name:
//
//%s
//
//%s`,
//			res,
//			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
//		) + "\n"
//		b.WriteString(view)
//		return b.String()
//	}
//	b.WriteString(view)
//	return b.String()
//}
//
//func (m *BinaryFileModel) NextInput() {
//	m.FocusedBinaryFile++
//	if m.FocusedBinaryFile == 2 {
//		m.FocusedBinaryFile = 0
//	}
//}
//
//// PrevInput focuses the previous input field.
//func (m *BinaryFileModel) PrevInput() {
//	m.FocusedBinaryFile--
//	// Wrap around
//	if m.FocusedBinaryFile < 0 {
//		m.FocusedBinaryFile = 0
//	}
//}
//
// readFile reads file by path.
func (m *BinaryFileModel) readFile(path string) ([]byte, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return f, nil
}
