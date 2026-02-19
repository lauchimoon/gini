package gini

import (
    "fmt"
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
    ini := Ini{}
    l := &lexer{Source: iniString, LenSource: len(iniString), Cursor: 0}
    fmt.Println(l.lex())
    return &ini, nil
}

