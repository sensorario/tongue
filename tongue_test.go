package tongue

import (
	"testing"
	"strings"
	"reflect"
	"fmt"
)

func TestDictionaryCanBeLoadedByFunc(t *testing.T) {
    d := LoadDictionary()
    d.Add("it", "foo", "bar")
    sentence := d.Get("it", "foo") 
    if sentence != "bar" {
        t.Fatal(strings.Join([]string{
            "Ops ",
            "'", sentence, "'", 
            " should be equal to bar",
        }, ""))
    }
}

func TestGenerateMap(t *testing.T) {
    d := LoadDictionary()
    d.Add("it", "foo", "bar")
    d.Add("en", "foo", "barr")

    mappa := make(map[string]map[string]string)
    mappa["it"] = make(map[string]string)
    mappa["it"]["foo"] = "bar"
    mappa["en"] = make(map[string]string)
    mappa["en"]["foo"] = "barr"

    if !reflect.DeepEqual(mappa, d.Map()) {
        t.Fatal("Oops! Generated map is not correct")
    }
}
