package gherkin

import (
	"sync"

	"github.com/l3pp4rd/go-behat/gherkin/lexer"
)

type item struct {
	next, prev *item
	value      *lexer.Token
}

type AST struct {
	head, tail *item
	mut        *sync.Mutex
}

func newAST() *AST {
	return &AST{mut: &sync.Mutex{}}
}

func (l *AST) addTail(t *lexer.Token) *item {
	l.mut.Lock()
	defer l.mut.Unlock()

	it := &item{next: nil, prev: l.tail, value: t}
	if l.head == nil {
		l.head = it
	} else {
		l.tail.next = it
	}
	l.tail = it
	return l.tail
}

func (l *AST) addBefore(t *lexer.Token, i *item) *item {
	l.mut.Lock()
	defer l.mut.Unlock()

	it := &item{next: i, prev: i.prev, value: t}
	i.prev = it
	if it.prev == nil {
		l.head = it
	}
	return it
}

func (l *AST) addAfter(t *lexer.Token, i *item) *item {
	l.mut.Lock()
	defer l.mut.Unlock()

	it := &item{next: i.next, prev: i, value: t}
	i.next = it
	if it.next == nil {
		l.tail = it
	}
	return it
}