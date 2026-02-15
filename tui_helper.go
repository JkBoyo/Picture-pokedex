package main

import (
	"fmt"
	"strings"

	"github.com/leaanthony/go-ansi-parser"
)

var objTop = " _%s_ \n/ %s \\\n"
var objMiddle = "| %s |\n"
var objBottom = "\\_%s_/"

func PrintPokePage(info, image string) {
	surInfoLines := strings.Split(surroundObj("Info", info), "\n")
	surImageLines := strings.Split(surroundObj("Picture", image), "\n")

	endInfo := len(surInfoLines)
	infoWidth := len(surInfoLines[0])
	formattedPP := ""
	for i, line := range surInfoLines {
		formattedPP += fmt.Sprintf(" %s %s\n", line, surImageLines[i])
	}

	for i := endInfo - 1; i < len(surImageLines); i++ {
		formattedPP += fmt.Sprintf(
			" %s %s \n",
			strings.Repeat(" ", infoWidth),
			surImageLines[i],
		)
	}

	surPokePage := surroundObj("Poke Page", formattedPP)

	fmt.Println(surPokePage)
}

func surroundObj(title, str string) string {
	objLines := strings.Split(str, "\n")

	maxLen, err := findMaxLen(objLines)
	if err != nil {
		return "Error surrounding obj"
	}

	formattedStr := fmt.Sprintf(
		objTop,
		strings.Repeat("_", maxLen),
		" "+title+strings.Repeat(" ", maxLen-1-len(title)),
	)
	for _, line := range objLines {
		prLine, err := ansi.Length(line)
		if err != nil {
			return "Error Surrounding obj due to ansi length"
		}
		formattedStr += fmt.Sprintf(
			objMiddle,
			line+strings.Repeat(" ", maxLen-prLine),
		)
	}
	formattedStr += fmt.Sprintf(
		objBottom,
		strings.Repeat("_", maxLen),
	)
	return formattedStr
}

func findMaxLen(strs []string) (int, error) {
	maxLen := 0
	for _, str := range strs {
		length, err := ansi.Length(str)
		if err != nil {
			return 0, err
		}
		if length < maxLen {
			continue
		}
		maxLen = length
	}
	return maxLen, nil
}
