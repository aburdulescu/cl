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
		header := `%s [OPTIONS] -cx regex INPUT

Read INPUT and print the lines that match the PATTERN with the color specified by OPTION
E.g.:
1) print the lines from the file file.txt that contain the pattern "foo" with color blue:
    %s -b "foo" file.txt

Options:

`
		fmt.Fprintf(flag.CommandLine.Output(), header, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	type Flag struct {
		Pattern string
		Print   func(format string, a ...interface{})
	}

	flags := map[string]*Flag{
		"blue":    &Flag{Print: color.Blue},
		"cyan":    &Flag{Print: color.Cyan},
		"green":   &Flag{Print: color.Green},
		"magenta": &Flag{Print: color.Magenta},
		"red":     &Flag{Print: color.Red},
		"yellow":  &Flag{Print: color.Yellow},
	}
	flag.StringVar(&flags["blue"].Pattern, "cb", "", "blue color")
	flag.StringVar(&flags["cyan"].Pattern, "cc", "", "cyan color")
	flag.StringVar(&flags["green"].Pattern, "cg", "", "green color")
	flag.StringVar(&flags["magenta"].Pattern, "cm", "", "magenta color")
	flag.StringVar(&flags["red"].Pattern, "cr", "", "red color")
	flag.StringVar(&flags["yellow"].Pattern, "cy", "", "yellow color")

	var mode string
	flag.StringVar(&mode, "m", "line", "mode - select mode(line or match)")

	flag.Parse()

	switch {
	case len(os.Args) == 1:
		fmt.Fprintf(os.Stderr, "error: no flags provided\n\n")
		flag.CommandLine.Usage()
		os.Exit(1)
	case len(os.Args) == 3 && os.Args[1] == "-m":
		fmt.Fprintf(os.Stderr, "error: no color->pattern provided\n\n")
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	type Color struct {
		Re    *regexp.Regexp
		Print func(format string, a ...interface{})
	}

	var colors []Color
	for _, v := range flags {
		if v.Pattern == "" {
			continue
		}
		re, err := regexp.Compile(v.Pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot create color from given pattern: %s\n\n", v.Pattern)
			flag.CommandLine.Usage()
			os.Exit(1)
		}
		color := Color{Re: re, Print: v.Print}
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
			if c.Re.MatchString(l) {
				c.Print(l)
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
