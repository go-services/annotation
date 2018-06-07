package annotation

// ValueType is a string that tells the type of the parsed parameter
type ValueType string

const (
	STRING  ValueType = "string"
	INT     ValueType = "int"
	FLOAT   ValueType = "float"
	BOOL    ValueType = "bool"
	UNKNOWN ValueType = "unknown"
)

// Value is the interface that wraps the annotation parameter
// tha parameter can be represented in all of the below types,
// if the parameter value is not convertible to a type the value
// will be the zero value of the type
type Value interface {
	String() string
	Int() int
	Float() float64
	Bool() bool
	Type() ValueType
}

// Annotation is the parsed annotation
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

// Fet returns the parameter value by name
// if that parameter does not exist it will return an empty value of type 'UNKNOWN'
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
