package optionlist

import (
	"fmt"

	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	result   string
}

func InitialState() Model {

	m := Model{
		choices:  options(),
		selected: make(map[int]struct{}),
	}

	return m

}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.result = m.choices[m.cursor]
			return m, tea.Quit
		}

	}
	return m, nil
}

func (m Model) View() string {
	var s string
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\n(esc to quit)\n"
	return s
}

func options() []string {
	return []string{"a", "b"}
}

func GetOption() string {
	choice := ""
	p := tea.NewProgram(InitialState())
	model, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	if m, ok := model.(Model); ok {
		choice = m.result
	}
	return choice
}
