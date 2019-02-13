package printer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/kylelemons/godebug/diff"
)

func detectNewline(text string) string {
	win, _ := regexp.MatchString(`\r\n`, text)
	if win {
		return "\r\n"
	}
	return "\n"
}

func colorDiffLine(str string) string {
	add, _ := regexp.MatchString(`^\+.*`, str)
	subtract, _ := regexp.MatchString(`^\-.*`, str)

	if add {
		return color.GreenString(str)
	} else if subtract {
		return color.RedString(str)
	}

	return str
}

func osFriendlyNewlineSplit(str string) []string {
	nl := detectNewline(str)
	return strings.Split(str, nl)
}

func prepareText(text string) string {
	lines := osFriendlyNewlineSplit(text)

	var buffer []string
	for _, line := range lines {
		buffer = append(buffer, colorDiffLine(line))

	}
	return strings.Join(buffer, "\n")
}

// PrintChange print the diff of changes
func PrintChange(original, new string) {
	diffText := diff.Diff(original, new)

	text := prepareText(diffText)
	if strings.ContainsAny(text, "+-") {
		fmt.Printf("\n%v\n\n", text)
	}
}
