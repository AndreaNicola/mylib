package ittaxid

import (
	"testing"
)

func TestExtractChars(t *testing.T) {
	tests := []struct {
		taxId      string
		wantCons   string
		wantVowels string
	}{
		{"RSSMRA85M01H501Z", "RSSMRMHZ", "A"},
		{"VRNGNI85M01H501Z", "VRNGNMHZ", "I"},
		{"", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.taxId, func(t *testing.T) {
			gotCons, gotVowels := extractChars(tt.taxId)
			if gotCons != tt.wantCons {
				t.Errorf("extractChars() gotCons = %v, want %v", gotCons, tt.wantCons)
			}
			if gotVowels != tt.wantVowels {
				t.Errorf("extractChars() gotVowels = %v, want %v", gotVowels, tt.wantVowels)
			}
		})
	}
}

func TestExtractDayOfBirthAndSex(t *testing.T) {
	tests := []struct {
		taxId      string
		wantDay    int
		wantIsMale bool
		wantErr    bool
	}{
		{"RSSMRA85M01H501Z", 1, true, false},
		{"rssmra85M23H501B", 23, true, false},
		{"RSSMRA85M41H501Z", 1, false, false},
		{"rssmra85M63H501B", 23, false, false},
		{"RSSMRA85M42H501Z", 2, false, false},
		{"RSSMRA85M72H501Z", 0, false, true},
		{"RSSMRA85M00H501Z", 0, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.taxId, func(t *testing.T) {
			gotDay, gotIsMale, gotErr := extractDayOfBirthAndSex(tt.taxId)
			if gotDay != tt.wantDay {
				t.Errorf("extractDayOfBirthAndSex() gotDay = %v, want %v", gotDay, tt.wantDay)
			}
			if gotIsMale != tt.wantIsMale {
				t.Errorf("extractDayOfBirthAndSex() gotIsMale = %v, want %v", gotIsMale, tt.wantIsMale)
			}
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("extractDayOfBirthAndSex() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestExtractMonthOfBirth(t *testing.T) {
	tests := []struct {
		taxId     string
		wantMonth string
		wantErr   bool
	}{

		{"RSSMRA85AA001Z", "01", false},
		{"RSSMRA85BA001Z", "02", false},
		{"RSSMRA85CA001Z", "03", false},
		{"RSSMRA85DA001Z", "04", false},
		{"RSSMRA85EA001Z", "05", false},
		{"RSSMRA85HA001Z", "06", false},
		{"RSSMRA85LA001Z", "07", false},
		{"RSSMRA85MA001Z", "08", false},
		{"RSSMRA85PA001Z", "09", false},
		{"RSSMRA85RA001Z", "10", false},
		{"RSSMRA85SA001Z", "11", false},
		{"RSSMRA85TA001Z", "12", false},
		{"RSSMRA85KA001Z", "", true},
		{"RSSMRA85ZA001Z", "", true},
		{"RSSMRA85YA001Z", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.taxId, func(t *testing.T) {
			gotMonth, err := extractMonthOfBirth(tt.taxId)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractMonthOfBirth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMonth != tt.wantMonth {
				t.Errorf("extractMonthOfBirth() gotMonth = %v, want %v", gotMonth, tt.wantMonth)
			}
		})
	}

}

func TestExtractYearOfBirth(t *testing.T) {
	tests := []struct {
		taxId     string
		wantYears []string
		wantErr   bool
	}{
		{"RSSMRA85M01A001Z", []string{"1985"}, false},
		{"RSSMRA00M01B001Z", []string{"2000"}, false},
		{"RSSMRA10M01B101Z", []string{"1910", "2010"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.taxId, func(t *testing.T) {
			gotYears, err := extractYearOfBirth(tt.taxId)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractYearOfBirth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotYears) != len(tt.wantYears) {
				t.Errorf("extractYearOfBirth() gotYears = %v, want %v", gotYears, tt.wantYears)
			}
		})
	}

}

func TestCalculateControlDigit(t *testing.T) {
	tests := []struct {
		taxId     string
		wantDigit string
	}{
		{"RSSMRA80A01H501", "U"},
		{"NCLNDR85M23L565", "B"},
		{"BNCLGU80A01A757", "T"},
		{"VRDMRA80A41L750", "I"},
	}
	for _, tt := range tests {
		t.Run(tt.taxId, func(t *testing.T) {
			gotDigit := calculateControlDigit(tt.taxId)
			if gotDigit != tt.wantDigit {
				t.Errorf("calculateControlDigit() gotDigit = %v, want %v", gotDigit, tt.wantDigit)
			}
		})
	}
}

func TestExtractInfo(t *testing.T) {
	tests := []struct {
		taxId      string
		wantedInfo *Info
	}{
		{"RSSMRA85M23L565B", &Info{Sex: "M", BirthDate: []string{"1985-08-23"}, BirthPlace: "L565"}},
		{"RSSMRA80A01A757T", &Info{Sex: "M", BirthDate: []string{"1980-01-01"}, BirthPlace: "A757"}},
		{"RSSMRA18H24H501A", &Info{Sex: "M", BirthDate: []string{"1918-06-24", "2018-06-24"}, BirthPlace: "H501"}},
		{"RSSMRA18H64H501A", &Info{Sex: "F", BirthDate: []string{"1918-06-24", "2018-06-24"}, BirthPlace: "H501"}},
	}
	for _, tt := range tests {
		t.Run(tt.taxId, func(t *testing.T) {
			gotInfo, err := ExtractInfo(tt.taxId)
			if err != nil {
				t.Errorf("ExtractInfo() error = %v", err)
			}
			if gotInfo == nil {
				t.Errorf("ExtractInfo() gotInfo = %v, want %v", gotInfo, tt.wantedInfo)
			}

			if !gotInfo.Equals(tt.wantedInfo) {
				t.Errorf("ExtractInfo() gotInfo = %v, want %v", gotInfo, tt.wantedInfo)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	tests := []struct {
		lastname  string
		name      string
		sex       string
		birthDate string
		taxId     string
		wantErr   bool
	}{
		{"Rossi", "Mario", "M", "1980-01-01", "RSSMRA80A01A061H", false},
		{"Rossi", "Mario", "F", "1980-01-01", "RSSMRA80A01A061H", true},
		{"Rossi", "Mario", "M", "1980-04-01", "RSSMRA80A01A061H", true},
		{"Rossi", "Mario", "M", "1900-01-01", "BNCGAI00A41H501H", true},
		{"Bianchi", "Gaia", "F", "1900-01-01", "BNCGAI00A41H501H", true},
		{"Bianchi", "Gaia", "F", "2000-01-01", "BNCGAI00A41H501H", false},
		{"Bianchi", "Gaia", "F", "2100-01-01", "BNCGAI00A41H501H", true},
		{"Verdi", "Lucia", "F", "1905-01-01", "VRDLCU05A41L736X", false},
		{"Verdi", "Lucia", "F", "2005-01-01", "VRDLCU05A41L736X", false},
		{"Verdi", "Lucia", "F", "1905-01-01", "VRDLCU05A41L736X", false},
		{"Verdi", "Lucia", "F", "2005-01-01", "VRDLCU05A41L736X", false},
		{"Verdi", "Gianni", "M", "1904-01-01", "VRDGNN04A01L219N", true},
		{"Verdi", "Gianni", "M", "2004-01-01", "VRDGNN04A01L219N", false},
	}
	for _, tt := range tests {
		t.Run(tt.taxId, func(t *testing.T) {
			err := Verify(tt.lastname, tt.name, tt.sex, tt.birthDate, tt.taxId)
			if (err != nil) && !tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
