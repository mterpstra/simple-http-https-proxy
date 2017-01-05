package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

var host string

func handler(w http.ResponseWriter, r *http.Request) {

	// Dump the original request to the terminal
	println("\x1b[31m")
	println("-- In From Client --------------------------------------------------------------------------")
	b, _ := httputil.DumpRequest(r, true)
	println(string(b))
	println("--------------------------------------------------------------------------------------------\n\n")

	// Make the new request for the server using HTTPS
	req := &http.Request{
		Method:     r.Method,
		URL:        r.URL,
		Header:     r.Header,
		Proto:      r.Proto,
		ProtoMajor: r.ProtoMajor,
		ProtoMinor: r.ProtoMinor,
		Body:       r.Body,
	}

	// Override only two values
	req.URL.Scheme = "https"
	req.URL.Host = host

	// Dump the new request to the server out to the terminal
	println("\x1b[33m")
	println("-- Out to Server ---------------------------------------------------------------------------")
	b, _ = httputil.DumpRequest(req, true)
	println(string(b))
	println("--------------------------------------------------------------------------------------------\n\n")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// @todo:  Something should be sent down to the client
		println(err.Error())
		return
	}

	// Dump the server response
	println("\x1b[35m")
	println("-- Response from Server --------------------------------------------------------------------")
	b, _ = httputil.DumpResponse(resp, true)
	println(string(b))
	body, err := ioutil.ReadAll(resp.Body)
	println("--------------------------------------------------------------------------------------------\n\n")

	// Move the headers from the server response to the client response
	for key, value := range resp.Header {
		w.Header().Add(key, value[0])
	}
	w.WriteHeader(resp.StatusCode)

	// Write the response back to the client
	w.Write(body)
}

// @todo: Write some help commands
func main() {
	flag.Parse()
	args := flag.Args()
	host = args[0]
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
