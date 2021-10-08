package scan

type tokenType uint8

const (
	// Single-character tokens
	leftParen tokenType = iota
	rightParen
	leftBrace
	rightBrace
	comma
	dot
	minus
	plus
	semicolon
	slash
	star

	// One or two character tokens
	bang
	bangEqual
	equal
	equalEqual
	greater
	greateEqual
	less
	lessEqual

	// Literals
	identifier
	str
	number

	// Keywords
	// collisions with Go are avoided by appending G (for Glox)
	and
	class
	elseG
	falseG
	fun
	forG
	ifG
	nilG
	or
	print
	returnG
	super
	this
	trueG
	varG
	while

	eof
)

type token struct {
	tType   tokenType
	lexeme  string
	literal string // TODO: this is Object in CI?
	line    int
}
