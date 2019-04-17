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
	Re    *regexp.Regexp
	Print func(format string, a ...interface{})
}

func main() {
	flag.CommandLine.Usage = func() {
		header := `%s [OPTIONS PATTERN] INPUT

Read INPUT and print the lines that match the PATTERN with the color specified by OPTION
E.g.:
1) print the lines from the file file.txt that contain the pattern "foo" with color blue:
    %s -b "foo" file.txt

Options:

`
		fmt.Fprintf(flag.CommandLine.Output(), header, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	var blue, cyan, green, magenta, red, yellow string
	flag.StringVar(&blue, "b", "", "blue - color the matching line with blue")
	flag.StringVar(&cyan, "c", "", "cyan - color the matching line with cyan")
	flag.StringVar(&green, "g", "", "green - color the matching line with green")
	flag.StringVar(&magenta, "m", "", "magenta - color the matching line with magenta")
	flag.StringVar(&red, "r", "", "red - color the matching line with red")
	flag.StringVar(&yellow, "y", "", "yellow - color the matching line with yellow")
	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "error: no flags provided\n\n")
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	var colors []Color
	if blue != "" {
		blueRE, err := regexp.Compile(blue)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n\n", err)
		} else {
			colors = append(colors, Color{Re: blueRE, Print: color.Blue})
		}
	}
	if cyan != "" {
		cyanRE, err := regexp.Compile(cyan)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n\n", err)
		} else {
			colors = append(colors, Color{Re: cyanRE, Print: color.Cyan})
		}
	}
	if green != "" {
		greenRE, err := regexp.Compile(green)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n\n", err)
		} else {
			colors = append(colors, Color{Re: greenRE, Print: color.Green})
		}
	}
	if magenta != "" {
		magentaRE, err := regexp.Compile(magenta)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n\n", err)
		} else {
			colors = append(colors, Color{Re: magentaRE, Print: color.Magenta})
		}
	}
	if red != "" {
		redRE, err := regexp.Compile(red)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n\n", err)
		} else {
			colors = append(colors, Color{Re: redRE, Print: color.Red})
		}
	}
	if yellow != "" {
		yellowRE, err := regexp.Compile(yellow)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n\n", err)
		} else {
			colors = append(colors, Color{Re: yellowRE, Print: color.Yellow})
		}
	}

	filecontent := `white line
blue line
cyan line
white line
green line
magenta line
white line
white line
white line
red line
white line
white line
yellow line
white line
`

	scanner := bufio.NewScanner(strings.NewReader(filecontent))
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
