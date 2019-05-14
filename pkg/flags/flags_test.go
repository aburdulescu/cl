package flags

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/fatih/color"
)

func TestFromFilterFilterFileDoesntExist(t *testing.T) {
	var flags Flags
	err := FromFilter("foo.bar", flags)
	if err == nil {
		t.Errorf("FilterToFlags should return error if filter file doesn't exist")
	}
}

func TestFromFilterJsonUmarshalFails(t *testing.T) {
	filterFileName := "example.cl"
	err := ioutil.WriteFile(filterFileName, []byte("foo"), os.FileMode(0666))
	if err != nil {
		t.Fatal(err)
	}
	var flags Flags
	err = FromFilter(filterFileName, flags)
	if err == nil {
		t.Errorf("FilterToFlags should return error if filter file has wrong content")
	}
	if err := os.Remove("example.cl"); err != nil {
		t.Fatal(err)
	}
}

func TestFromFilterSuccesfull(t *testing.T) {
	writtenFlags := Flags{
		"blue":   {".*foo", color.FgBlue},
		"yellow": {".*bar", color.FgYellow},
	}
	if err := ToFilter(writtenFlags, "example.cl"); err != nil {
		t.Fatal(err)
	}
	readFlags := Flags{
		"blue":   {ColorAttr: color.FgBlue},
		"yellow": {ColorAttr: color.FgYellow},
	}
	if err := FromFilter("example.cl", readFlags); err != nil {
		t.Fatal(err)
	}
	if !readFlags.Equals(writtenFlags) {
		t.Error("written flags aren't equal to read flags")
	}
	if err := os.Remove("example.cl"); err != nil {
		t.Fatal(err)
	}
}
