package main

import (
	"io"
	"net/http"
	"log"
	"encoding/json"
	"net"
)

// Echo http headers back
func HelloServer(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Header)
	log.Println(req.RemoteAddr)
	//io.WriteString(w, "hello, world!\n")
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.MarshalIndent(req.Header, "", "    ")
	if err != nil {
		io.WriteString(w, "error")
	} else {
		//log.Println(string(jsonData))
		io.WriteString(w, string(jsonData))
	}
}

// Echo http headers back
func LocServer(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Header)
	log.Println(req.RemoteAddr)
	if loc := req.Header.Get("Cf-Ipcountry"); loc == "" {
		io.WriteString(w, "IP region not found")
	} else {
		io.WriteString(w, loc)
	}
}

// Echo IP back
func EchoIPServer(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Header)
	log.Println(req.RequestURI)
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	if ip := req.Header.Get("Cf-Connecting-Ip"); ip != "" {
		io.WriteString(w, ip)
		return
	}
	if ip := req.Header.Get("X-Forwarded-For"); ip == "" {
		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			log.Fatal(err)
			return
		}
		io.WriteString(w, ip)
	} else {
		io.WriteString(w, ip)
	}
}

func main() {
	http.HandleFunc("/headers", HelloServer)
	http.HandleFunc("/loc", LocServer)
	http.HandleFunc("/", EchoIPServer)
	log.Fatal(http.ListenAndServe("172.17.0.1:8080", nil))
	//log.Fatal(http.ListenAndServe(":80", nil))
}