package annotation

import (
	"errors"
	"strings"

	"github.com/alecthomas/participle"
	"strconv"
)

type attrValue struct {
	Str    *string  `parser:"@String"`
	RStr   *string  `parser:"| @RawString"`
	I      *int     `parser:"| @Int"`
	F      *float64 `parser:"| @Float"`
	VTrue  bool     `parser:"| @'true'"`
	VFalse bool     `parser:"| @'false'"`
}

// value is a helper struct for the parser to parse `parameter="value"`  pairs.
type value struct {
	Key   string     `parser:"@Ident'='"`
	Value *attrValue `parser:"@@"`
}

// ann is the struct that is used to parse parameters in comments.
type ann struct {
	Name   string   `parser:"'@' @Ident'('"`
	Values []*value `parser:"[@@{','@@}]')'"`
}

// Parse finds an ann in a string.
func Parse(s string) (*Annotation, error) {
	s = prepareString(s)
	if !strings.HasPrefix(s, "@") {
		return nil, errors.New("annotation not found in string")
	}
	a := &ann{}
	err := parse(a, s)
	if err != nil {
		return nil, err
	}
	ant := NewAnnotation(a.Name)
	for _, v := range a.Values {
		ant.set(v.Key, *v.Value)
	}
	return &ant, err
}

// parse is a helper function that builds the parser.
func parse(a interface{}, s string) (err error) {
	p, err := participle.Build(a)
	if err != nil {
		return err
	}
	if err := p.ParseString(s, a); err != nil {
		return err
	}
	return
}

func prepareString(s string) string {
	return strings.TrimSpace(s)
}

func (v attrValue) String() string {
	switch v.Type() {
	case STRING:
		return *v.Str
	case INT:
		return strconv.Itoa(*v.I)
	case FLOAT:
		return strconv.FormatFloat(*v.F, 'f', 4, 64)
	case BOOL:
		if v.VTrue {
			return "true"
		}
		return "false"
	default:
		return ""
	}
}

func (v attrValue) Int() int {
	switch v.Type() {
	case INT:
		return *v.I
	case FLOAT:
		return int(*v.F)
	default:
		i, _ := strconv.ParseInt(v.String(), 10, strconv.IntSize)
		return int(i)
	}
}

func (v attrValue) Float() float64 {
	switch v.Type() {
	case FLOAT:
		return *v.F
	case INT:
		return float64(*v.I)
	default:
		f, _ := strconv.ParseFloat(v.String(), 64)
		return f
	}
}

func (v attrValue) Bool() bool {
	return v.VTrue
}

func (v attrValue) Type() ValueType {
	if v.I != nil {
		return INT
	} else if v.F != nil {
		return FLOAT
	} else if v.VTrue || v.VFalse {
		return BOOL
	} else if v.Str != nil {
		return STRING
	}
	return UNKNOWN
}
