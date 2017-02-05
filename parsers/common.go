package parsers

type Parser interface {
	Parse(string) map[string]string
}

