package parser

import (
	"fmt"
	"github.com/mark-marushak/number-parser/logger"
)

type Service interface {
	Parse(string) []int
}

type parserService struct {
	repository Repository
	logger     logger.Logger
}

func NewParser(repo Repository, logger logger.Logger) Service {
	return &parserService{repository: repo, logger: logger}
}

func (p parserService) Parse(s string) []int {
	output, err := p.repository.Parse(s)
	if err != nil {
		p.logger.Error(fmt.Errorf("Error while parsing string: %v", err))
		return nil
	}

	return output
}
