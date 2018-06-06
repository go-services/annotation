package annotation

import (
	"errors"
	"strings"

	"github.com/alecthomas/participle"
)

// value is a helper struct for the parser to parse `parameter="value"`  pairs.
type value struct {
	Key   string `parser:"@Ident'='"`
	Value string `parser:"@String"`
}

// ann is the struct that is used to parse parameters in comments.
type ann struct {
	Name   string   `parser:"'@' @Ident'('"`
	Values []*value `parser:"[@@{','@@}]')'"`
}

// ParseAnnotation finds an ann in a string.
func ParseAnnotation(s string) (Ann, error) {
	if !strings.HasPrefix(s, "@") {
		return nil, nil
	}
	a := &ann{}
	err := parse(a, s)
	if err != nil {
		return nil, err
	}
	ant := NewAnnotation(a.Name)
	for _, v := range a.Values {
		ant.Set(v.Key, v.Value)
	}
	return ant, err
}

// parse is a helper function that builds the parser.
func parse(a interface{}, s string) (err error) {
	p, err := participle.Build(a, nil)
	if err != nil {
		return
	}
	if err := p.ParseString(s, a); err != nil {
		return errors.New("error while parsing annotation")
	}
	return
}
