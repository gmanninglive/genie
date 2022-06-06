package lexer

import (
	"github.com/gmanninglive/golex"
)

const (
	TokenText golex.TokenType = iota
	TokenNewLine
	TokenOpenBlock
	TokenCloseBlock
	TokenDotIdentifier // An Identifier starting with . aka helper function
	TokenIdentifier    // An Identifier to be parsed with context / helper
)

const OpenBlock = "{{"
const CloseBlock = "}}"
const NewLine = "\n"
const EOF = golex.EOF

func baseStateFn(l *golex.Lexer) golex.StateFn {
	for {
		if l.NextHasPrefix(NewLine) {
			l.CheckEmit(TokenText)
			return lexNewLine
		}
		if l.NextHasPrefix(OpenBlock) {
			l.CheckEmit(TokenText)
			return lexOpenBlock
		}
		if l.NextHasPrefix(CloseBlock) {
			l.CheckEmit(TokenText)
			return lexCloseBlock
		}
		if l.Next() == golex.EOF {
			break
		}
	}

	l.CheckEmit(TokenText)

	l.Emit(golex.TokenEOF)
	return nil
}

func lexOpenBlock(l *golex.Lexer) golex.StateFn {
	l.Current += len(OpenBlock)
	l.Emit(TokenOpenBlock)
	for {
		if l.NextHasPrefix(".") {
			return lexDotIndentifier
		}
		switch r := l.Next(); {
		case golex.IsSpace(r):
			l.Ignore()
		case golex.IsAlpha(r):
			return lexIdentifier
		}
	}
}

func lexDotIndentifier(l *golex.Lexer) golex.StateFn {
	for {
		if l.NextHasPrefix(CloseBlock) {
			l.CheckEmit(TokenDotIdentifier)
			return lexCloseBlock
		}

		switch r := l.Next(); {
		case r == EOF || r == '\n':
			l.Errorf("Error Lexing a variable in a {{ block }}, No closing block found")
		case golex.IsSpace(r):
			l.Backup()
			l.CheckEmit(TokenDotIdentifier)

			l.Next()
			l.Ignore()
			return lexIdentifier
		}
	}
}

func lexIdentifier(l *golex.Lexer) golex.StateFn {
	for {
		if l.NextHasPrefix(CloseBlock) {
			l.CheckEmit(TokenIdentifier)

			return lexCloseBlock
		}

		switch r := l.Next(); {
		case r == EOF || r == '\n':
			l.Errorf("Error Lexing a variable in a {{ block }}, No closing block found")
		case golex.IsSpace(r):
			l.Backup()
			l.CheckEmit(TokenIdentifier)

			l.Next()
			l.Ignore()
			return lexIdentifier
		}
	}
}

func lexCloseBlock(l *golex.Lexer) golex.StateFn {
	l.Current += len(CloseBlock)
	l.Emit(TokenCloseBlock)
	return baseStateFn
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
