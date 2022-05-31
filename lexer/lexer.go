package lexer

import (
	"github.com/gmanninglive/golex"
)

const (
	TokenText golex.TokenType = iota
	TokenNewLine
)

const NewLine = "\n"

func baseStateFn(l *golex.Lexer) golex.StateFn {
	for {
		if l.NextHasPrefix(NewLine) {
			l.Emit(TokenText)
			return lexNewLine
		}
		if l.Next() == golex.EOF {
			break
		}
	}
	l.Emit(TokenText)

	l.Emit(golex.TokenEOF)

	return nil
}

func lexNewLine(l *golex.Lexer) golex.StateFn {
	l.Current += len(NewLine)
	l.Emit(TokenNewLine)
	return baseStateFn
}

func Lex(s string) *golex.Lexer {
	l := golex.New("hdb-lexer", string(s), baseStateFn)
	l.RunConc()

	return l
}
