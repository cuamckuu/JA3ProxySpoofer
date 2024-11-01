package main

import (
	"io"
	"log"
	"strings"

	http "github.com/Danny-Dasilva/fhttp"
	//"net/http"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ja3 := "771,52393-52392-52244-52243-49195-49199-49196-49200-49171-49172-156-157-47-53-10,65281-0-23-35-13-5-18-16-30032-11-10,29-23-24,0"
	ua := "Chrome Version 57.0.2987.110 (64-bit) Linux"

	cycleClient := &http.Client{
		Transport: cycletls.NewTransport(ja3, ua),
	}

	resp, err := cycleClient.Get(strings.Replace(r.URL.String(), "http", "https", 1))
	if err != nil {
		log.Print("Request Failed: " + err.Error())
		http.Error(w, "Error sending proxy request", http.StatusInternalServerError)
		return
	}
	log.Println(resp)

	// Copy the headers from the proxy response to the original response
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, string(value))
		}
	}

	// Set the status code of the original response to the status code of the proxy response
	w.WriteHeader(resp.StatusCode)

	// Copy the body of the proxy response to the original response
	io.Copy(w, resp.Body)
}

func main() {
	// Create a new HTTP server with the handleRequest function as the handler
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handleRequest),
	}

	// Start the server and log any errors
	log.Println("Starting proxy server on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting proxy server: ", err)
	}
}
