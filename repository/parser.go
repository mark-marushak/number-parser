package repository

import (
	"github.com/mark-marushak/number-parser/parser"
	"strconv"
	"strings"
)

type parserRepository struct{}

func NewParser() parser.Repository {
	return &parserRepository{}
}

func (p parserRepository) Parse(s string) (output []int, err error) {
	numbers := strings.Split(s, ";")
	for i := 0; i < len(numbers); i++ {
		number, err := strconv.Atoi(numbers[i])
		if err != nil {
			rangeStrings := strings.Split(numbers[i], "-")
			rangeNumbers := make([]int, 0, 2)
			for k := 0; k < len(rangeStrings); k++ {
				number, err = strconv.Atoi(rangeStrings[k])
				if err != nil {
					panic(err)
				}
				rangeNumbers = append(rangeNumbers, number)
			}

			for j := rangeNumbers[0]; j <= rangeNumbers[1]; j++ {
				output = append(output, j)
			}

			continue
		}

		output = append(output, number)
	}

	return
}
