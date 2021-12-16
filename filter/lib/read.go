package lib

import (
	"bufio"
	"os"
)

func Read(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	rows := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
	}
	return rows, nil
}
