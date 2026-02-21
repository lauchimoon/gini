package gini

import (
    "errors"
    "fmt"
    "strings"
)

type parser struct {
    Tokens []token
    Cursor int
}

func (p *parser) parse() (*Ini, error) {
    ini := make(Ini)
    currentSection := ""
    for p.Cursor < len(p.Tokens) {
        tok := p.current()
        // Ignore comments
        for tok.id == tokenComment {
            p.advance()
            tok = p.current()
        }
        // Make section
        if tok.id == tokenOpenBracket {
            p.advance()
            sectionName := p.current().value
            next := p.next()
            if next.id != tokenCloseBracket {
                return nil, fmt.Errorf("expected closing bracket, found %s", next.value)
            }
            p.advance()
            currentSection = sectionName
            ini[currentSection] = make(Section)
        }
        // Add items to section
        if tok.id == tokenEqual {
            prev := p.previous()
            if prev.id != tokenSymbol {
                return nil, fmt.Errorf("expected symbol, found %s", prev.id)
            }
            key := prev.value
            next := p.next()
            p.advance()
            var value string
            if p.current().id == tokenString {
                var err error
                value, err = p.parseString()
                if err != nil {
                    return nil, err
                }
            } else {
                value = next.value
            }
            ini[currentSection][key] = value
        }
        p.advance()
    }
    return &ini, nil
}

func (p *parser) parseString() (string, error) {
    value := strings.Builder{}
    value.WriteByte('"')
    if p.current().id != tokenString {
        return "", fmt.Errorf("expected string, found %s", p.current().id)
    }
    p.advance()

    for p.Cursor < len(p.Tokens) && p.current().id != tokenString {
        value.WriteString(p.current().value)
        p.advance()
    }

    if p.Cursor >= len(p.Tokens) {
        return "", errors.New("string was not terminated")
    }
    value.WriteByte('"')
    return value.String(), nil
}

func (p *parser) current() token {
    if p.Cursor >= len(p.Tokens) {
        return token{}
    }
    return p.Tokens[p.Cursor]
}

func (p *parser) previous() token {
    if p.Cursor - 1 < 0 {
        return token{}
    }
    return p.Tokens[p.Cursor - 1]
}

func (p *parser) next() token {
    if p.Cursor + 1 >= len(p.Tokens) {
        return token{}
    }
    return p.Tokens[p.Cursor + 1]
}

func (p *parser) advance() {
    if p.Cursor >= len(p.Tokens) {
        return
    }
    p.Cursor++
}
