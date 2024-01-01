package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "8080", "specify port")
	flag.Parse()
	localIp, err := getLocalAddress()
	if err != nil {
		fmt.Println("Error getting local IP", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test Server running on http://%s:%s", localIp, port)
	})
	fmt.Printf("Started webserver on http://%s:%s\n", localIp, port)
	err = http.ListenAndServe(localIp+":"+port, nil)
	if err != nil {
		fmt.Println("Error starting webServer:", err)
	}
}

func getLocalAddress() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localIP := conn.LocalAddr().(*net.UDPAddr)
	return strings.Split(localIP.String(), ":")[0], nil
}
