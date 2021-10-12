package scan

import "github.com/mmeinzer/glox/report"

type Scanner struct {
	source        string
	characters    []string
	tokens        []token
	start         int
	current       int
	line          int
	errorReporter report.ErrorReporter
}

// NewScanner creates a Lox scanner from a source code string
func NewScanner(source string, reporter report.ErrorReporter) *Scanner {
	p := Scanner{source: source, characters: []string{}, tokens: []token{}, start: 0, current: 0, line: 1}
	p.sourceToChars()
	return &p
}

// ScanTokens scans the source string and returns a slice of tokens
func (s *Scanner) ScanTokens() []token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.addToken(eof)
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case "(":
		s.addToken(leftParen)
	case ")":
		s.addToken(rightParen)
	case "{":
		s.addToken(leftBrace)
	case "}":
		s.addToken(rightBrace)
	case ",":
		s.addToken(comma)
	case ".":
		s.addToken(dot)
	case "+":
		s.addToken(plus)
	case ";":
		s.addToken(semicolon)
	case "*":
		s.addToken(star)
	default:
		s.errorReporter.Error(s.line, "Unexpected character.")
	}
}

// advance gets the character pointed at by current, increments current, and returns the character
func (s *Scanner) advance() string {
	char := s.characters[s.current]
	s.current++
	return char
}

func (s *Scanner) addToken(t tokenType) {
	var bytes []byte
	for _, char := range s.characters[s.start:s.current] {
		bytes = append(bytes, []byte(char)...)
	}

	s.tokens = append(s.tokens, token{tType: t, lexeme: string(bytes), literal: "", line: s.line})
}

func (s *Scanner) sourceToChars() {
	for _, k := range s.source {
		s.characters = append(s.characters, string(k))
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.characters)
}
