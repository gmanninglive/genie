package lexer

import (
	"testing"

	"github.com/gmanninglive/golex"
)

func TestLex(t *testing.T) {
	testStrings := []string{
		"<div>{{.dotIdentifier Name}}</div>",
		"<div>{{ .dotIdentifier Name}}</div>",
		"<div>{{ .dotIdentifier Name}}</div>",
		"<div>{{.dotIdentifier Name}}</div>",
	}
	for i := range testStrings {
		l := Lex(testStrings[i])

		var received []golex.Token
		for {
			token, done := l.Listen()
			if done {
				break
			}
			t.Logf("Token: %s\n", token.Val)
			received = append(received, token)
		}

		if len(received) != 6 {
			t.Errorf("Expected 6 tokens received: %o\n", len(received))
		}
		if received[0].Typ != TokenText {
			t.Errorf("run: %o. Expected %o, Received: %o\n", i, TokenText, received[0].Typ)
		}
		if received[1].Typ != TokenOpenBlock {
			t.Errorf("run: %o. Expected %o, Received: %o\n", i, TokenOpenBlock, received[1].Typ)
		}
		if received[2].Typ != TokenDotIdentifier {
			t.Errorf("run: %o. Expected %o, Received: %o\n", i, TokenDotIdentifier, received[2].Typ)
		}
		if received[3].Typ != TokenIdentifier {
			t.Errorf("run: %o. Expected %o, Received: %o\n", i, TokenIdentifier, received[3].Typ)
		}
		if received[4].Typ != TokenCloseBlock {
			t.Errorf("run: %o. Expected %o, Received: %o\n", i, TokenCloseBlock, received[4].Typ)
		}
		if received[5].Typ != TokenText {
			t.Errorf("run: %o. Expected %o, Received: %o\n", i, TokenText, received[5].Typ)
		}
	}
}
