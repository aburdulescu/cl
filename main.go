package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type Color struct {
	Re   *regexp.Regexp
	Func func(a ...interface{}) string
}

func main() {
	flag.CommandLine.Usage = func() {
		header := `%s -x regex INPUT

Read INPUT and color the part of the line that matches regex with the
color specified by -x flag(see Colors).

Examples:
1) color the lines from file "file.txt" that end with "foo" with color blue:
    %s -b ".*foo$" file.txt

Colors:
`
		fmt.Fprintf(flag.CommandLine.Output(), header, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	type Flag struct {
		pattern   string
		colorAttr color.Attribute
	}

	flags := map[string]*Flag{
		"blue":    &Flag{colorAttr: color.FgBlue},
		"cyan":    &Flag{colorAttr: color.FgCyan},
		"green":   &Flag{colorAttr: color.FgGreen},
		"magenta": &Flag{colorAttr: color.FgMagenta},
		"red":     &Flag{colorAttr: color.FgRed},
		"yellow":  &Flag{colorAttr: color.FgYellow},
	}
	flag.StringVar(&flags["blue"].pattern, "b", "", "color blue")
	flag.StringVar(&flags["cyan"].pattern, "c", "", "color cyan")
	flag.StringVar(&flags["green"].pattern, "g", "", "color green")
	flag.StringVar(&flags["magenta"].pattern, "m", "", "color magenta")
	flag.StringVar(&flags["red"].pattern, "r", "", "color red")
	flag.StringVar(&flags["yellow"].pattern, "y", "", "color yellow")
	flag.Parse()

	if len(os.Args) == 1 {
		mainError(fmt.Errorf("no flags provided"))
	}

	var colors []Color
	for _, v := range flags {
		if v.pattern == "" {
			continue
		}
		re, err := regexp.Compile(v.pattern)
		if err != nil {
			mainError(err)
		}
		color := Color{Re: re, Func: color.New(v.colorAttr).SprintFunc()}
		colors = append(colors, color)
	}

	var scanner *bufio.Scanner
	if len(flag.Args()) == 0 {
		scanner = bufio.NewScanner(bufio.NewReader(os.Stdin))
	} else {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			mainError(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(bufio.NewReader(file))
	}

	color.NoColor = false // always output color
	w := bufio.NewWriter(os.Stdout)
	for scanner.Scan() {
		_, err := w.WriteString(colorLine(colors, scanner.Text()) + "\n")
		if err != nil {
			mainError(err)
		}
	}
	w.Flush()
	if err := scanner.Err(); err != nil {
		mainError(err)
	}
}

func mainError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}

func contains(idxs [][]int, i int) bool {
	for _, idx := range idxs {
		if i >= idx[0] && i < idx[1] {
			return true
		}
	}
	return false
}

func colorLine(colors []Color, line string) string {
	var idxPairs [][]int
	var f func(a ...interface{}) string
	for _, c := range colors {
		idxPairs = c.Re.FindAllIndex([]byte(line), -1)
		if idxPairs != nil {
			f = c.Func
			break
		}
	}
	if idxPairs == nil {
		return line
	}

	var output strings.Builder
	output.Grow(len(line))

	idxs := make([]int, len(idxPairs)*2)
	k := 0
	for _, idx := range idxPairs {
		idxs[k] = idx[0]
		k++
		idxs[k] = idx[1]
		k++
	}

	if idxs[0] != 0 {
		output.WriteString(line[:idxs[0]])
	}
	coloredStr := f(line[idxs[0]:idxs[1]])
	for start, end := 0, 1; end < len(idxs); {
		if start%2 == 0 {
			output.WriteString(coloredStr)
		} else {
			output.WriteString(line[idxs[start]:idxs[end]])
		}
		start++
		end++
	}
	if idxs[len(idxs)-1] < len(line) {
		output.WriteString(line[idxs[len(idxs)-1]:len(line)])
	}

	return output.String()
}
