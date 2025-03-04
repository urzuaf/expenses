package barchart

import (
	"fmt"

	"github.com/NimbleMarkets/ntcharts/barchart"
	"github.com/charmbracelet/lipgloss"
)

func getColor(idx int) string {

	colors := []string{
		"#A7528D", // darker 3
		"#BE5D9F", // darker 2
		"#D567B0", // darker 1
		"#E37AC1", // base
		"#E992CD", // lighter 1
		"#F0AAD9", // lighter 2
		"#F6C2E5", // lighter 3
	}
	return colors[idx]

}

func DisplayBarchart(title string, data map[string]int) {
	var values []barchart.BarData
	var count int
	for key, value := range data {
		aux := barchart.BarData{
			Label: key,
			Values: []barchart.BarValue{
				{
					Name:  fmt.Sprintf("(%d)", value),
					Value: float64(value),
					Style: lipgloss.NewStyle().Foreground(lipgloss.Color(getColor(count)))}},
		}
		values = append(values, aux)
		count++
	}

	bc := barchart.New(64, 10)
	bc.PushAll(values)
	bc.Draw()

	fmt.Println(title)
	fmt.Println()
	fmt.Println(bc.View())
}
