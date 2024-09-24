# Ittaxid Package

The `ittaxid` package provides functions for validating Italian tax identification numbers (codice fiscale).

The Italian tax identification number is based on the name, surname, sex, birth date, and birth place of the person. In
some
cases, there may be multiple people with the same name, surname, and birth date, which results in the same tax id. This
is known as "omocodia." The `ittaxid` package does not handle omocodia cases.

Refer to the [Wikipedia page](https://en.wikipedia.org/wiki/https://en.wikipedia.org/wiki/Italian_fiscal_code) for more information on the
Italian tax identification number.

## Functions

### `Verify`

```go
func Verify(lastname, name, sex, birthDate, taxId string) error 
```

Verifies if the provided Italian tax identification number is valid. The function returns an error with a descroption if
the tax id is invalid.

#### Parameters

- `lastname` (string): The last name of the person.
- `name` (string): The first name of the person.
- `sex` (string {`M`, `F`})
- `birthDate` (string): The birth date of the person in the format `YYYY-MM-DD`.

#### Returns

- `error`: An error if the tax id is invalid.

#### Example

```go
err := ittaxid.Verify("Rossi", "Mario", "M", "1980-01-01", "RSSMRA80A01H501Z")
if err != nil {
log.Fatal(err)
}
```

### `Extract`

```go
func ExtractInfo(taxId string) (*Info, error)
```

Extracts the base information from an Italian tax identification number.

#### Parameters

- `taxId` (string): The Italian tax identification number to extract information from.

#### Returns

- `Info`: A struct containing the extracted information.
- `error`: An error if the tax id is invalid.

#### Example

```go
info, err := ittaxid.Extract("RSSMRA80A01H501Z")
if err != nil {
log.Fatal(err)
}

#### `Info` Struct

The `Info` struct contains the extracted information from an Italian tax identification number.

```go
type Info struct {
	Sex        string
	BirthDate  []string
	BirthPlace string
}
```

The fields of the `Info` struct are:

- `Sex` (string {`M`, `F`})
- `BirthDate` ([]string): The possibile birth date of the person in the format `YYYY-MM-DD`. When the codice fiscale
  standard was created in 1973, the birth date was not included in the tax id. In this case, the function returns a
  slice of possible birth dates. Kudos to the Italian bureaucracy for creating such a mess.
- `BirthPlace` (string): The birth place of the person. The birth place is extracted from the tax id and is in "Codice
  Catastale" format.