package tongue

import (
	"testing"
	"strings"
	"reflect"
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
    d.Add("it", "fizz", "buzz")
    d.Add("en", "foo", "barr")
    d.Add("en", "fizz", "buzzz")

    mappa := make(map[string]map[string]string)
    mappa["it"] = make(map[string]string)
    mappa["it"]["foo"] = "bar"
    mappa["it"]["fizz"] = "buzz"
    mappa["en"] = make(map[string]string)
    mappa["en"]["foo"] = "barr"
    mappa["en"]["fizz"] = "buzzz"

    if !reflect.DeepEqual(mappa, d.Map()) {
        t.Fatal("Oops! Generated map is not correct")
    }
}

func TestReturnTranslation(t *testing.T) {
    d := LoadDictionary()
    d.Add("it", "foo", "bar")
    d.Add("it", "fizz", "buzz")
    d.Add("en", "foo", "barr")
    d.Add("en", "fizz", "buzzz")


    if d.Get("it", "foo") != "bar" {
        t.Fatal("Oops! Wrong translation found")
    }
}
