package tongue

type Dict struct {
    sentences map[string]map[string]string
}

func (d *Dict) Map () map[string]map[string]string {
    return d.sentences
}

func (d *Dict) Add (lang string, id string, sentence string) {
    if d.sentences == nil {
        d.sentences = make(map[string]map[string]string)
    }

    if d.sentences[lang] == nil {
        d.sentences[lang] = make(map[string]string)
    }

    d.sentences[lang][id] = sentence
}

func (d *Dict) Get (lang string, id string) string {
    return d.sentences[lang][id]
}

func LoadDictionary() Dict {
    return Dict{}
}
