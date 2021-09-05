package main

import (
	"fmt"
	"os"
)

const ROW_LENGTH = 16

func printHeaderOffset() {
	fmt.Printf("\n%12c", ' ')
	for i := 0; i < ROW_LENGTH; i++ {
		fmt.Printf("%02X ", i)
	}
	fmt.Printf("\n")
}

func printHexBytesRow(bytes []byte, offset int) {
	for i := offset; i < offset+ROW_LENGTH; i++ {
		if i < len(bytes) {
			fmt.Printf("%02X ", bytes[i])
		} else {
			fmt.Printf("%3c", ' ')
		}
	}
}

func printASCIIRow(bytes []byte, offset int) {
	fmt.Printf(" | ")
	for i := offset; i < offset+ROW_LENGTH; i++ {
		if i < len(bytes) {
			if bytes[i] >= 32 && bytes[i] <= 126 {
				fmt.Printf("%c", bytes[i])
			} else {
				fmt.Printf(".")
			}
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Printf(" |")
}

func printHexDump(bytes []byte) {
	printHeaderOffset()
	for i := 0; i < len(bytes); i += ROW_LENGTH {
		fmt.Printf("\n%08X%4c", i, ' ')

		printHexBytesRow(bytes, i)
		printASCIIRow(bytes, i)
	}
}

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

	printHexDump(bytes)

	fmt.Printf("\n")
}
