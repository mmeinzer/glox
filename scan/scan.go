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
	switch s.advance() {
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
	case "!":
		if s.match("=") {
			s.addToken(bangEqual)
		} else {
			s.addToken(bang)
		}
	case "=":
		// token: ==
		if s.match("=") {
			s.addToken(equalEqual)
		} else {
			s.addToken(equal)
		}
	case "<":
		if s.match("=") {
			// token: <=
			s.addToken(lessEqual)
		} else {
			s.addToken(less)
		}
	case ">":
		// token: >=
		if s.match("=") {
			s.addToken(greateEqual)
		} else {
			s.addToken(greater)
		}
	case "/":
		// token: // (start of a comment)
		if s.match("/") {
			for s.peek() != "\n" && s.isAtEnd() {
				// Drop everything until the end of the line
				s.advance()
			}
		} else {
			s.addToken(slash)
		}
	case " ":
	case "\r":
	case "\t":
		// Ignore whitepsace
		break
	case "\n":
		s.line++
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

// match determines if the current character matches match and consumes it if so
func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}

	if s.characters[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "" // TODO: This is \0 in CI but that gives a syntax error in Go
	}

	return s.characters[s.current]
}
