package annotation

import "fmt"

// Definition describes the Annotation Definition.
type Definition struct {
	// name is the name of the Annotation e.x Hello for // @Hello().
	name string

	// should the definition allow unknown parameters
	allowUnknownParameters bool

	// parameters holds a map of parameters, it is intended to be used as : "parameter_name" : is_it_required
	// if the parameter is required kit will fail if it is missing.
	parameters []ParameterDefinition
}

type ParameterDefinition struct {
	name     string
	required bool
	tp       ValueType
}

func NewParameterDefinition(name string, required bool, parameterType ValueType) ParameterDefinition {
	return ParameterDefinition{
		name:     name,
		required: required,
		tp:       parameterType,
	}
}

// NewDefinition creates a new Annotation Definition.
func NewDefinition(name string, allowUnknownParameters bool, parameters ...ParameterDefinition) Definition {
	return Definition{
		name:                   name,
		allowUnknownParameters: allowUnknownParameters,
		parameters:             parameters,
	}
}

func (d Definition) allowParameter(name string) bool {
	if d.allowUnknownParameters {
		return true
	}
	for _, p := range d.parameters {
		if p.name == name {
			return true
		}
	}
	return false
}

func (d *Definition) Check(annotation Annotation) error {
	if d.name != annotation.name {
		return fmt.Errorf("annotation name `%s` does not match the definition name %s", annotation.name, d.name)
	}
	for k := range annotation.parameters {
		if !d.allowParameter(k) {
			return fmt.Errorf("unknown parameter: `%s` in `@%s()` Annotation", k, d.name)
		}
	}
	for _, p := range d.parameters {
		if err := p.checkParameter(annotation); err != nil {
			return err
		}
	}
	return nil
}

func (p *ParameterDefinition) checkParameter(annotation Annotation) error {
	parameter := annotation.Get(p.name)
	if p.required && parameter.Type() == UNKNOWN {
		return fmt.Errorf("the `%s` parameter is required for @%s() Annotation", p.name, annotation.name)
	}
	if parameter.Type() == UNKNOWN {
		return nil
	}
	if p.tp != parameter.Type() {
		return fmt.Errorf(
			"the `%s` parameter for @%s() Annotation should have be of type `%s`",
			p.name,
			annotation.name,
			p.tp,
		)
	}
	return nil
}
