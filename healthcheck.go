package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var tcpFlags arrayFlags
var httpFlags arrayFlags

func main() {
	flag.Var(&tcpFlags, "tcp", "tcpcheck: ex: my.domain.com:80")
	flag.Var(&httpFlags, "http", "httpcheck: ex: my.domain.com:80")
	flag.Parse()

	for _, tcpAddr := range tcpFlags {
		conn, err := net.Dial("tcp", tcpAddr)
		if err != nil {
			log.Println("Connection error:", err)
			log.Println(tcpAddr, ": Unreachable")
			os.Exit(1)
		} else {
			defer conn.Close()
			log.Println(tcpAddr, ": Online")
		}
	}

	for _, httpURL := range httpFlags {
		resp, err := http.Get(httpURL)
		if err != nil {
			// handle error
			log.Println("Connection error:", err)
			log.Println(httpURL, ": Unreachable")
			os.Exit(1)
		} else {
			if resp.StatusCode < 400 {
				log.Println(httpURL, ": Online - response:", resp.StatusCode)
			} else {
				log.Println(httpURL, ": Error - response:", resp.StatusCode)
				os.Exit(1)
			}
		}
	}

}
