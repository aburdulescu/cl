package cl

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
)

func TestColorLine(t *testing.T) {
	flags := map[string]*Flag{
		"blue":    {"blue", color.FgBlue},
		"cyan":    {"cyan", color.FgCyan},
		"green":   {"green", color.FgGreen},
		"magenta": {"magenta", color.FgMagenta},
		"red":     {"red", color.FgRed},
		"yellow":  {"yellow", color.FgYellow},
	}
	colors, err := CreateColors(flags)
	if err != nil {
		t.Fatalf("couldn't create colors: %v", err)
	}

	escS := "\x1b[%dm"
	escE := "\x1b[0m"
	escFmt := escS + "%s" + escE
	testData := []struct {
		input    string
		expected string
	}{
		{
			"bla" + flags["blue"].Pattern + "bla" + flags["blue"].Pattern + "bla",
			fmt.Sprintf(
				"bla"+escFmt+"bla"+escFmt+"bla",
				flags["blue"].ColorAttr, flags["blue"].Pattern,
				flags["blue"].ColorAttr, flags["blue"].Pattern,
			),
		},
		{
			"bla" + flags["cyan"].Pattern + "bla" + flags["cyan"].Pattern + "bla",
			fmt.Sprintf(
				"bla"+escFmt+"bla"+escFmt+"bla",
				flags["cyan"].ColorAttr, flags["cyan"].Pattern,
				flags["cyan"].ColorAttr, flags["cyan"].Pattern,
			),
		},
		{
			"bla" + flags["green"].Pattern + "bla" + flags["green"].Pattern + "bla",
			fmt.Sprintf(
				"bla"+escFmt+"bla"+escFmt+"bla",
				flags["green"].ColorAttr, flags["green"].Pattern,
				flags["green"].ColorAttr, flags["green"].Pattern,
			),
		},
		{
			"bla" + flags["magenta"].Pattern + "bla" + flags["magenta"].Pattern + "bla",
			fmt.Sprintf(
				"bla"+escFmt+"bla"+escFmt+"bla",
				flags["magenta"].ColorAttr, flags["magenta"].Pattern,
				flags["magenta"].ColorAttr, flags["magenta"].Pattern,
			),
		},
		{
			"bla" + flags["red"].Pattern + "bla" + flags["red"].Pattern + "bla",
			fmt.Sprintf(
				"bla"+escFmt+"bla"+escFmt+"bla",
				flags["red"].ColorAttr, flags["red"].Pattern,
				flags["red"].ColorAttr, flags["red"].Pattern,
			),
		},
		{
			"bla" + flags["yellow"].Pattern + "bla" + flags["yellow"].Pattern + "bla",
			fmt.Sprintf(
				"bla"+escFmt+"bla"+escFmt+"bla",
				flags["yellow"].ColorAttr, flags["yellow"].Pattern,
				flags["yellow"].ColorAttr, flags["yellow"].Pattern,
			),
		},
	}

	for _, v := range testData {
		output := ColorLine(colors, v.input)
		if output != v.expected {
			t.Errorf("output doesn't match expected!\noutput: %s\nexpected: %s", output, v.expected)
		}
	}
}
