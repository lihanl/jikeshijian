package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/healthz", myfunc)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func myfunc(w http.ResponseWriter, r *http.Request) {
	for k, _ := range r.Header {
		strs := r.Header[k]
		for i, _ := range strs {
			w.Header().Add(k, strs[i])
		}
	}
	resp := make(map[string]int)
	resp["message"] = http.StatusOK
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	log.Printf("client ip: %s, response status code: %s", r.RemoteAddr, strconv.Itoa(resp["message"]))

}
