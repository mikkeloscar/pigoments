package main

import "fmt"

func main() {
	hex := 0xbeef
	if 060 > int(60) {
		fmt.Printf("value: %d", hex)
	}
}
