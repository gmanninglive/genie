package main

import (
	"fmt"
	"genie/lexer"
	"os"
	"strings"

	"github.com/gmanninglive/golex"
)

type Token = golex.Token

type Parser struct {
	Queue *Pqueue
	out   strings.Builder
}

// Public method for initialising and running the parser
func (p Parser) Parse(templatePath string, vars TplVars) string {
	file, err := os.ReadFile(templatePath)
	if err != nil {
		panic(err)
	}

	fileStr := string(file)

	l := lexer.Lex(fileStr)

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
				// ignore
			case lexer.TokenCloseBlock:
				// ignore
			case lexer.TokenIdentifier:
				p.parseIdentifier(tok)
			case lexer.TokenVariable:
				p.parseVariable(tok)
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

func (p *Parser) parseIdentifier(tok Token) {
	p.out.Grow(len(tok.Val))
	fmt.Fprintf(&p.out, tok.Val)
}

func (p *Parser) parseVariable(tok Token) {
	p.out.Grow(len(tok.Val))
	fmt.Fprintf(&p.out, tok.Val)
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
		tokens:   make(chan Token, 2),
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
