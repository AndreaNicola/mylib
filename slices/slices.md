# Package `slices`

The `slices` package provides functions for manipulating slices. It includes the following functionality:

- Calculating the difference between two slices

## Functions

### `Difference`

```go
func Difference(left, right []string) []string
```

Returns the difference between two slices.

#### Parameters

- `left`: The first slice to compare.
- `right`: The second slice to compare.

#### Returns

- `[]string`: A slice containing elements that are in the `left` slice but not in the `right` slice.

## Examples

### Calculating the difference between two slices

```go
left := []string{"a", "b", "c"}
right := []string{"b", "c", "d"}
diff := slices.Difference(left, right)
fmt.Println(diff) // Output: ["a"]
```