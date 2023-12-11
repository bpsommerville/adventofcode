package main

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestDay01ParseLine(t *testing.T) {
	// always start tests with this
	is := is.New(t)

	tests := []struct {
		name  string
		input string
		total uint
		err   error
	}{
		{
			name:  "EmptyInputGivesZero",
			input: "",
			total: 0,
		},
		{
			name:  "OneDigitReusesSameDigit",
			input: "1",
			total: 11,
		},
		{
			name:  "TwoDigitsGivesValue",
			input: "12",
			total: 12,
		},
		{
			name:  "MultipleDigitsUsesFirstLast",
			input: "9984324",
			total: 94,
		},
		{
			name:  "IgnoreNonDigits",
			input: "ase110i942afd7asd",
			total: 17,
		},
		{
			name:  "SpellDigitsWithLetters",
			input: "onetwothree",
			total: 13,
		},
		{
			name:  "OverlappingWordsTakesFirst",
			input: "twonethreeight",
			total: 28,
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotTotal, gotErr := d01Parseline(tt.input)

			if tt.err != nil {
				is.Equal(tt.err.Error(), gotErr.Error())
				is.Equal(0, gotTotal)
			} else {
				is.NoErr(gotErr)
				is.Equal(tt.total, gotTotal)
			}
		})
	}
}

func TestDay01(t *testing.T) {
	// always start tests with this
	is := is.New(t)

	tests := []struct {
		name  string
		input []string
		total uint
		err   error
	}{
		{
			name:  "Example",
			input: []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
			total: 142,
		},
		{
			name:  "Example 2",
			input: strings.Split("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen", "\n"),
			total: 281,
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotTotal, gotErr := Day01(tt.input)

			if tt.err != nil {
				is.Equal(tt.err.Error(), gotErr.Error())
				is.Equal(0, gotTotal)
			} else {
				is.NoErr(gotErr)
				is.Equal(tt.total, gotTotal)
			}
		})
	}
}
