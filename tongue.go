package tongue

type dict struct {
    mappa map[string]map[string]string
}

func (d * dict) Map () map[string]map[string]string {
    return d.mappa
}

func (d * dict) Add (lang string, id string, sentence string) {
    if d.mappa == nil {
        d.mappa = make(map[string]map[string]string)
    }

    d.mappa[lang] = make(map[string]string)
    d.mappa[lang][id] = sentence
}

func (d *dict) Get (lang string, id string) string {
    return d.mappa[lang][id]
}

func LoadDictionary() dict {
    return dict{}
}
