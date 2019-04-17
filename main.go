package main

import (
	"flag"

	"github.com/fatih/color"
)

func main() {
	var red, green, yellow, blue string
	flag.StringVar(&red, "r", "", "red - color the matching line with red")
	flag.StringVar(&green, "g", "", "green - color the matching line with green")
	flag.StringVar(&yellow, "y", "", "yellow - color the matching line with yellow")
	flag.StringVar(&blue, "b", "", "blue - color the matching line with red")
	flag.Parse()

	if red != "" {
		color.Red("red: %s\n", red)
	}
	if green != "" {
		color.Green("green: %s\n", green)
	}
	if yellow != "" {
		color.Yellow("yellow: %s\n", yellow)
	}
	if blue != "" {
		color.Blue("blue: %s\n", blue)
	}
}
