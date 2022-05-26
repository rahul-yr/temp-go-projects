package employee

import "net/http"

func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}

func Read(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("read"))
}
