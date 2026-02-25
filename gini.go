package gini

import (
    "fmt"
    "io"
    "sort"
    "os"
)

type Ini map[string]Section

type Section map[string]string

func LoadFromFile(filePath string) (Ini, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    contents, err := io.ReadAll(f)
    if err != nil {
        return nil, err
    }
    return LoadFromString(string(contents))
}

func LoadFromString(iniString string) (Ini, error) {
    l := &lexer{Source: iniString, LenSource: len(iniString), Cursor: 0}
    p := &parser{Tokens: l.lex(), Cursor: 0}
    return p.parse()
}

func Dump(data Ini, outFile io.Writer) {
    sections := []string{}
    for section := range data {
        sections = append(sections, section)
    }
    sort.Strings(sections)

    for i, section := range sections {
        fmt.Fprintf(outFile, "[%s]\n", section)
        keys := []string{}
        for key := range data[section] {
            keys = append(keys, key)
        }

        for _, key := range keys {
            fmt.Fprintln(outFile, key, "=", data[section][key])
        }

        // Don't write last newline
        if i < len(sections) - 1 {
            fmt.Fprintln(outFile)
        }
    }
}

func (ini Ini) Get(section, key string) (string, error) {
    sec, ok := ini[section]
    if !ok {
        return "", fmt.Errorf("section '%s' does not exist.", section)
    }
    val, ok := sec[key]
    if !ok {
        return "", fmt.Errorf("key '%s' does not exist in section '%s'.", key, section)
    }
    return val, nil
}
