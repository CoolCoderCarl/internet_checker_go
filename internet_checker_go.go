package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	argURL := flag.String("url", ".", "Destination")
	argRetryCount := flag.String("retry", ".", "Retry count")
	argProtocolICMP := flag.String("icmp", ".", "Bool")
	flag.Parse()

	switch os.Args[1] {
	case "--url":
		fmt.Println(*argURL)
	case "--retry":
		fmt.Println(*argRetryCount)
	case "--icmp":
		fmt.Println(*argProtocolICMP)
	}

}
