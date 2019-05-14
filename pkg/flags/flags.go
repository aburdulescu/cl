package flags

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

type Flag struct {
	Pattern   string
	ColorAttr color.Attribute
}

type Flags map[string]*Flag

func (f Flags) FromFilter(filter string) error {
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
		f[k].Pattern = v
	}
	return nil
}

func (f Flags) ToFilter(filter string) error {
	data := make(map[string]string)
	for k, v := range f {
		data[k] = v.Pattern
	}
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filter, b, os.FileMode(0666))
	if err != nil {
		return err
	}
	return nil
}

func (f0 Flags) Equals(f1 Flags) bool {
	for f1k, f1v := range f1 {
		f0v, ok := f0[f1k]
		if !ok {
			return false
		}
		if f0v.Pattern != f1v.Pattern {
			return false
		}
		if f0v.ColorAttr != f1v.ColorAttr {
			return false
		}
	}
	return true
}
