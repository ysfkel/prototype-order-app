package data_reader

import (
	"bufio"
	"os"
)

func openFile(filePath string) (*os.File, error) {

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	return file, nil

}

func getScanner(file *os.File) *bufio.Scanner {

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return scanner
}
