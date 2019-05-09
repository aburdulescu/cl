package clutil

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type Color struct {
	Re   *regexp.Regexp
	Func func(a ...interface{}) string
}

type Flag struct {
	Pattern   string
	ColorAttr color.Attribute
}

func contains(idxs [][]int, i int) bool {
	for _, idx := range idxs {
		if i >= idx[0] && i < idx[1] {
			return true
		}
	}
	return false
}

func idxPairsToIdxSlice(idxPairs [][]int) []int {
	idxs := make([]int, len(idxPairs)*2)
	k := 0
	for _, idx := range idxPairs {
		idxs[k] = idx[0]
		k++
		idxs[k] = idx[1]
		k++
	}
	return idxs
}

func ColorLine(colors []Color, line string) string {
	var idxPairs [][]int
	var f func(a ...interface{}) string
	for _, c := range colors {
		idxPairs = c.Re.FindAllIndex([]byte(line), -1)
		if idxPairs != nil {
			f = c.Func
			break
		}
	}
	if idxPairs == nil {
		return line
	}

	idxs := idxPairsToIdxSlice(idxPairs)

	var output strings.Builder
	output.Grow(len(line))

	if idxs[0] != 0 {
		output.WriteString(line[:idxs[0]])
	}
	coloredStr := f(line[idxs[0]:idxs[1]])
	for start, end := 0, 1; end < len(idxs); {
		if start%2 == 0 {
			output.WriteString(coloredStr)
		} else {
			output.WriteString(line[idxs[start]:idxs[end]])
		}
		start++
		end++
	}
	if idxs[len(idxs)-1] < len(line) {
		output.WriteString(line[idxs[len(idxs)-1]:])
	}

	return output.String()
}

func CreateColors(flags map[string]*Flag) ([]Color, error) {
	var colors []Color
	for _, v := range flags {
		if v.Pattern == "" {
			continue
		}
		re, err := regexp.Compile(v.Pattern)
		if err != nil {
			return nil, err
		}
		color := Color{
			Re:   re,
			Func: color.New(v.ColorAttr).SprintFunc(),
		}
		colors = append(colors, color)
	}
	return colors, nil
}

func FilterToFlags(filter string, flags map[string]*Flag) error {
	d, err := ioutil.ReadFile(filter)
	if err != nil {
		return err
	}
	data := make(map[string]string)
	err = json.Unmarshal(d, &data)
	if err != nil {
		return err
	}
	for k, v := range data {
		flags[k].Pattern = v
	}
	return nil
}
