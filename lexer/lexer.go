package lexer

import "fmt"

// represents a token returned from the lexer
type token struct {
	typ tokenType
	val string 
}

// tokenType represents the type of tokens
type tokenType int

const (
		tokenError tokenType = iota // value is text of error

		tokenCloseBlock						// end of block, }}					
		tokenDot										// the cursor, '.'
		tokenEOF
		tokenElse									// else keyword
		tokenEnd										// end keyword
		tokenHelper								// helper method
		tokenIf										// if keyword
		tokenNumber								// number
		tokenOpenBlock							// start of block, {{
		tokenRawString 						// raw quoted string
		tokenString								// quoted string
		tokenText									// plain text
		)

func (t token) String() string {
	switch t.typ {
	case tokenEOF:
		return "EOF"
	case tokenError:
		return t.val
	}
	if len(t.val) > 10 {
		return fmt.Sprintf("%.10q...", t.val)
	}
	return fmt.Sprintf("%q", t.val)
}
