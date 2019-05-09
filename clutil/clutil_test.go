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
	testData := []struct {
		input    string
		expected string
	}{
		{
			"bla foo bla foo bla",
			fmt.Sprintf(
				"bla "+escS+"foo"+escE+" bla "+escS+"foo"+escE+" bla",
				color.FgBlue, color.FgBlue,
			),
		},
		{
			"bla bar bla bar bla",
			fmt.Sprintf(
				"bla "+escS+"bar"+escE+" bla "+escS+"bar"+escE+" bla",
				color.FgGreen, color.FgGreen,
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
