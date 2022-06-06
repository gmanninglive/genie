package main

import (
	"fmt"
	"genie/helpers"
	"genie/lexer"
	"reflect"
	"strings"

	"github.com/gmanninglive/golex"
)

type Token = golex.Token

type Parser struct {
	Queue   *Pqueue
	out     strings.Builder
	vars    TplVars
	Helpers map[string]interface{}
}

// Public method for initialising and running the parser
func (p Parser) Parse(template string, vars TplVars) string {
	p.Helpers = helpers.Init()
	p.vars = vars

	l := lexer.Lex(template)

	p.run(l)

	return p.out.String()
}

// Runs the parser
// Initialises a queue
// Runs the parser state machine concurrently
// Listening for tokens from the lexer
func (p *Parser) run(l *golex.Lexer) {
	p.Queue = p.Queue.Init()
	go p.stateMachine()
	for {
		token, done := l.Listen()
		if done {
			p.Queue.finished <- true
			return
		}
		//fmt.Printf("Token: %s, Type: %s\n", token.Val, token.Typ)
		p.Queue.push(token)
	}
}

// Parser statemachine
// Receives tokens from the queue and performs actions
// Until either EOf token received or queue finished message received
func (p *Parser) stateMachine() {
	for {
		select {
		case tok := <-p.Queue.tokens:
			switch tok.Typ {
			case lexer.TokenText:
				p.parseText(tok)
			case lexer.TokenNewLine:
				p.parseNewLine(tok)
			case lexer.TokenOpenBlock:
				// skip
			case lexer.TokenCloseBlock:
				// skip
			case lexer.TokenDotIdentifier:
				p.parseDotIdentifier(tok)
			case lexer.TokenIdentifier:
				p.parseIdentifier(tok)
			case golex.TokenEOF:
				return
			}
		case <-p.Queue.finished:
			return
		}
	}
}

// Sends the token value to the output string builder
func (p *Parser) parseText(tok Token) {
	p.out.Grow(len(tok.Val))
	fmt.Fprintf(&p.out, tok.Val)
}

func (p *Parser) parseNewLine(tok Token) {
	p.out.Grow(len("\n"))
	fmt.Fprintf(&p.out, tok.Val)
}

// Parse Dot Identifier aka Helper function
// If specified helper is not valid parser will skip to next state
// If args for helper are not valid parser will skip to next state
func (p *Parser) parseDotIdentifier(tok Token) {
	helper, valid := p.Helpers[tok.Val[1:]]
	if !valid {
		fmt.Printf("Helper Error: %s Is not a valid helper function\n", tok.Val[1:])
		return
	}

	helperType := reflect.TypeOf(helper)
	var helperArgs []reflect.Kind

	for i := 0; i < helperType.NumIn(); i++ {
		helperArgs = append(helperArgs, helperType.In(i).Kind())
	}

	// Array of args, taken from the next tokens enqueued by the lexer
	// Checks if arg type matches type required by helper function
	// Returns early if invalid type
	args := make([]reflect.Value, len(helperArgs))
	for i := range args {
		token := <-p.Queue.tokens
		args[i] = reflect.ValueOf(p.vars[token.Val])
		if args[i].Kind() != helperArgs[i] {
			fmt.Printf("Helper Error: %s Is not a valid argument for helper %s. Expected type: %s\n",
				args[i], tok.Val[1:], helperArgs[i])
			return
		}
	}

	// Call function and convert to string,
	// Currently all helpers should return a string
	fnCall := reflect.ValueOf(helper).Call(args)
	res := fnCall[0].Interface().(string)

	p.out.Grow(len(res))
	fmt.Fprintf(&p.out, res)
}

// Replaces Indentifier with the ctx variable
func (p *Parser) parseIdentifier(tok Token) {
	if val, isIn := p.vars[tok.Val]; isIn {
		p.out.Grow(len(val))
		fmt.Fprintf(&p.out, val)
	}
}

// Parser queue
// Containing a channel for tokens,
// And a channel to notify when finished
type Pqueue struct {
	tokens   chan Token
	finished chan bool
}

// Initialise parser queue, returns queue pointer
func (q *Pqueue) Init() *Pqueue {
	return &Pqueue{
		tokens:   make(chan Token),
		finished: make(chan bool),
	}
}

// Add job to back of queue
func (q *Pqueue) push(t Token) {
	q.tokens <- t
}

// Pop from front of queue
// Not really needed right now... may delete
func (q *Pqueue) Pop() {
	<-q.tokens
}
