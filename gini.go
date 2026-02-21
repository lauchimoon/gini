package gini

import (
    "io"
    "os"
)

type Ini map[string]Section

type Section map[string]string

type reader struct {
    Source string
    LenSource int
    Cursor int
}

func NewFromFile(filePath string) (*Ini, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    contents, err := io.ReadAll(f)
    if err != nil {
        return nil, err
    }
    return NewFromString(string(contents))
}

func NewFromString(iniString string) (*Ini, error) {
    l := &lexer{Source: iniString, LenSource: len(iniString), Cursor: 0}
    p := &parser{Tokens: l.lex(), Cursor: 0}
    return p.parse()
}

