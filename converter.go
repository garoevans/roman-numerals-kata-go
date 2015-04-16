package converter

type RomanNumeralConverter struct {
}

type RomanNumeralConvertion struct {
    number int
    numeral string
}

var mapping = []RomanNumeralConvertion{
    RomanNumeralConvertion{1000, "M"},
    RomanNumeralConvertion{900, "CM"},
    RomanNumeralConvertion{500, "D"},
    RomanNumeralConvertion{400, "CD"},
    RomanNumeralConvertion{100, "C"},
    RomanNumeralConvertion{90, "XC"},
    RomanNumeralConvertion{50, "L"},
    RomanNumeralConvertion{40, "XL"},
    RomanNumeralConvertion{10, "X"},
    RomanNumeralConvertion{9, "IX"},
    RomanNumeralConvertion{5, "V"},
    RomanNumeralConvertion{4, "IV"},
    RomanNumeralConvertion{1, "I"},
}

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
