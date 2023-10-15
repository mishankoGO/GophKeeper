// Package binary_file offers an interface to work with binary files tea Model.
// BUG(Михаил Михайлов): странно работает выбор файлов. Нужно нажат стрелки влево вправо, а потом немного подвигать окно терминала.
package binary_file

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/mishankoGO/GophKeeper/internal/client"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/binary_files"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
)

// used colors
const (
	hotPink = lipgloss.Color("#FF06B7")
)

// used styles
var (
	inputStyle   = lipgloss.NewStyle().Foreground(hotPink)
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	helpStyle    = blurredStyle.Copy()
)

// BinaryFileModel is a struct for current binary model state.
type BinaryFileModel struct {
	FileInsertInputs []textinput.Model // insert page
	FileGetInputs    []textinput.Model // get page
	FileDeleteInputs []textinput.Model // delete page
	FileUpdateInputs []textinput.Model // update page
	filepicker       filepicker.Model  // file picker instance
	selectedFile     string            // currently selected file
	Finish           bool              // flag if user terminated the process
	Step             string            // current step
	User             *users.User       // user instance
	InsertResult     string            // result of insertion
	GetResult        string            // result of retrieval
	DeleteResult     string            // result of deletion
	UpdateResult     string            // result of update
	Client           *client.Client    // client instance
	Err              error             // occurred error
}

// NewBinaryFileModel function creates new BinaryFileModel instance.
func NewBinaryFileModel(client *client.Client) BinaryFileModel {
	var fileInsertInputs = make([]textinput.Model, 1)
	fileInsertInputs[0] = textinput.New()
	fileInsertInputs[0].Placeholder = "Enter name"
	fileInsertInputs[0].Focus()
	fileInsertInputs[0].CharLimit = 20
	fileInsertInputs[0].Width = 30
	fileInsertInputs[0].Prompt = ""

	var fileUpdateInputs = make([]textinput.Model, 1)
	fileUpdateInputs[0] = textinput.New()
	fileUpdateInputs[0].Placeholder = "Enter name"
	fileUpdateInputs[0].Focus()
	fileUpdateInputs[0].CharLimit = 20
	fileUpdateInputs[0].Width = 30
	fileUpdateInputs[0].Prompt = ""

	var fileGetInputs = make([]textinput.Model, 1)
	fileGetInputs[0] = textinput.New()
	fileGetInputs[0].Placeholder = "Enter name"
	fileGetInputs[0].Focus()
	fileGetInputs[0].CharLimit = 20
	fileGetInputs[0].Width = 30
	fileGetInputs[0].Prompt = ""

	var fileDeleteInputs = make([]textinput.Model, 1)
	fileDeleteInputs[0] = textinput.New()
	fileDeleteInputs[0].Placeholder = "Enter name"
	fileDeleteInputs[0].Focus()
	fileDeleteInputs[0].CharLimit = 20
	fileDeleteInputs[0].Width = 30
	fileDeleteInputs[0].Prompt = ""

	fp := filepicker.New()
	fp.AllowedTypes = []string{".mod", ".sum", ".go", ".txt", ".md"}
	d, _ := os.UserHomeDir()
	fp.CurrentDirectory = d
	fp.ShowHidden = false

	m := BinaryFileModel{
		filepicker:       fp,
		FileInsertInputs: fileInsertInputs,
		FileGetInputs:    fileGetInputs,
		FileDeleteInputs: fileDeleteInputs,
		FileUpdateInputs: fileUpdateInputs,
		Client:           client,
		GetResult:        "",
		DeleteResult:     "",
		Finish:           false,
	}
	return m
}

// clearErrorMsg for filepicker errors.
type clearErrorMsg struct{}

// clearErrorAfter function clears error after t duration.
func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

// Init method for tea Model interface.
func (m *BinaryFileModel) Init() tea.Cmd {
	cmds := []tea.Cmd{m.filepicker.Init(), textinput.Blink}
	return tea.Batch(cmds...)
}

// View method displays view according to step.
func (m *BinaryFileModel) View() string {
	if m.Step == "Binary File_INSERT" {
		return m.ViewInsert()
	} else if m.Step == "Binary File_DELETE" {
		return m.ViewDelete()
	} else if m.Step == "Binary File_GET" {
		return m.ViewGet()
	} else if m.Step == "Binary File_UPDATE" {
		return m.ViewUpdate()
	}
	return m.ViewGet()
}

