package textinput

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type Model struct {
	title     string
	textInput textinput.Model
	err       error
	done      bool
	Result    string
}

func InitialModel(title, placeholder string) Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Model{
		title:     title,
		textInput: ti,
		err:       nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.done = true
			m.Result = m.textInput.Value()
			return m, tea.Quit
		case "ctrl+c", "esc":
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n",
		m.title,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func GetInput(title, placeholder string) string {
	choice := ""
	p := tea.NewProgram(InitialModel(title, placeholder))
	model, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	if m, ok := model.(Model); ok {
		choice = m.Result
	}
	return choice
}
