package files

import (
	"bufio"
	"fmt"
	"os"
)

// ReadAllLinesAsMap reads all lines from a file and returns them as a map.
func ReadAllLinesAsMap(fileName string) (map[string]interface{}, error) {
	lines, err := ReadAllLines(fileName)
	if err != nil {
		return nil, err
	}

	linesMap := make(map[string]interface{}, len(lines))
	for _, line := range lines {
		linesMap[line] = nil
	}

	return linesMap, nil
}

// ReadAllLines reads all lines from a file and returns them as a map.
func ReadAllLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return lines, nil
}
