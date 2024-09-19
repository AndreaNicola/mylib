package files

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestReadCsv(t *testing.T) {
	tests := []struct {
		filename  string
		separator rune
		expected  []map[string]string
		wantErr   bool
	}{
		{
			filename:  "../testdata/files/test.csv",
			separator: ',',
			expected: []map[string]string{
				{"Name": "John", "Age": "30", "City": "New York"},
				{"Name": "Jane", "Age": "25", "City": "Los Angeles"},
			},
			wantErr: false,
		},
		{
			filename:  "../testdata/files/nonexistent.csv",
			separator: ',',
			expected:  nil,
			wantErr:   true,
		},
	}

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Current dir: %v\n", currentDir)
	fmt.Printf("Executable: %v\n", executable)

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := ReadCsv(tt.filename, tt.separator)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ReadCsv() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
