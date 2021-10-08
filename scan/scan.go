package scan

type scanner struct {
	source     string
	characters []string
	tokens     []token
	start      int
	current    int
	line       int
}

func NewScanner(source string) *scanner {
	p := scanner{source: source, characters: []string{}, tokens: []token{}, start: 0, current: 0, line: 1}
	p.sourceToChars()
	return &p
}

func (s *scanner) ScanTokens() []token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.addToken(eof)
	return s.tokens
}

func (s *scanner) scanToken() {
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
	}
}

// advance gets the character pointed at by current, increments current, and returns the character
func (s *scanner) advance() string {
	char := s.characters[s.current]
	s.current++
	return char
}

func (s *scanner) addToken(t tokenType) {
	var bytes []byte
	for _, char := range s.characters[s.start:s.current] {
		bytes = append(bytes, []byte(char)...)
	}

	s.tokens = append(s.tokens, token{tType: t, lexeme: string(bytes), literal: "", line: s.line})
}

func (s *scanner) sourceToChars() {
	for _, k := range s.source {
		s.characters = append(s.characters, string(k))
	}
}

func (s *scanner) isAtEnd() bool {
	return s.current >= len(s.characters)
}
