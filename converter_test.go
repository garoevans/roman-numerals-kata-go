package converter

import "testing"

func TestNumbersConvertToCorrectRomanNumerals(t *testing.T) {
    cases := []struct {
        number int
        numeral string
    }{
        {1, "I"},
        {2, "II"},
        {4, "IV"},
        {5, "V"},
        {49, "XLIX"},
        {872, "DCCCLXXII"},
        {2015, "MMXV"},
        {1999, "MCMXCIX"},
        {1738, "MDCCXXXVIII"},
        {0, ""},
    }

    r := RomanNumeralConverter{}
    for _, c := range cases {
        got := r.Convert(c.number)
        if got != c.numeral {
            t.Errorf("Convert(%d) == %q, want %q", c.number, got, c.numeral)
        }
    }
}
