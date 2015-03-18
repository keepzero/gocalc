package gocalc

var Cache bool = false

var rpnCaches map[string]string

func init() {
	rpnCaches = make(map[string]string)
}
