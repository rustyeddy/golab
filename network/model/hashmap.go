package model

// General purpose hash map with string index
type Hashmap map[string]interface{}

func (hm Hashmap) Fetch(name string) (interface{}, bool) {
	item, exists := hm[name]
	return item, exists
}

func (hm Hashmap) Exists(name string) bool {
	_, ex := hm.Fetch(name)
	return ex
}

func (hm Hashmap) Get(name string) interface{} {
	if it, ex := hm.Fetch(name); ex {
		return it
	}
	return nil
}

func (hm Hashmap) Set(name string, obj interface{}) {
	hm[name] = obj
}

func (hm Hashmap) NameValList() (names []string, vals []interface{}) {
	for n, v := range hm {
		names = append(names, n)
		vals = append(vals, v)
	}
	return names, vals
}

func (hm Hashmap) Names() (names []string) {
	names, _ = hm.NameValList()
	return names
}

func (hm Hashmap) Values() (vals []interface{}) {
	_, vals = hm.NameValList()
	return vals
}
