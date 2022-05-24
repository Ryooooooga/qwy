package main

import (
	"fmt"
)

var (
	version string = "dev"
	commit  string = "HEAD"
	date    string = "unknown"
)

func main() {
	fmt.Println("Hello, world!", version, commit, date)
}
