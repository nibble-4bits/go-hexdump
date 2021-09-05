package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("USAGE: hexdump [FILE]")
		os.Exit(1)
	}

	file := args[1]
	bytes, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("\n            ")
	for i := 0; i < 16; i++ {
		fmt.Printf("%02X ", i)
	}
	fmt.Printf("\n")

	for i := 0; i < len(bytes); i += 16 {
		fmt.Printf("\n%08X    ", i)

		for j := i; j < i+16; j++ {
			if j < len(bytes) {
				fmt.Printf("%02X ", bytes[j])
			} else {
				fmt.Printf("   ")
			}
		}

		fmt.Printf(" | ")
		for j := i; j < i+16; j++ {
			if j < len(bytes) {
				if bytes[j] >= 32 && bytes[j] <= 126 {
					fmt.Printf("%c", bytes[j])
				} else {
					fmt.Printf(".")
				}
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf(" |")
	}
	fmt.Printf("\n")
}
