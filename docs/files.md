# Package `files`

The `files` package provides functions for reading, writing, and comparing text files. It includes the following
functionalities:

- Writing all lines to a file
- Reading all lines from a file as a slice or map
- Comparing two files and writing the differences to a third file

## Functions

### SaveAllLines

```go
func SaveAllLines(fileName string, lines []string) error
```

Writes all provided lines to a specified file.

#### Parameters

- `fileName`: The name of the file to write the lines to.
- `lines`: A slice of strings to write to the file.

#### Returns

- `lines`: A slice of strings to write to the file.
- `error`: An error if there is an issue creating or writing to the file.

### ReadAllLinesAsMap

```go
func ReadAllLinesAsMap(fileName string) (map[string]interface{}, error)
```

Reads all lines from a file and returns them as a map.

#### Parameters

- `fileName`: The name of the file to read the lines from.

#### Returns

- `map[string]interface{}`: A map of the lines read from the file.
- `error`: An error if there is an issue opening or reading the file.

### ReadAllLines

```go
func ReadAllLines(fileName string) ([]string, error)
```

Reads all lines from a file and returns them as a slice.

#### Parameters

- `fileName`: The name of the file to read the lines from.

#### Returns

- `error`: An error if there is an issue opening or reading the file.

### `Difference`

Compares two files and writes the differences to a third file.

```go
func Difference(leftFile, rightFile, outputFile string) error
```

#### Parameters

- `leftFile`: The name of the first file to compare.
- `rightFile`: The name of the second file to compare.
- `outputFile`: The name of the file to write the differences to.

#### Returns

- `error`: An error if there is an issue opening or reading the file.

#### Usage

```go
err := files.Difference("file1.txt", "file2.txt", "diff.txt")
if err != nil {
log.Fatal(err)
}
```

### `ReadCsv`

The `ReadCsv` function reads a CSV file and returns its contents as a slice of maps, where each map represents a row in
the CSV file with column names as keys.

#### Parameters

- `filename`: The name of the CSV file to read.
- `separator`: The rune used as a separator in the CSV file.

#### Returns

- `[]map[string]string`: A slice of maps, where each map represents a row in the CSV file.
- `error`: An error if there is an issue opening or reading the CSV file.

#### Example Usage

```go
package main

import (
	"fmt"
	"log"
	"myproject/files"
)

func main() {
	records, err := files.ReadCsv("testdata/files/test.csv", ',')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records)
}
```