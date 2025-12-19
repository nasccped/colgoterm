package main

import "fmt"

func intoBrightRed(s string) string {
	return fmt.Sprintf("\x1b[92m%s\x1b[0m", s)
}

func main() {
	fmt.Printf("Welcome to the %s!\n", intoBrightRed("colgoterm"))
}