// ViewInsert method for insert step.
func (m *BinaryFileModel) ViewInsert() string {
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
		return viewFilePicker(m)
	}
	s.WriteString("\n\n" + m.filepicker.View() + "\n")
	return s.String()
}

// ViewGet method for get step.
func (m *BinaryFileModel) ViewGet() string {
	if m.Finish {
		return ""
	}
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during file retrieval: %v", m.Err))
	}

	view := fmt.Sprintf(
		`Enter file name:

%s
%s

%s`,
		inputStyle.Width(30).Render("File Name"),
		m.FileGetInputs[0].View(),
		helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
	)

	// if GetResult is empty then the view is different
	if m.GetResult != "" {
		file := m.GetResult

		view = fmt.Sprintf(
			`Enter file name:

%s
%s

%s`,

			inputStyle.Render("Retrieved file:"),
			file,
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
		return b.String()
	}
	b.WriteString(view)
	return b.String()
}

// ViewDelete method for delete step.
func (m *BinaryFileModel) ViewDelete() string {
	if m.Finish {
		return ""
	}
	var b strings.Builder

	if m.Err != nil {
		b.WriteString(fmt.Sprintf("Error occured during file removal: %v", m.Err))
	}

	view := fmt.Sprintf(
		`Enter file name:

%s
%s

%s`,
		inputStyle.Width(30).Render("File Name"),
		m.FileDeleteInputs[0].View(),
		helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
	)

	if m.DeleteResult != "" {
		file := m.DeleteResult

		view = fmt.Sprintf(
			`Enter file name:

%s
%s

%s`,

			inputStyle.Render("Retrieved file:"),
			file,
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
		) + "\n"
		b.WriteString(view)
		return b.String()
	}
	b.WriteString(view)
	return b.String()
}

// ViewUpdate method for update step.
func (m *BinaryFileModel) ViewUpdate() string {
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
		return viewFilePickerUpdate(*m)
	}
	s.WriteString("\n\n" + m.filepicker.View() + "\n")
	return s.String()
}

// viewFilePicker function shows updated filepicker after insertion.
func viewFilePicker(m *BinaryFileModel) string {
	view := fmt.Sprintf(`
Selected file: %s

%s
%s

%s
`,
		m.filepicker.Styles.Selected.Render(m.selectedFile),
		inputStyle.Width(30).Render("File Name:"),
		m.FileInsertInputs[0].View(),
		helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
	)

	if m.InsertResult != "" {
		view = fmt.Sprintf(`
Selected file: %s

%s

%s`,
			m.filepicker.Styles.Selected.Render(m.selectedFile),
			m.InsertResult,
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"))
		return view
	}

	return view
}

// viewFilePickerUpdate function shows updated filepicker after update.
func viewFilePickerUpdate(m BinaryFileModel) string {
	view := fmt.Sprintf(`
Selected file: %s

%s
%s

%s
`,
		m.filepicker.Styles.Selected.Render(m.selectedFile),
		inputStyle.Width(30).Render("File Name:"),
		m.FileUpdateInputs[0].View(),
		helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"),
	)

	if m.UpdateResult != "" {
		view = fmt.Sprintf(`
Selected file: %s

%s

%s`,
			m.filepicker.Styles.Selected.Render(m.selectedFile),
			m.UpdateResult,
			helpStyle.Render("\nctrl+c to quit | ctrl+z to return\n"))
		return view
	}

	return view
}

// Update method for tea Model interface.
func (m *BinaryFileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.Step == "Binary File_INSERT" {
		return m.UpdateInsert(msg)
	} else if m.Step == "Binary File_GET" {
		return m.UpdateGet(msg)
	} else if m.Step == "Binary File_DELETE" {
		return m.UpdateDelete(msg)
	} else if m.Step == "Binary File_UPDATE" {
		return m.UpdateUpdate(msg)
	}
	return m.UpdateGet(msg)
}

// updateBinaryFileInsert function updates insert inputs.
func updateBinaryFileInsert(m *BinaryFileModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.Step = "Binary File_INSERT"
			return m, nil
		}
	}

	m.FileInsertInputs[0], cmd = m.FileInsertInputs[0].Update(msg)
	return m, cmd
}

