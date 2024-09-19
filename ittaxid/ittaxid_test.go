package ittaxid

import "testing"

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
	}{
		{"RSSMRA85M01H501Z", 1, true},
		{"rssmra85M23H501B", 23, true},
		{"RSSMRA85M41H501Z", 1, false},
		{"rssmra85M63H501B", 23, false},
		{"RSSMRA85M42H501Z", 2, false},
	}

	for _, tt := range tests {
		t.Run(tt.taxId, func(t *testing.T) {
			gotDay, gotIsMale := extractDayOfBirthAndSex(tt.taxId)
			if gotDay != tt.wantDay {
				t.Errorf("extractDayOfBirthAndSex() gotDay = %v, want %v", gotDay, tt.wantDay)
			}
			if gotIsMale != tt.wantIsMale {
				t.Errorf("extractDayOfBirthAndSex() gotIsMale = %v, want %v", gotIsMale, tt.wantIsMale)
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
