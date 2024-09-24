package ittaxid

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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

func extractDayOfBirthAndSex(taxId string) (day int, isMale bool, err error) {
	taxIdPart, _ := strconv.ParseInt(taxId[9:11], 10, 64)

	if taxIdPart < 1 || taxIdPart > 71 {
		err = fmt.Errorf("day of birth not valid: %d", taxIdPart)
		return
	}

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

func extractYearOfBirth(taxId string) ([]string, error) {
	// extract the year code from the taxId from chars 7 to 8
	if len(taxId) != 16 {
		return nil, fmt.Errorf("taxId must be 16 characters long")
	}

	yearCode := taxId[6:8]

	// calculate current date
	currentYear := int64(time.Now().Year())
	year1900, err := strconv.ParseInt("19"+yearCode, 10, 64)
	if err != nil {
		return nil, err
	}
	year2000, err := strconv.ParseInt("20"+yearCode, 10, 64)
	if err != nil {
		return nil, err
	}

	possibleYears := make([]string, 0)

	if year1900 >= currentYear-120 {
		possibleYears = append(possibleYears, strconv.Itoa(int(year1900)))
	}

	if year2000 <= currentYear {
		possibleYears = append(possibleYears, strconv.Itoa(int(year2000)))
	}

	return possibleYears, nil

}

func calculateControlDigit(s string) string {
	var c string
	var charPosPari, charPosDispari strings.Builder
	counter := 0

	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			charPosDispari.WriteByte(s[i]) // perché per l'algoritmo la stringa comincia da 1 e non da 0
		} else {
			charPosPari.WriteByte(s[i])
		}
	}

	for i := 0; i < charPosDispari.Len(); i++ {
		switch charPosDispari.String()[i] {
		case '0', 'A':
			counter += 1
		case '1', 'B':
			counter += 0
		case '2', 'C':
			counter += 5
		case '3', 'D':
			counter += 7
		case '4', 'E':
			counter += 9
		case '5', 'F':
			counter += 13
		case '6', 'G':
			counter += 15
		case '7', 'H':
			counter += 17
		case '8', 'I':
			counter += 19
		case '9', 'J':
			counter += 21
		case 'K':
			counter += 2
		case 'L':
			counter += 4
		case 'M':
			counter += 18
		case 'N':
			counter += 20
		case 'O':
			counter += 11
		case 'P':
			counter += 3
		case 'Q':
			counter += 6
		case 'R':
			counter += 8
		case 'S':
			counter += 12
		case 'T':
			counter += 14
		case 'U':
			counter += 16
		case 'V':
			counter += 10
		case 'W':
			counter += 22
		case 'X':
			counter += 25
		case 'Y':
			counter += 24
		case 'Z':
			counter += 23
		}
	}

	for i := 0; i < charPosPari.Len(); i++ {
		switch charPosPari.String()[i] {
		case '0':
		case '1', 'B':
			counter += 1
		case '2', 'C':
			counter += 2
		case '3', 'D':
			counter += 3
		case '4', 'E':
			counter += 4
		case '5', 'F':
			counter += 5
		case '6', 'G':
			counter += 6
		case '7', 'H':
			counter += 7
		case '8', 'I':
			counter += 8
		case '9', 'J':
			counter += 9
		case 'A':
		case 'K':
			counter += 10
		case 'L':
			counter += 11
		case 'M':
			counter += 12
		case 'N':
			counter += 13
		case 'O':
			counter += 14
		case 'P':
			counter += 15
		case 'Q':
			counter += 16
		case 'R':
			counter += 17
		case 'S':
			counter += 18
		case 'T':
			counter += 19
		case 'U':
			counter += 20
		case 'V':
			counter += 21
		case 'W':
			counter += 22
		case 'X':
			counter += 23
		case 'Y':
			counter += 24
		case 'Z':
			counter += 25
		}
	}

	switch counter % 26 {
	case 0:
		c = "A"
	case 1:
		c = "B"
	case 2:
		c = "C"
	case 3:
		c = "D"
	case 4:
		c = "E"
	case 5:
		c = "F"
	case 6:
		c = "G"
	case 7:
		c = "H"
	case 8:
		c = "I"
	case 9:
		c = "J"
	case 10:
		c = "K"
	case 11:
		c = "L"
	case 12:
		c = "M"
	case 13:
		c = "N"
	case 14:
		c = "O"
	case 15:
		c = "P"
	case 16:
		c = "Q"
	case 17:
		c = "R"
	case 18:
		c = "S"
	case 19:
		c = "T"
	case 20:
		c = "U"
	case 21:
		c = "V"
	case 22:
		c = "W"
	case 23:
		c = "X"
	case 24:
		c = "Y"
	case 25:
		c = "Z"
	}

	return c
}
