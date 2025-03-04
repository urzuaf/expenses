package title

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func DisplayTitle() {
	titleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#E37AC1")).Bold(true)
	fmt.Printf("%s\nSelect an option:\n", titleStyle.Render("FinGo"))
}
