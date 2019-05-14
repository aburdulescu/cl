package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/aburdulescu/cl/pkg/cl"
	"github.com/aburdulescu/cl/pkg/flags"
	"github.com/fatih/color"
)

func main() {
	flag.CommandLine.Usage = usage

	f := flags.Flags{
		"blue":    {ColorAttr: color.FgBlue},
		"cyan":    {ColorAttr: color.FgCyan},
		"green":   {ColorAttr: color.FgGreen},
		"magenta": {ColorAttr: color.FgMagenta},
		"red":     {ColorAttr: color.FgRed},
		"yellow":  {ColorAttr: color.FgYellow},
	}
	flag.StringVar(&f["blue"].Pattern, "cb", "", "color blue")
	flag.StringVar(&f["cyan"].Pattern, "cc", "", "color cyan")
	flag.StringVar(&f["green"].Pattern, "cg", "", "color green")
	flag.StringVar(&f["magenta"].Pattern, "cm", "", "color magenta")
	flag.StringVar(&f["red"].Pattern, "cr", "", "color red")
	flag.StringVar(&f["yellow"].Pattern, "cy", "", "color yellow")
	var filter string
	flag.StringVar(&filter, "f", "", "apply color filter")
	flag.Parse()

	if len(os.Args) == 1 {
		mainError(fmt.Errorf("no flags provided"))
	}

	if filter != "" {
		err := flags.FromFilter(filter, f)
		if err != nil {
			mainError(err)
		}
	}

	colors, err := cl.CreateColors(f)
	if err != nil {
		mainError(err)
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
		_, err := w.WriteString(cl.ColorLine(colors, scanner.Text()) + "\n")
		if err != nil {
			mainError(err)
		}
	}
	w.Flush()
	if err := scanner.Err(); err != nil {
		mainError(err)
	}
}

func usage() {
	header := `%s -x regex INPUT

Read INPUT and color the part of the line that matches regex with the
color specified by -x flag(see Colors).

Examples:
1) color the lines from file "file.txt" that end with "foo" with color blue:
    %s -b ".*foo$" file.txt

Colors:
`
	fmt.Fprintf(flag.CommandLine.Output(),
		header, os.Args[0], os.Args[0])
	flag.PrintDefaults()

}

func mainError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
