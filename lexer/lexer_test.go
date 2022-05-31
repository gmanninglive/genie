package lexer

import (
	"os"
	"testing"
)

const testString = "<div>{{name}}</div>"

func TestLex(t *testing.T) {

  l := New("test", testString)
	var received []Token
  for l.State != nil {
    received = append(received, l.NextToken())
  }

	t.Log("length", len(received))

	if len(received) != 5 {
		t.Errorf("Expected 5 tokens, got %d", len(received))
	}

	t.Run("With hbs template", func (t *testing.T) {
		f, err := os.ReadFile("../examples/example_template.hbs")
  	if err != nil {
			panic(err)
		}

		l := New("hbs", string(f))
		var received []Token
		for {
			select {
			case tok := <-l.Tokens:
			l.Run()
			received = append(received, tok)
			t.Log("Token", tok.Val)
			}
		}

		t.Log("length", len(received))

		if len(received) != 9 {
			t.Errorf("Expected 9 tokens, got %d", len(received))
		}
	})
}