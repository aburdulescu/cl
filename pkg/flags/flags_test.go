package flags

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/fatih/color"
)

func TestFromFilterFilterFileDoesntExist(t *testing.T) {
	var flags Flags
	err := flags.FromFilter("foo.bar")
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
	err = flags.FromFilter(filterFileName)
	if err == nil {
		t.Errorf("FilterToFlags should return error if filter file has wrong content")
	}
}

func TestFromFilterSuccesfull(t *testing.T) {
	writtenFlags := Flags{
		"blue":   {".*foo", color.FgBlue},
		"yellow": {".*bar", color.FgYellow},
	}
	err := writtenFlags.ToFilter("example.cl")
	if err != nil {
		t.Fatal(err)
	}
	readFlags := make(Flags)
	err = readFlags.FromFilter("example.cl")
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(writtenFlags, readFlags) {
		t.Error("written flags aren't equal to read flags")
	}
}
