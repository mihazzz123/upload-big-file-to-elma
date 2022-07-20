package casey

import (
	"strings"
	"unicode"
)

// Tokens after split
type Tokens []string

// Camel splits camelCase
func Camel(s string) Tokens {
	cs := new(camelSplitter)
	return cs.Split(s)
}

const (
	unknown = iota
	lower
	number
	upper
)

type camelSplitter struct {
	state  int
	chunk  []rune
	chunks []string
}

// Split string
func (cs *camelSplitter) Split(s string) Tokens {
	if len(s) == 0 {
		return nil
	}

	runes := []rune(s)

	cs.state = unknown
	cs.chunk = make([]rune, 0, len(runes))
	cs.chunks = cs.chunks[:0]

	for i := len(runes) - 1; i >= 0; i-- {
		cs.addRune(runes[i])
	}

	cs.flush(false)
	cs.reverse()

	return cs.chunks
}

func (cs *camelSplitter) addRune(r rune) {
	cs.chunk = append(cs.chunk, r)
	switch {
	case unicode.IsUpper(r):
		if cs.state == lower {
			cs.flush(false)
			cs.state = unknown
			return
		}
		cs.state = upper

	case unicode.IsNumber(r):
		cs.state = number

	default:
		if cs.state == upper {
			cs.flush(true)
		}
		cs.state = lower
	}
}

func (cs *camelSplitter) flush(preserveLast bool) {
	if len(cs.chunk) > 0 {
		for i, j := 0, len(cs.chunk)-1; i < j; i, j = i+1, j-1 {
			cs.chunk[i], cs.chunk[j] = cs.chunk[j], cs.chunk[i]
		}
		if preserveLast {
			cs.chunks = append(cs.chunks, string(cs.chunk[1:]))
			cs.chunk = cs.chunk[:1]
		} else {
			cs.chunks = append(cs.chunks, string(cs.chunk))
			cs.chunk = cs.chunk[:0]
		}
	}
}

func (cs *camelSplitter) reverse() {
	for i, j := 0, len(cs.chunks)-1; i < j; i, j = i+1, j-1 {
		cs.chunks[i], cs.chunks[j] = cs.chunks[j], cs.chunks[i]
	}
}

// Map tokens with function
func (tokens Tokens) Map(fn func(string) string) Tokens {
	res := make(Tokens, len(tokens))
	for i := range tokens {
		res[i] = fn(tokens[i])
	}
	return res
}

// Capital convert tokens to CapitalCase
func (tokens Tokens) Capital() string {
	return strings.Join(tokens.Map(strings.Title), "")
}

// Camel convert tokens to camelCase
func (tokens Tokens) Camel() string {
	res := tokens.Map(strings.Title)
	if len(res) > 0 {
		res[0] = strings.ToLower(res[0])
	}
	return strings.Join(res, "")
}

// SNAKE convert tokens to SNAKE_CASE
func (tokens Tokens) SNAKE() string {
	return strings.Join(tokens.Map(strings.ToUpper), "_")
}

// Snake convert tokens to snake_case
func (tokens Tokens) Snake() string {
	return strings.Join(tokens.Map(strings.ToLower), "_")
}

// Kebab convert tokens to kebab-case
func (tokens Tokens) Kebab() string {
	return strings.Join(tokens, "-")
}
