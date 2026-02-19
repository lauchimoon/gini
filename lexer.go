package gini

import (
    "strings"
    "unicode"
)

const (
    tokenComment = iota
    tokenOpenBracket
    tokenCloseBracket
    tokenSymbol
    tokenEqual
    tokenString
)

type lexer struct {
    Source    string
    LenSource int
    Cursor    int
}

type token struct {
    id    int
    value string
}

func (l *lexer) lex() []token {
    tokens := []token{}
    valueBuilder := strings.Builder{}
    for l.Cursor < l.LenSource {
        c := l.consume()
        if c == ';' {
            tokens = append(tokens, token{
                id: tokenComment,
                value: ";",
            })

            c = l.consume()
            for c != '\n' {
                valueBuilder.WriteRune(c)
                c = l.consume()
            }
            tokens = append(tokens, token{
                id: tokenComment,
                value: valueBuilder.String(),
            })
        }

        if unicode.IsLetter(c) || unicode.IsDigit(c) {
            valueBuilder.WriteRune(c)
            c = l.consume()
            for c != '\n' &&
                (unicode.IsLetter(c) || unicode.IsDigit(c) ||
                 unicode.IsSpace(c) || (unicode.IsPunct(c) && c != '"' && c != ']')) {
                valueBuilder.WriteRune(c)
                c = l.consume()
            }
            tokens = append(tokens, token{
                id: tokenSymbol,
                value: strings.TrimSpace(valueBuilder.String()),
            })
        }

        if c == '[' {
            tokens = append(tokens, token{
                id: tokenOpenBracket,
                value: "[",
            })
        } else if c == ']' {
            tokens = append(tokens, token{
                id: tokenCloseBracket,
                value: "]",
            })
        } else if c == '=' {
            tokens = append(tokens, token{
                id: tokenEqual,
                value: "=",
            })
        } else if c == '"' {
            tokens = append(tokens, token{
                id: tokenString,
                value: "\"",
            })
        }
        valueBuilder.Reset()
    }
    return tokens
}

func (l *lexer) consume() rune {
    if l.Cursor >= l.LenSource {
        return ' '
    }
    c := l.Source[l.Cursor]
    l.Cursor++
    return rune(c)
}
