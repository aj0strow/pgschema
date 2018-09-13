package sql

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type itemType int

const (
	itemError itemType = iota
	itemEOF
	itemToken
	itemSpecial
)

const eof = -1

type item struct {
	typ itemType
	val string
}

type lexer struct {
	input string
	start int
	pos   int
	width int
	items chan item
}

func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.ignore()
}

func (l *lexer) run(state stateFn) {
	for state != nil {
		state = state(l)
	}
	close(l.items)
}

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{
		itemError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

func nextRune(l *lexer, rs ...rune) bool {
	r := l.next()
	for i := 0; i < len(rs); i++ {
		if r == rs[i] {
			return true
		}
	}
	l.backup()
	return false
}

func nextString(l *lexer, s string) bool {
	if strings.HasPrefix(l.input[l.pos:], s) {
		l.pos += len(s)
		return true
	}
	return false
}

func nextKeyword(l *lexer, s string) bool {
	return nextString(l, strings.ToLower(s)) || nextString(l, strings.ToUpper(s))
}

type stateFn func(*lexer) stateFn

func lex(input string) *lexer {
	l := &lexer{
		input: input,
		items: make(chan item),
	}
	go l.run(lexSQL)
	return l
}

func lexSQL(l *lexer) stateFn {
	lexSpace(l)
	if l.peek() == eof {
		l.emit(itemEOF)
		return nil
	}
	if nextString(l, "--") {
		return lexComment
	}
	if nextString(l, "/*") {
		return lexCommentBlock
	}
	if nextRune(l, ';', '(', ')', '[', ']', ',', '*') {
		l.emit(itemSpecial)
		return lexSQL
	}
	if r := l.peek(); unicode.IsLetter(r) || r == '_' {
		return lexToken
	}
	return nil
}

func lexSpace(l *lexer) {
	for unicode.IsSpace(l.next()) {
		l.ignore()
	}
	l.backup()
}

func lexComment(l *lexer) stateFn {
	for {
		r := l.next()
		if r == '\n' || r == '\r' || r == eof {
			break
		}
	}
	l.backup()
	l.ignore()
	return lexSQL
}

func lexCommentBlock(l *lexer) stateFn {
	n := 1
	for l.next() != eof {
		if nextString(l, "/*") {
			n += 1
		} else if nextString(l, "*/") {
			n -= 1
		} else if n == 0 {
			break
		}
	}
	l.backup()
	l.ignore()
	return lexSQL
}

func lexToken(l *lexer) stateFn {
	for r := l.next(); unicode.IsLetter(r) || r == '_' || unicode.IsDigit(r); r = l.next() {
		// continue
	}
	l.backup()
	l.emit(itemToken)
	return lexSQL
}
