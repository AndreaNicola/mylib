package ittaxid

import (
	"fmt"
	"strconv"
	"strings"
)

func extractChars(taxId string) (consonants string, vowels string) {

	consonants = ""
	vowels = ""
	taxId = strings.ToUpper(taxId)

	for _, r := range taxId {

		// if r in not a letter, skip it
		if r < 'A' || r > 'Z' {
			continue
		}

		//if r is a vowel, add it to the vowels string
		if r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
			vowels += string(r)
		} else {
			//if r is a consonant, add it to the consonants string
			consonants += string(r)
		}

	}

	return
}

func extractDayOfBirthAndSex(taxId string) (day int, isMale bool) {
	taxIdPart, _ := strconv.ParseInt(taxId[9:11], 10, 64)
	isMale = taxIdPart < 40
	if taxIdPart > 40 {
		day = int(taxIdPart - 40)
	} else {
		day = int(taxIdPart)
	}
	return
}

func extractMonthOfBirth(taxId string) (string, error) {
	monthCode := taxId[8]

	switch monthCode {
	case 'A':
		return "01", nil
	case 'B':
		return "02", nil
	case 'C':
		return "03", nil
	case 'D':
		return "04", nil
	case 'E':
		return "05", nil
	case 'H':
		return "06", nil
	case 'L':
		return "07", nil
	case 'M':
		return "08", nil
	case 'P':
		return "09", nil
	case 'R':
		return "10", nil
	case 'S':
		return "11", nil
	case 'T':
		return "12", nil
	default:
		return "", fmt.Errorf("birth month not valid: %c", monthCode)
	}

}
