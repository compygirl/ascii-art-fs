package asciifs

import (
	"bufio"
	"fmt"
	"os"
)

func FileReader(filename string) ([]string, error) {
	var data []string

	file, err := os.Open(filename + ".txt")
	if err != nil {
		return nil, fmt.Errorf("open file err: %v", err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	defer file.Close()
	return data, nil
}
