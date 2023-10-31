package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"asciifs/asciifs"
)

func main() {
	input := os.Args[1:]
	filename := ""
	banner := ""
	inputStr := ""

	if len(input) == 1 {
		inputStr = input[0]
		banner = "standard"
	} else if len(input) == 2 {
		inputStr = input[0]
		banner = input[1]
	} else {
		reg := regexp.MustCompile(`--output=\s*([^"]*).txt\s*([^"]*)`)

		command := input[0]
		if !reg.MatchString(command) {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\ngo run . --output=<fileName.txt> something standard")
			return
		}
		filename = command[9 : len(command)-4]
		inputStr = input[1]
		banner = input[2]
	}

	reg := regexp.MustCompile(`standard|shadow|thinkertoy`)

	if !reg.MatchString(banner) || len(input) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard\nOR\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	if !asciifs.ValidityChecker(inputStr) {
		fmt.Println("error: not in the ASCII range")
		return
	}

	data, err := asciifs.FileReader(banner)
	if err != nil {
		log.Fatalln(err)
	}
	mapData := asciifs.Mapper(data)
	output := asciifs.AsciiPrinter(mapData, inputStr)
	if len(filename) != 0 {
		output = output + "\n" + "\n"
		os.WriteFile(filename+".txt", []byte(output), 0o644)
	} else {
		fmt.Println(output)
	}
}
