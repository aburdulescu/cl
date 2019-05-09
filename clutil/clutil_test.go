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

	testData := []struct {
		input    string
		expected string
	}{
		{"bla bla foo bla", fmt.Sprintf("bla bla \x1b[%dmfoo\x1b[0m bla", color.FgBlue)},
		{"bla bla bar bla", fmt.Sprintf("bla bla \x1b[%dmbar\x1b[0m bla", color.FgGreen)},
	}

	for _, v := range testData {
		output := ColorLine(colors, v.input)
		if output != v.expected {
			t.Errorf("output doesn't match expected!\noutput: %s\nexpected: %s", output, v.expected)
		}
	}
}
