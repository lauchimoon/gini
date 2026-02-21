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
    l := &lexer{Source: iniString, LenSource: len(iniString), Cursor: 0}
    p := &parser{Tokens: l.lex(), Cursor: 0}
    return p.parse()
}

func (ini *Ini) Get(section, key string) (string, error) {
    sec, ok := (*ini)[section]
    if !ok {
        return "", fmt.Errorf("section %s does not exist.", section)
    }
    val, ok := sec[key]
    if !ok {
        return "", fmt.Errorf("key %s does not exist in section %s.", key, section)
    }
    return val, nil
}
