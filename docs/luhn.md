# Luhn Package

The `luhn` package provides functions to calculate and validate numbers using the Luhn algorithm.

## Functions

### `CalculateLuhn`

Calculates the Luhn check number for a given number.

```go
func CalculateLuhn(number int) int
```

**Parameters:**
- `number` (int): The number for which to calculate the Luhn check number.

**Returns:**
- `int`: The Luhn check number.

**Example:**
```go
checkNumber := luhn.CalculateLuhn(12345)
```

### `Valid`

Checks if a number is valid based on the Luhn algorithm.

```go
func Valid(number int) bool
```

**Parameters:**
- `number` (int): The number to validate.

**Returns:**
- `bool`: `true` if the number is valid, `false` otherwise.

**Example:**
```go
isValid := luhn.Valid(12345)
```

### `checksum`

Calculates the checksum for a given number using the Luhn algorithm.

```go
func checksum(number int) int
```

**Parameters:**
- `number` (int): The number for which to calculate the checksum.

**Returns:**
- `int`: The checksum.

**Example:**
```go
sum := luhn.checksum(12345)
```

## Usage

To use the `luhn` package, import it into your Go project and call the provided functions as needed.

```go
package main

import (
    "fmt"
    "luhn"
)

func main() {
    number := 12345
    checkNumber := luhn.CalculateLuhn(number)
    fmt.Println("Check Number:", checkNumber)

    isValid := luhn.Valid(number)
    fmt.Println("Is Valid:", isValid)
}
```

This documentation provides an overview of the functions available in the `luhn` package and how to use them7