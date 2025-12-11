package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetLinesFromMultilineInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan opened file: %w", err)
	}

	return lines, nil
}

func GetLinesFromSingleLineBySeparator(path string, sep string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = strings.Split(scanner.Text(), sep)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan opened file: %w", err)
	}

	return lines, nil
}
