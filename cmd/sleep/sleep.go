package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func usage() {
	fmt.Printf("usage: sleep seconds\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	seconds, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
	}

	time.Sleep(time.Duration(seconds) * time.Second)
}