// UpdateInsert method to update insert view.
func (m *BinaryFileModel) UpdateInsert(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.Finish = true
			m.selectedFile = ""
			return m, tea.Quit
		case "ctrl+z":
			m.Step = "DataTypes"
			m.selectedFile = ""
			m.InsertResult = ""
			m.FileInsertInputs[0].Reset()
			return m, nil
		case "enter":
			if m.selectedFile != "" {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				// get name and extension
				name_ := m.FileInsertInputs[0].Value()
				splittedName := strings.Split(m.selectedFile, ".")
				extension := "." + splittedName[len(splittedName)-1]

				// read file
				file, err := m.readFile(m.selectedFile)
				if err != nil {
					m.Err = err
					m.selectedFile = ""
					m.InsertResult = ""
					m.FileInsertInputs[0].Blur()
					m.FileInsertInputs[0].Reset()
					return m, nil
				}

				// create proto file
				pbUser := converters.UserToPBUser(m.User)
				f := &binary_files.Files{UserID: pbUser.GetUserId(), Name: name_, File: file, Extension: []byte(extension), UpdatedAt: time.Now()}
				pbBinaryFile, err := converters.BinaryFileToPBBinaryFile(f)
				if err != nil {
					m.Err = err
					m.selectedFile = ""
					m.InsertResult = ""
					m.FileInsertInputs[0].Blur()
					m.FileInsertInputs[0].Reset()
					return m, nil
				}

				// insert file
				_, err = m.Client.BinaryFilesClient.Insert(ctx, &pb.InsertBinaryFileRequest{User: pbUser, File: pbBinaryFile})
				if err != nil {
					m.Err = err
					m.selectedFile = ""
					m.InsertResult = ""
					//m.FileInsertInputs[0].Blur()
					m.FileInsertInputs[0].Reset()
					return m, nil
				}
				// update state
				m.InsertResult = "BinaryFile inserted successfully!"
				m.Step = "Binary File_INSERT"
				m.FileInsertInputs[0].Reset()
				m.FileInsertInputs[0].Focus()
				m.filepicker.FileSelected = ""
				return m, nil
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

// updateBinaryFileUpdate function updates update inputs.
func updateBinaryFileUpdate(m *BinaryFileModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.Step = "Binary File_UPDATE"
			return m, nil
		}
	}

	m.FileUpdateInputs[0], cmd = m.FileUpdateInputs[0].Update(msg)
	return m, cmd
}

// UpdateUpdate method to update update view.
func (m *BinaryFileModel) UpdateUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.Finish = true
			m.selectedFile = ""
			return m, tea.Quit
		case "ctrl+z":
			m.Step = "DataTypes"
			m.selectedFile = ""
			m.FileUpdateInputs[0].Reset()
		case "enter":
			if m.selectedFile != "" {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				// get name and extension
				name_ := m.FileUpdateInputs[0].Value()
				splittedName := strings.Split(m.selectedFile, ".")
				extension := "." + splittedName[len(splittedName)-1]

				// read file
				file, err := m.readFile(m.selectedFile)
				if err != nil {
					m.Err = err
					m.selectedFile = ""
					m.FileUpdateInputs[0].Blur()
					m.FileUpdateInputs[0].Reset()
					return m, nil
				}

				// create proto file
				pbUser := converters.UserToPBUser(m.User)
				f := &binary_files.Files{UserID: pbUser.GetUserId(), Name: name_, File: file, Extension: []byte(extension), UpdatedAt: time.Now()}
				pbBinaryFile, err := converters.BinaryFileToPBBinaryFile(f)
				if err != nil {
					m.Err = err
					m.selectedFile = ""
					m.FileUpdateInputs[0].Blur()
					m.FileUpdateInputs[0].Reset()
					return m, nil
				}

				// update file
				_, err = m.Client.BinaryFilesClient.Update(ctx, &pb.UpdateBinaryFileRequest{User: pbUser, File: pbBinaryFile})
				if err != nil {
					m.Err = err
					m.selectedFile = ""
					m.FileUpdateInputs[0].Blur()
					m.FileUpdateInputs[0].Reset()
					return m, nil
				}

				// update state
				m.UpdateResult = "BinaryFile updated successfully!"
				m.Step = "Binary File_UPDATE"
				m.FileUpdateInputs[0].Reset()
				m.selectedFile = ""
				m.filepicker.FileSelected = ""
				return m, nil
			}
		}
	}

	if m.selectedFile != "" {
		return updateBinaryFileUpdate(m, msg)
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

// updateBinaryFileGet function to update get inputs.
func updateBinaryFileGet(m *BinaryFileModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.Step = "Binary File_GET"
			return m, nil
		}
	}

	m.FileGetInputs[0], cmd = m.FileGetInputs[0].Update(msg)
	return m, cmd
}

