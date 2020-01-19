package main

import (
	"fmt"
	"math"
	"strings"
)

func createAdidasAsciiLogo(width int) string {
	if width < 2 {
		panic("Error, minimum width is 2")
	}

	var sb strings.Builder
	numStripes := 3 // Can be any number of numStripes
	smallHeight := int(math.Round(math.Sqrt(float64(width))))
	totalHeight := smallHeight * numStripes

	// Some vars to speed up things a little bit
	stripe := strings.Repeat("@", width)
	separator := strings.Repeat(" ", smallHeight)
	stripes := make([]string, numStripes)
	for i := range stripes {
		stripes[i] = stripe
	}

	for h := 0; h < totalHeight; h++ {
		offset := h % smallHeight
		skip := numStripes - 1 - (h / smallHeight)

		sb.WriteString(strings.Repeat(" ", (skip*width)+offset))
		sb.WriteString(strings.Join(stripes[skip:], separator))

		// Avoid last carriage return
		if h != totalHeight-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func main() {

	widths := []int{2, 3, 5, 7, 9, 16, 21}

	for _, width := range widths {
		fmt.Println(fmt.Sprintf("\nadidas (width %d)", width))
		fmt.Println("\n-------------------------------------------------------------")
		fmt.Println(fmt.Sprintf("\n%s\n\n", createAdidasAsciiLogo(width)))
	}
}
