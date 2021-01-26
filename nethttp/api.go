package main

import (
	"io/ioutil"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {

		//http.Error(w, "Not FOUND ", 400)
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Sample web server "))

}

//setting url path and send error status
func sampletwo(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/sampletwo" {

		http.Error(w, "Not FOUND ", 400)
		return
	}

	w.Write([]byte("Sample web server 2 "))

}

//adding header with request

func sampleheader(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/sampleheader" {

		http.Error(w, "Not FOUND ", 400)
		return
	}

	if r.Header.Get("Accept") == "application/json" {

		w.Header().Set("content-type", "application/json")

		w.Write([]byte(`{ "name" : "helloword"} `))

		return

	}

	w.Write([]byte("hello"))

}

//request that manage GET and POST method , and take data from URL

func updatedapi(w http.ResponseWriter, r *http.Request) {

	//for GET call

	switch r.Method {

	case "GET": //curl -is localhost:8081/updateapi?foo=bar

		value := r.URL.Query().Get("foo")

		w.Write([]byte(value))
		return

	case "POST": //curl -is -X POST -d "I am rajibul " localhost:8081/updateapi

		data, _ := ioutil.ReadAll(r.Body)
		w.Write(data)
		return

	default:

		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
		return
	}
}
