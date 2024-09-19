package files

import "github.com/AndreaNicola/mylib/slices"

// Difference compares two files and writes the difference to a third file. The leftInput and rightInput are the files to compare, and the output is the file to write the difference to.
func Difference(leftFile, rightFile, outputFile string) error {
	left, err := ReadAllLines(leftFile)
	if err != nil {
		return err
	}

	right, err := ReadAllLines(rightFile)
	if err != nil {
		return err
	}

	output := slices.Difference(left, right)

	return SaveAllLines(outputFile, output)

}
