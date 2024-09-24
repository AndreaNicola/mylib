# MyLib

MyLib is a library where I store common functions that I use in my projects. The library is written in Go and is
organized into packages based on the functionality they provide. The packages are described below.

## Packages

### [Files](docs/files.md)

The `files` package provides functions for reading, writing, and comparing text files. It includes the following
functionalities:

- Writing all lines to a file
- Reading all lines from a file as a slice or map
- Comparing two files and writing the differences to a third file
- Reading a csv file and returning a slice of maps

### [Google](docs/google.md)

The `google` package provides functions for interacting with Google services.

### [Italian Tax Id](docs/ittaxid.md)

The `ittaxid` package provides functions for:

- extraction of base information from an Italian tax id;
- validation of an Italian tax id;

### [Luhn](docs/luhn.md)

The `luhn` package provides functions for generating and validating Luhn numbers.

### [Slices](docs/slices.md)

The `slices` package provides functions for manipulating slices. It includes the following functionalities:

- Difference between two strings slices

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
