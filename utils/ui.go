package utils

import (
	"fmt"
	"github.com/gookit/color"
	"strings"
)

func DisplayTitle(title string) string {
	spacer := fmt.Sprintf("%s", color.New(color.FgGreen, color.BgCyan).Sprintf(strings.Repeat("=", len(title))))
	titleColored := fmt.Sprintf("%s", color.New(color.FgBlack, color.BgCyan).Sprintf(title))

	return fmt.Sprintf("\n%s\n%s\n%s\n", spacer, titleColored, spacer)
}
