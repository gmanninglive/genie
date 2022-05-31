package lexer

import (
	"github.com/gmanninglive/golex"
)

const (
	TokenText golex.TokenType = iota
	TokenNewLine
	TokenOpenBlock
	TokenCloseBlock
	TokenIdentifier
	TokenVariable
)

const OpenBlock = "{{"
const CloseBlock = "}}"
const NewLine = "\n"

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
	return baseStateFn
}

func lexIndentifier(l *golex.Lexer) golex.StateFn {
	for {
		if golex.IsSpace(l.Next()) {
			l.Backup()
			l.Emit(TokenIdentifier)
			l.Ignore()
			return lexVariable
		}
		if l.NextHasPrefix(CloseBlock) {
			l.Emit(TokenIdentifier)
			return lexCloseBlock
		}
		if l.NextHasPrefix(NewLine) {
			l.Errorf("Incorrect formatting! Received New Line within a {{ block }}")
		}
		l.Errorf("Error Lexing a {{ block }}, No closing block found")
	}
}

func lexVariable(l *golex.Lexer) golex.StateFn {
	for {
		if golex.IsSpace(l.Next()) {
			l.Backup()
			l.Emit(TokenVariable)
			l.Ignore()
			return lexVariable
		}
		if l.NextHasPrefix(CloseBlock) {
			l.Emit(TokenVariable)
			return lexCloseBlock
		}
		if l.NextHasPrefix(NewLine) {
			l.Errorf("Incorrect formatting! Received New Line within a {{ block }}")
		}
		l.Errorf("Error Lexing a variable in a {{ block }}, No closing block found")
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
