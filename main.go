package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime/pprof"
	"strings"

	"github.com/fatih/color"
)

type Color struct {
	Re   *regexp.Regexp
	Func func(a ...interface{}) string
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
	var idxs [][]int
	var f func(a ...interface{}) string
	for _, c := range colors {
		idxs = c.Re.FindAllIndex([]byte(line), -1)
		if idxs != nil {
			f = c.Func
			break
		}
	}
	if idxs == nil {
		return line
	}

	var output strings.Builder
	output.Grow(len(line))

	idxsSer := make([]int, len(idxs)*2)
	k := 0
	for _, idx := range idxs {
		idxsSer[k] = idx[0]
		k++
		idxsSer[k] = idx[1]
		k++
	}

	if idxsSer[0] != 0 {
		output.WriteString(line[:idxsSer[0]])
	}
	for start, end := 0, 1; end < len(idxsSer); {
		if start%2 == 0 {
			output.WriteString(f(line[idxsSer[start]:idxsSer[end]]))
		} else {
			output.WriteString(line[idxsSer[start]:idxsSer[end]])
		}
		start++
		end++
	}
	if idxsSer[len(idxsSer)-1] < len(line) {
		output.WriteString(line[idxsSer[len(idxsSer)-1]:len(line)])
	}

	return output.String()
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

	var cpuprof string
	flag.StringVar(&cpuprof, "cpuprof", "", "write cpu profile to file")
	flag.Parse()

	if cpuprof != "" {
		f, err := os.Create(cpuprof)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "error: no flags provided\n")
		os.Exit(1)
	}

	var colors []Color
	for _, v := range flags {
		if v.pattern == "" {
			continue
		}
		re, err := regexp.Compile(v.pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
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
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		scanner = bufio.NewScanner(bufio.NewReader(file))
	}

	color.NoColor = false // always output color
	w := bufio.NewWriter(os.Stdout)
	for scanner.Scan() {
		line := colorLine(colors, scanner.Text())
		var output strings.Builder
		fmt.Fprintf(&output, "%s\n", line)
		if _, err := w.WriteString(output.String()); err != nil {
			fmt.Fprintln(os.Stderr, "error: ", err)
			os.Exit(1)
		}
	}
	w.Flush()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}
}
