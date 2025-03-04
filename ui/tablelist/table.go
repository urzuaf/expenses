package tablelist

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table    table.Model
	selected string
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "backspace":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "esc", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if !m.table.Focused() {
				return m, nil
			}
			m.selected = m.table.SelectedRow()[0]
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var s string
	s += baseStyle.Render(m.table.View()) + "\n"
	s += fmt.Sprintf("\n\n%s\n", "(esq to quit)")
	return s
}

func DisplayTable(columnNames []table.Column, rowData []table.Row) string {
	var choice string
	columns := columnNames

	rows := rowData

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	p := tea.NewProgram(model{table: t, selected: ""})
	m, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	if m, ok := m.(model); ok {
		choice = m.selected
	}
	return choice

}
