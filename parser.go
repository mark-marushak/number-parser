package code_interview

import (
	"strconv"
	"strings"
)

func Parse(s string) (output []int, err error) {
	splitted := strings.Split(s, "-")
	var digit int
	for i := 0; i < len(splitted); i++ {
		numbers := strings.Split(splitted[i], ";")

		if i > 0 {
			digit, err = strconv.Atoi(numbers[0])
			if err != nil {
				return nil, err
			}
			for k := output[len(output)-1] + 1; k <= digit; k++ {
				output = append(output, k)
			}
			numbers = numbers[1:]
		}

		for j := 0; j < len(numbers); j++ {
			digit, err = strconv.Atoi(numbers[j])
			if err != nil {
				return nil, err
			}
			output = append(output, digit)
		}
	}

	return
}
