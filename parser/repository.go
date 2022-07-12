package parser

type Repository interface {
	Parse(string) ([]int, error)
}
