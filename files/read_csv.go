package files

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsv(filename string, separator rune) ([]map[string]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error while opening csv file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = separator

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error while reading csv file: %v", err)
	}

	fieldsName := records[0]
	recordsSlice := make([]map[string]string, len(records)-1)
	for i, record := range records[1:] {
		recordMap := make(map[string]string)
		recordsSlice[i] = recordMap
		for j, field := range record {
			recordMap[fieldsName[j]] = field
		}
	}

	return recordsSlice, nil

}
