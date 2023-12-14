package utils

import (
	"fmt"
	"github.com/gookit/color"
	"strings"
)

func DisplayTitle(title string) {
	color.New(color.FgGreen, color.BgCyan).Println(strings.Repeat("=", len(title)))
	color.New(color.FgBlack, color.BgCyan).Println(title)
	color.New(color.FgGreen, color.BgCyan).Println(strings.Repeat("=", len(title)))
	fmt.Println("")
}
