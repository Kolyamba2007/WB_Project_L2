package main

import (
	"fmt"
	"os"

	ntp "github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "No response from NTP server: %v\n", err)
		os.Exit(111)
	}

	fmt.Println(time)
}
