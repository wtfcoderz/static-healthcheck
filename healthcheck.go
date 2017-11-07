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

	for _, tcp_addr := range tcpFlags {
		conn, err := net.Dial("tcp", tcp_addr)
		if err != nil {
			log.Println("Connection error:", err)
			log.Println(tcp_addr, ": Unreachable")
			os.Exit(1)
		} else {
			defer conn.Close()
			log.Println(tcp_addr, ": Online")
		}
	}

	for _, http_url := range httpFlags {
		resp, err := http.Get(http_url)
		if err != nil {
			// handle error
			log.Println("Connection error:", err)
			log.Println(http_url, ": Unreachable")
			os.Exit(1)
		} else {
			if resp.StatusCode < 400 {
				log.Println(http_url, ": Online - response:", resp.StatusCode)
			} else {
				log.Println(http_url, ": Error - response:", resp.StatusCode)
				os.Exit(1)
			}
		}
	}

}
