package main

type Flag struct {
	Color   string
	Pattern string
	Name    string
	Usage   string
}

const (
	fRed = iota
	fGreen
	fYellow
	fBlue
	fPurple
	fCyan
	fGray
	fWhite
	fMax
)

var flags = func() []Flag {
	flags := make([]Flag, fMax)
	flags[fRed] = Flag{
		Color:   Red,
		Pattern: "",
		Name:    "red",
		Usage:   "Color given pattern with red",
	}
	flags[fGreen] = Flag{
		Color:   Green,
		Pattern: "",
		Name:    "green",
		Usage:   "Color given pattern with green",
	}
	flags[fYellow] = Flag{
		Color:   Yellow,
		Pattern: "",
		Name:    "yellow",
		Usage:   "Color given pattern with yellow",
	}
	flags[fBlue] = Flag{
		Color:   Blue,
		Pattern: "",
		Name:    "blue",
		Usage:   "Color given pattern with blue",
	}
	flags[fPurple] = Flag{
		Color:   Purple,
		Pattern: "",
		Name:    "purple",
		Usage:   "Color given pattern with purple",
	}
	flags[fCyan] = Flag{
		Color:   Cyan,
		Pattern: "",
		Name:    "cyan",
		Usage:   "Color given pattern with cyan",
	}
	flags[fGray] = Flag{
		Color:   Gray,
		Pattern: "",
		Name:    "gray",
		Usage:   "Color given pattern with gray",
	}
	flags[fWhite] = Flag{
		Color:   White,
		Pattern: "",
		Name:    "white",
		Usage:   "Color given pattern with white",
	}
	return flags
}()
