package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseWithNormalData(t *testing.T) {
	tests := []struct {
		name   string
		data   string
		expect []int
	}{
		{
			"test range",
			"1;4-7;9",
			[]int{1, 4, 5, 6, 7, 9},
		},
		{
			"with one number",
			"1",
			[]int{1},
		},
		{
			"empty",
			"",
			[]int{},
		},
	}

	repository := parserRepository{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := repository.Parse(test.data)
			if err != nil {
				return
			}
			assert.Equal(t, test.expect, output)
		})
	}

}

func TestParseWithBadCases(t *testing.T) {
	tests := []struct {
		name string
		data string
	}{
		{
			"test double range",
			"1;4--7;9",
		},
		{
			"wrong separator",
			"11,11",
		},
	}

	repository := parserRepository{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := repository.Parse(test.data)
			if err == nil {
				t.Fail()
			}
		})
	}

}
