package tongue

import "fmt"

type Dict struct {
    mappa map[string]map[string]string
}

func (d *Dict) Map () map[string]map[string]string {
    return d.mappa
}

func (d *Dict) Add (lang string, id string, sentence string) {
    if d.mappa == nil {
        d.mappa = make(map[string]map[string]string)
    }

    if d.mappa[lang] == nil {
        d.mappa[lang] = make(map[string]string)
    }

    d.mappa[lang][id] = sentence
}

func (d *Dict) Get (lang string, id string) string {
    fmt.Println("!!!")
    fmt.Println(d.mappa)
    fmt.Println(d.mappa[lang])
    fmt.Println(d.mappa[lang][id])

    return d.mappa[lang][id]
}

func LoadDictionary() Dict {
    return Dict{}
}
