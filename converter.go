package converter

// RomanNumeralConverter is the public strcut for the roman numeral convertion
type RomanNumeralConverter struct {
}

type romanNumeralConvertion struct {
    number int
    numeral string
}

var mapping = []romanNumeralConvertion{
    romanNumeralConvertion{1000, "M"},
    romanNumeralConvertion{900, "CM"},
    romanNumeralConvertion{500, "D"},
    romanNumeralConvertion{400, "CD"},
    romanNumeralConvertion{100, "C"},
    romanNumeralConvertion{90, "XC"},
    romanNumeralConvertion{50, "L"},
    romanNumeralConvertion{40, "XL"},
    romanNumeralConvertion{10, "X"},
    romanNumeralConvertion{9, "IX"},
    romanNumeralConvertion{5, "V"},
    romanNumeralConvertion{4, "IV"},
    romanNumeralConvertion{1, "I"},
}

// Convert will convert any number to it's roman numeral equivalent
func (r *RomanNumeralConverter) Convert(number int) string {
    var numeral string;

    for _, convertion := range mapping {
        for (number - convertion.number) >= 0 {
            number -= convertion.number
            numeral += convertion.numeral
        }
    }

    return numeral;
}
