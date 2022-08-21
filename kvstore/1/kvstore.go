package kvstore

var data = map[string]string{} // no! data race!

func Set(k, v string) {
	data[k] = v
}

func Get(k string) string {
	return data[k]
}
