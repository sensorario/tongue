package tongue

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

    d.mappa[lang] = make(map[string]string)
    d.mappa[lang][id] = sentence
}

func (d *Dict) Get (lang string, id string) string {
    return d.mappa[lang][id]
}

func LoadDictionary() Dict {
    return Dict{}
}
