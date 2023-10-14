package store

var storeMap = make(map[string]string)

func Set(alias string, url string) {
	storeMap[alias] = url
}

func Get(alias string) (string, bool) {
	url, ok := storeMap[alias]
	return url, ok
}
