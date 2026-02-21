package gini

const (
    tokenComment = iota
    tokenOpenBracket
    tokenCloseBracket
    tokenSymbol
    tokenEqual
    tokenString
)

type tokenId int

type token struct {
    id    tokenId
    value string
}

func (id tokenId) String() string {
    switch id {
        case tokenComment: return "Comment"
        case tokenOpenBracket: return "Open section"
        case tokenCloseBracket: return "Close section"
        case tokenSymbol: return "Symbol"
        case tokenEqual: return "Equal"
        case tokenString: return "String"
    }
    return "Unknown"
}
