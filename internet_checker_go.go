package main

import (
	"flag"
	"fmt"
)

func main() {
	var icmp bool

	argURL := flag.String("url", ".", "Destination")
	argRetryCount := flag.String("retry", ".", "Retry count")
	flag.BoolVar(&icmp, "icmp", false, "Bool")
	flag.Parse()

	if icmp {
		fmt.Print("ICMP")
	} else {
		fmt.Println(*argURL)
		fmt.Println(*argRetryCount)
	}

}
