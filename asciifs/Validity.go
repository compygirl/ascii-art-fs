package asciifs

import (
	"fmt"
	"log"
)

func ValidityChecker(inputStr string) bool {
	count := 0
	count2 := 0
	if inputStr == "" {
		return false
	}
	for _, val := range inputStr {
		if val != '\n' && (val < ' ' || val > '~') {
			log.Print("Error: your string is not part of ascii table ")
			return false
		}
		if val == 'n' || val == '\n' {
			count2++
		}
		if val == '\\' {
			count++
		}
	}
	if count+count2 == len(inputStr) {
		for i := 0; i < count2; i++ {
			fmt.Println()
		}
		return false
	} else {
		return true
	}
}
