package api

import (
	"net/http"
)

func handleTopPage(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		writer.WriteHeader(404)
		writer.Write(nil)
		return
	}
	if request.Method == "GET" || request.Method == "HEAD" {
		writer.Header().Add("Content-Type", "text/html; charset=utf-8")
		writer.WriteHeader(200)
		writer.Write([]byte(`<a href="https://neodb.net">NeoDB</a> relay`))
	} else {
		writer.WriteHeader(400)
		writer.Write(nil)
	}
}
