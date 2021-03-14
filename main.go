package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/tabwriter"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	flag.Usage = CustomUsage

	for i := range flags {
		flag.StringVar(&flags[i].Pattern, flags[i].Name, "", flags[i].Usage)
	}

	var printVersion bool
	flag.BoolVar(&printVersion, "v", false, "print version")

	var printColorPalette bool
	flag.BoolVar(&printColorPalette, "print-palette", false, "print available color palette")

	var exportFilter bool
	flag.BoolVar(&exportFilter, "export-filter", false, "export filter specified by provided flags")

	flag.Parse()

	if printVersion {
		fmt.Println(version)
		return nil
	}

	if printColorPalette {
		w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
		for _, f := range flags {
			fmt.Fprintf(w, "%s\t%scolored text%s\n", f.Name, f.Color, Reset)
		}
		w.Flush()
		return nil
	}

	if exportFilter {
		f, err := os.Create("filter.cl")
		if err != nil {
			return err
		}
		defer f.Close()
		for i := range flags {
			if flags[i].Pattern == "" {
				continue
			}
			f.WriteString(flags[i].Name + "=" + flags[i].Pattern + "\n")
		}
		return nil
	}

	if len(os.Args) == 1 {
		return fmt.Errorf("no flags provided")
	}

	// TODO: detect duplicate patterns: --red xyz --green xyz

	colors := []Color{}
	for _, f := range flags {
		if f.Pattern == "" {
			continue
		}
		re, err := regexp.Compile(f.Pattern)
		if err != nil {
			return err
		}
		colors = append(colors, Color{re, f.Color})
	}

	var scanner *bufio.Scanner
	if len(flag.Args()) == 0 {
		scanner = bufio.NewScanner(bufio.NewReader(os.Stdin))
	} else {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			return err
		}
		defer file.Close()
		scanner = bufio.NewScanner(bufio.NewReader(file))
	}

	w := bufio.NewWriter(os.Stdout)
	for scanner.Scan() {
		line := colorLine(colors, scanner.Text())
		if _, err := w.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	w.Flush()
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

type Color struct {
	re    *regexp.Regexp
	color string
}

type Position struct {
	position int
	color    string
}

type Positions []Position

func (p Positions) Len() int           { return len(p) }
func (p Positions) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Positions) Less(i, j int) bool { return p[i].position < p[j].position }

func colorLine(colors []Color, line string) string {
	var positions []Position
	var extraLen int

	for _, c := range colors {
		p := c.re.FindAllIndex([]byte(line), -1)
		for i := range p {
			positions = append(positions, Position{p[i][0], c.color})
			extraLen += len(c.color)
			positions = append(positions, Position{p[i][1], Reset})
			extraLen += len(Reset)
		}
	}

	if positions == nil {
		return line
	}

	sort.Sort(Positions(positions))

	var output strings.Builder
	output.Grow(len(line) + extraLen)

	var last int
	for i := range positions {
		output.WriteString(line[last:positions[i].position])
		last = positions[i].position
		output.WriteString(positions[i].color)
	}

	if last < len(line) {
		output.WriteString(line[last:])
	}

	return output.String()
}

func CustomUsage() {
	w := tabwriter.NewWriter(os.Stderr, 0, 4, 2, ' ', 0)

	fmt.Fprintf(w, "Usage: %s [OPTION]|[--COLOR=PATTERN]... [FILE]\n", os.Args[0])

	validOptions := map[string]bool{
		"v":             true,
		"print-palette": true,
		"export-filter": true,
	}

	options := []*flag.Flag{}
	colorFlags := []*flag.Flag{}
	flag.VisitAll(func(f *flag.Flag) {
		if _, ok := validOptions[f.Name]; ok {
			options = append(options, f)
		} else {
			colorFlags = append(colorFlags, f)
		}
	})

	fmt.Fprintf(w, "\nOPTIONS:\n")
	for _, f := range options {
		var flagPrefix string
		if len(f.Name) > 1 {
			flagPrefix = "--"
		} else {
			flagPrefix = "-"
		}
		fmt.Fprintf(w, "\t%s%s\t%s\n", flagPrefix, f.Name, f.Usage)
	}

	fmt.Fprintf(w, "\nCOLOR FLAGS:\n")
	for _, f := range colorFlags {
		fmt.Fprintf(w, "\t--%s\t%s\n", f.Name, f.Usage)
	}

	w.Flush()
}
