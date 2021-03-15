package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

func customUsage() {
	w := tabwriter.NewWriter(os.Stderr, 0, 4, 2, ' ', 0)

	fmt.Fprintf(w, "Usage: %s [OPTION]|[--COLOR=PATTERN]... [FILE]\n", os.Args[0])

	validOptions := map[string]bool{
		"v":             true,
		"print-palette": true,
		"export-filter": true,
		"filter":        true,
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
