package cli

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mishankoGO/GophKeeper/internal/cli/binary_file"
	"github.com/mishankoGO/GophKeeper/internal/cli/card"
	"github.com/mishankoGO/GophKeeper/internal/cli/datatype"
	"github.com/mishankoGO/GophKeeper/internal/cli/index"
	"github.com/mishankoGO/GophKeeper/internal/cli/log_pass"
	"github.com/mishankoGO/GophKeeper/internal/cli/login"
	"github.com/mishankoGO/GophKeeper/internal/cli/tab"
	"github.com/mishankoGO/GophKeeper/internal/cli/text"
	"github.com/mishankoGO/GophKeeper/internal/client"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
	"os"
	"strings"
)

type Model struct {
	LoginModel      *login.LoginModel
	RegisterModel   *login.RegisterModel
	CardModel       *card.CardModel
	TextModel       *text.TextModel
	BinaryFileModel *binary_file.BinaryFileModel
	LogPassModel    *log_pass.LogPassModel
	TabModel        *tab.TabModel
	IndexModel      *index.IndexModel
	DataTypeModel   *datatype.DataTypeModel
	Cursor          int
	Finish          bool
	Err             error
	Step            string
	User            *users.User
	Client          *client.Client
}

func InitialModel(client *client.Client) *Model {

	loginModel := login.NewLoginModel(client)
	registerModel := login.NewRegisterModel(client)
	tabModel := tab.NewTabModel()
	cardModel := card.NewCardModel(client)
	textModel := text.NewTextModel(client)
	fileModel := binary_file.NewBinaryFileModel(client)
	logpassModel := log_pass.NewLogPassModel(client)
	indexModel := index.NewIndexModel()
	dataTypeModel := datatype.NewDataTypeModel()

	m := Model{
		LoginModel:      &loginModel,
		RegisterModel:   &registerModel,
		CardModel:       &cardModel,
		TextModel:       &textModel,
		BinaryFileModel: &fileModel,
		LogPassModel:    &logpassModel,
		TabModel:        &tabModel,
		IndexModel:      &indexModel,
		DataTypeModel:   &dataTypeModel,
		Err:             nil,
		Client:          client,
		Step:            "index",
	}
	p, _ := os.UserHomeDir()
	fmt.Println(p)
	p += "/go/GoLandProjects/GophKeeper"
	fmt.Println(p)
	return &m
}

/* MODEL */
func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.Step == "index" {
		m.IndexModel.Step = "index"
		_, cmd := m.IndexModel.Update(msg)
		m.Step = m.IndexModel.Step
		m.Finish = m.IndexModel.Finish
		return m, cmd
	} else if m.Step == "Login" {
		m.LoginModel.Step = "Login"
		_, cmd := m.LoginModel.Update(msg)
		m.Step = m.LoginModel.Step
		m.Finish = m.LoginModel.Finish
		m.User = m.LoginModel.User
		m.Err = m.LoginModel.Err
		return m, cmd
	} else if m.Step == "Register" {
		m.RegisterModel.Step = "Register"
		_, cmd := m.RegisterModel.Update(msg)
		m.Step = m.RegisterModel.Step
		m.Finish = m.RegisterModel.Finish
		m.Err = m.RegisterModel.Err
		return m, cmd
	} else if m.Step == "Tab" {
		m.TabModel.Step = "Tab"
		_, cmd := m.TabModel.Update(msg)
		m.Step = m.TabModel.Step
		m.DataTypeModel.Tab = m.TabModel.Tabs[m.TabModel.ActiveTab]
		m.Finish = m.TabModel.Finish
		return m, cmd
	} else if m.Step == "DataTypes" {
		m.DataTypeModel.Step = "DataTypes"
		_, cmd := m.DataTypeModel.Update(msg)
		m.Step = m.DataTypeModel.Step
		m.Finish = m.DataTypeModel.Finish
		return m, cmd
	} else if strings.Split(m.Step, "_")[0] == "Card" {
		m.CardModel.Step = m.Step
		m.CardModel.User = m.User
		_, cmd := m.CardModel.Update(msg)
		m.CardModel.User = m.User
		m.Step = m.CardModel.Step
		m.Finish = m.CardModel.Finish
		return m, cmd
	} else if strings.Split(m.Step, "_")[0] == "Text" {
		m.TextModel.Step = m.Step
		m.TextModel.User = m.User
		_, cmd := m.TextModel.Update(msg)
		m.TextModel.User = m.User
		m.Step = m.TextModel.Step
		m.Finish = m.TextModel.Finish
		return m, cmd
	} else if strings.Split(m.Step, "_")[0] == "Binary File" {
		m.BinaryFileModel.Step = m.Step
		m.BinaryFileModel.User = m.User
		_, cmd := m.BinaryFileModel.Update(msg)
		m.BinaryFileModel.User = m.User
		m.Step = m.BinaryFileModel.Step
		m.Finish = m.BinaryFileModel.Finish
		return m, cmd
	} else if strings.Split(m.Step, "_")[0] == "LogPass" {
		m.LogPassModel.Step = m.Step
		m.LogPassModel.User = m.User
		_, cmd := m.LogPassModel.Update(msg)
		m.LogPassModel.User = m.User
		m.Step = m.LogPassModel.Step
		m.Finish = m.LogPassModel.Finish
		return m, cmd
	}
	return m, nil
}

func (m Model) View() string {
	if m.Step == "Login" {
		return m.LoginModel.View()
	} else if m.Step == "Register" {
		return m.RegisterModel.View()
	} else if m.Step == "Tab" {
		return m.TabModel.View()
	} else if m.Step == "DataTypes" {
		return m.DataTypeModel.View()
	} else if strings.Split(m.Step, "_")[0] == "Card" {
		return m.CardModel.View()
	} else if strings.Split(m.Step, "_")[0] == "Text" {
		return m.TextModel.View()
	} else if strings.Split(m.Step, "_")[0] == "Binary File" {
		return m.BinaryFileModel.View()
	} else if strings.Split(m.Step, "_")[0] == "LogPass" {
		return m.LogPassModel.View()
	}
	return m.IndexModel.View()
}

func (m Model) GetUser() *users.User {
	return m.User
}
