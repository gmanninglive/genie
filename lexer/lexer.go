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
			l.Emit(TokenText)
			return lexNewLine
		}
		if l.NextHasPrefix(OpenBlock) {
			l.Emit(TokenText)
			return lexOpenBlock
		}
		if l.NextHasPrefix(CloseBlock) {
			l.Emit(TokenText)
			return lexCloseBlock
		}
		if l.Next() == golex.EOF {
			break
		}
	}
	if l.Current > l.Start {
		l.Emit(TokenText)
	}

	l.Emit(golex.TokenEOF)
	return nil
}

func lexOpenBlock(l *golex.Lexer) golex.StateFn {
	l.Current += len(OpenBlock)
	l.Emit(TokenOpenBlock)

	if golex.IsSpace(l.Next()) {
		l.Ignore()
	}

	if l.NextHasPrefix(".") {
		return lexDotIndentifier
	}

	return lexIdentifier
}

func lexDotIndentifier(l *golex.Lexer) golex.StateFn {
	for {
		if l.NextHasPrefix(CloseBlock) {
			if l.Current > l.Start {
				l.Emit(TokenDotIdentifier)
			}
			return lexCloseBlock
		}

		switch r := l.Next(); {
		case r == EOF || r == '\n':
			l.Errorf("Error Lexing a variable in a {{ block }}, No closing block found")
		case golex.IsSpace(r):
			l.Backup()
			if l.Current > l.Start {
				l.Emit(TokenDotIdentifier)
			}
			l.Next()
			l.Ignore()
			return lexIdentifier
		}
	}
}

func lexIdentifier(l *golex.Lexer) golex.StateFn {
	for {
		if l.NextHasPrefix(CloseBlock) {
			if l.Current > l.Start {
				l.Emit(TokenIdentifier)
			}
			return lexCloseBlock
		}

		switch r := l.Next(); {
		case r == EOF || r == '\n':
			l.Errorf("Error Lexing a variable in a {{ block }}, No closing block found")
		case golex.IsSpace(r):
			l.Backup()
			if l.Current > l.Start {
				l.Emit(TokenIdentifier)
			}
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
