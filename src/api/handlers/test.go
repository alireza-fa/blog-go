package handlers

import "net/http"

type TestHandler struct{}

func (handler TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Test handler" + r.Method))
	if err != nil {
		panic(err)
	}
}
