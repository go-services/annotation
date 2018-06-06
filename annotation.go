package annotation

type ValueType string

const (
	STRING  ValueType = "string"
	INT     ValueType = "int"
	FLOAT   ValueType = "float"
	BOOL    ValueType = "bool"
	UNKNOWN ValueType = "unknown"
)

type Value interface {
	String() string
	Int() int
	Float() float64
	Bool() bool
	Type() ValueType
}

type Annotation struct {
	Name       string
	parameters map[string]attrValue
}

// NewAnnotation creates a new Annotation.
func NewAnnotation(name string) Annotation {
	return Annotation{
		Name:       name,
		parameters: map[string]attrValue{},
	}
}

func (ad *Annotation) Get(name string) Value {
	if ad.parameters == nil {
		return attrValue{}
	}
	if v, ok := ad.parameters[name]; ok {
		return v
	}
	return attrValue{}
}

func (ad *Annotation) set(name string, value attrValue) {
	if ad.parameters != nil {
		ad.parameters[name] = value
	} else {
		ad.parameters = map[string]attrValue{
			name: value,
		}
	}
}