// UpdateGet method to update get view.
func (m *BinaryFileModel) UpdateGet(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.Finish = true
			return m, tea.Quit
		case "ctrl+z":
			m.GetResult = ""
			m.Step = "DataTypes"
			m.FileGetInputs[0].Reset()
		case "enter":
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// get name
			name_ := m.FileGetInputs[0].Value()

			// convert user to proto user
			pbUser := converters.UserToPBUser(m.User)

			// get binary file
			resp, err := m.Client.BinaryFilesClient.Get(ctx, &pb.GetRequest{User: pbUser, Name: name_})
			if err != nil {
				m.Err = err
				m.GetResult = ""
				m.FileGetInputs[0].Reset()
				return updateBinaryFileGet(m, msg)
			}

			// convert to binary file
			file, err := converters.PBBinaryFileToBinaryFile(pbUser.GetUserId(), resp.GetFile())
			if err != nil {
				m.Err = err
				m.GetResult = ""
				m.FileGetInputs[0].Reset()
				return updateBinaryFileGet(m, msg)
			}

			// create file at home dir
			home, _ := os.UserHomeDir()
			path := fmt.Sprintf("%s/%s%s", home, name_, file.Extension)
			f, err := os.Create(path)
			if err != nil {
				m.Err = err
				m.GetResult = ""
				m.FileGetInputs[0].Reset()
				return updateBinaryFileGet(m, msg)
			}

			// write file
			buf := bytes.NewBuffer(file.File)
			_, err = buf.WriteTo(f)
			if err != nil {
				m.Err = err
				m.GetResult = ""
				m.FileGetInputs[0].Reset()
				return updateBinaryFileGet(m, msg)
			}

			// update state
			m.GetResult = fmt.Sprintf("BinaryFile saved to %s!", path)
			m.Step = "Binary File_GET"
			m.Err = nil
			m.FileGetInputs[0].Reset()
			return m, nil
		}
		m.GetResult = ""

	}
	return updateBinaryFileGet(m, msg)
}

// updateBinaryFileDelete function to update delete inputs.
func updateBinaryFileDelete(m *BinaryFileModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.Step = "Binary File_DELETE"
			return m, nil
		}
	}

	m.FileDeleteInputs[0], cmd = m.FileDeleteInputs[0].Update(msg)
	return m, cmd
}

// UpdateDelete method to update delete view.
func (m *BinaryFileModel) UpdateDelete(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.Finish = true
			return m, tea.Quit
		case "ctrl+z":
			m.DeleteResult = ""
			m.Step = "DataTypes"
			m.FileDeleteInputs[0].Reset()
		case "enter":
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// get name
			name_ := m.FileDeleteInputs[0].Value()

			// convert user to proto user
			pbUser := converters.UserToPBUser(m.User)

			// delete file
			_, err := m.Client.BinaryFilesClient.Delete(ctx, &pb.DeleteBinaryFileRequest{User: pbUser, Name: name_})
			if err != nil {
				m.Err = err
				m.DeleteResult = ""
				m.FileDeleteInputs[0].Reset()
				return updateBinaryFileDelete(m, msg)
			}

			// update state
			m.DeleteResult = "BinaryFile deleted successfully!"
			m.Step = "Binary File_DELETE"
			m.Err = nil
			return m, nil
		}
		m.DeleteResult = ""

	}
	return updateBinaryFileDelete(m, msg)
}

// readFile reads file by path.
func (m *BinaryFileModel) readFile(path string) ([]byte, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return f, nil
}
