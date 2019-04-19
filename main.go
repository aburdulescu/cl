package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/fatih/color"
)

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
		fmt.Fprintf(os.Stderr, "error: no flags provided\n\n")
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	type Color struct {
		Re   *regexp.Regexp
		Func func(a ...interface{}) string
	}

	var colors []Color
	for _, v := range flags {
		if v.pattern == "" {
			continue
		}
		re, err := regexp.Compile(v.pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot create color from given pattern: %s\n\n", v.pattern)
			flag.CommandLine.Usage()
			os.Exit(1)
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
			log.Fatal(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(bufio.NewReader(file))
	}

	color.NoColor = false // always output color

	for scanner.Scan() {
		l := scanner.Text()
		matchFound := false
		for _, c := range colors {
			loc := c.Re.FindIndex([]byte(l))
			if loc != nil {
				fmt.Printf("%s%s%s\n", l[:loc[0]], c.Func(l[loc[0]:loc[1]]), l[loc[1]:])
				matchFound = true
				break
			}
		}
		if !matchFound {
			fmt.Println(l)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
