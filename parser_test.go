package code_interview

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := Parse(test.data)
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Parse(test.data)
			if err == nil {
				t.Fail()
			}
		})
	}

}

func TestUpdatedParse(t *testing.T) {
	s := "1;4-7;8;9;10-13;1-3;22"
	//s := "1;2;3;4;5"

	output := make([]int, 0, len(s))
	splitted := strings.Split(s, "-")
	for i := 0; i < len(splitted); i++ {
		numbers := strings.Split(splitted[i], ";")

		if i > 0 {
			digit, _ := strconv.Atoi(numbers[0])
			for k := output[len(output)-1] + 1; k <= digit; k++ {
				output = append(output, k)
			}
			numbers = numbers[1:]
		}

		for j := 0; j < len(numbers); j++ {
			digit, _ := strconv.Atoi(numbers[j])
			output = append(output, digit)
		}
	}

	fmt.Println(output)
}
