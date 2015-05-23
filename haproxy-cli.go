package main

import (
	"fmt"
	"os"

	hap "github.com/wayneeseguin/haproxy-cli/haproxy"
)

var haproxy *hap.Haproxy

func init() {
	socket := os.Getenv("HAPROXY_SOCK")
	if socket == "" {
		fmt.Printf("Error: HAPROXY_SOCK environment variable not set")
		os.Exit(1)
	}
	haproxy = &hap.Haproxy{Socket: socket}
}

func main() {
	action := os.Args[1]
	switch {
	case action == "stats":
		output, err := haproxy.Stats("all")
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%+v", output)
	case action == "info":
		output, err := haproxy.Info()
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%+v", output)
	default:
		fmt.Printf("Usage: {stats|info}")
		os.Exit(1)
	}
}
