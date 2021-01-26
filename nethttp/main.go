package main

import "net/http"

func main() {

	http.HandleFunc("/", helloworld)
	http.HandleFunc("/sampletwo", sampletwo)
	http.HandleFunc("/sampleheader", sampleheader)
	http.HandleFunc("/updateapi", updatedapi)
	http.ListenAndServe(":8081", nil)

}
