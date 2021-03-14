package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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

	if len(os.Args) == 1 {
		return fmt.Errorf("no flags provided")
	}

	type Color struct {
		re    *regexp.Regexp
		color string
	}
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
		// call re.FindAllIndex
		// add colors where indexes indicate
		_, err := w.WriteString(scanner.Text() + "\n")
		if err != nil {
			return err
		}
	}
	w.Flush()
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func CustomUsage() {
	w := tabwriter.NewWriter(os.Stderr, 0, 4, 2, ' ', 0)

	fmt.Fprintf(w, "Usage: %s [OPTION]|[--COLOR=PATTERN]... [FILE]\n", os.Args[0])

	validOptions := map[string]bool{"v": true, "print-palette": true}

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
		fmt.Fprintf(w, "\t--%s=regex\t%s\n", f.Name, f.Usage)
	}

	w.Flush()
}
