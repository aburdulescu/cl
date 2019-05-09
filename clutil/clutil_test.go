package clutil

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
)

func TestColorLine(t *testing.T) {
	flags := map[string]*Flag{
		"blue":  {"foo", color.FgBlue},
		"green": {"bar", color.FgGreen},
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
			"bla foo bla foo bla",
			fmt.Sprintf(
				"bla "+escFmt+" bla "+escFmt+" bla",
				flags["blue"].ColorAttr, flags["blue"].Pattern,
				flags["blue"].ColorAttr, flags["blue"].Pattern,
			),
		},
		{
			"bla bar bla bar bla",
			fmt.Sprintf(
				"bla "+escFmt+" bla "+escFmt+" bla",
				flags["green"].ColorAttr, flags["green"].Pattern,
				flags["green"].ColorAttr, flags["green"].Pattern,
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
