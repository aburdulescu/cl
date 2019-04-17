package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

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

	log.Printf("blue: %s, cyan: %s, green: %s, magenta: %s, red: %s, yellow: %s\n", blue, cyan, green, magenta, red, yellow)

	if blue != "" {
		color.Blue("blue: %s\n", blue)
	}
	if cyan != "" {
		color.Cyan("cyan: %s\n", cyan)
	}
	if green != "" {
		color.Green("green: %s\n", green)
	}
	if magenta != "" {
		color.Magenta("magenta: %s\n", magenta)
	}
	if red != "" {
		color.Red("red: %s\n", red)
	}
	if yellow != "" {
		color.Yellow("yellow: %s\n", yellow)
	}
}
