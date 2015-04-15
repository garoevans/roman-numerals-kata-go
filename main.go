package main

type RomanNumeralConverter struct {
}

type RomanNumeralConvertion struct {
    Number int
    Numeral string
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
    var numeralString string;

    for _, convertion := range mapping {
        for (number - convertion.Number) >= 0 {
            number -= convertion.Number
            numeralString += convertion.Numeral
        }
    }

    return numeralString;
}
