package annotation

import (
	"fmt"
)

// AnnDefinition describes a definition.
type AnnDefinition interface {
	// Name returns the annotation name.
	Name() string

	// Required tells if the annotation is required.
	Required() bool

	// AllowMultiple tells if the annotation is allowed to appear more than once.
	AllowMultiple() bool

	// Parameters returns the parameters definitions.
	Parameters() map[string]bool

	// HasParameter checks if the parameter is defined.
	HasParameter(name string) bool
}

// definition describes the annotation definition.
type definition struct {
	// name is the name of the annotation e.x Hello for // @Hello().
	name string

	// required tells if the annotation is required, if it is set to true kit will
	// fail if this annotation is missing.
	required bool

	// allowMultiple tells if there can be more than one annotations with the same name present.
	allowMultiple bool

	// parameters holds a map of parameters, it is intended to be used as : "parameter_name" : is_it_required
	// if the parameter is required kit will fail if it is missing.
	parameters map[string]bool
}

// NewDefinition creates a new annotation definition.
func NewDefinition(name string, required bool, parameters map[string]bool, allowMultiple bool) AnnDefinition {
	return &definition{
		name:          name,
		required:      required,
		parameters:    parameters,
		allowMultiple: allowMultiple,
	}
}

func (ad *definition) Name() string {
	return ad.name
}

func (ad *definition) Required() bool {
	return ad.required
}

func (ad *definition) AllowMultiple() bool {
	return ad.allowMultiple
}

func (ad *definition) Parameters() map[string]bool {
	return ad.parameters
}

func (ad *definition) HasParameter(name string) bool {
	if _, ok := ad.parameters[name]; ok {
		return true
	}
	return false
}

// Ann describes an annotation.
type Ann interface {
	// Name returns the name of the annotation.
	Name() string

	// Get returns the value of a given parameter
	// if parameter does not exist it returns an empty string.
	Get(name string) string

	// Set sets the value of the parameter.
	Set(name string, value string)

	// Parameters returns all the parameters and values.
	Parameters() map[string]string
}

type annotation struct {
	name       string
	parameters map[string]string
}

// NewAnnotation creates a new parsed annotation.
func NewAnnotation(name string) Ann {
	return &annotation{
		name:       name,
		parameters: map[string]string{},
	}
}

func (ad *annotation) Get(name string) string {
	if ad.parameters == nil {
		return ""
	}
	if v, ok := ad.parameters[name]; ok {
		return v
	}
	return ""
}

func (ad *annotation) Set(name string, value string) {
	if ad.parameters == nil {
		ad.parameters = map[string]string{
			name: value,
		}
	}
	ad.parameters[name] = value
}

func (ad *annotation) Parameters() map[string]string {
	return ad.parameters
}

func (ad *annotation) Name() string {
	return ad.name
}

// WithAnnotations describes an annotated node.
type WithAnnotations interface {
	GetAnnotation(name string) []Ann
	FindAnnotations(sl []string) error
	HasAllAnnotations(definitions []AnnDefinition) error
	CheckAnnotations(definitions []AnnDefinition) error
	AddAnnotation(an Ann)
}

// Annotated describes a node that has annotations and implements WithAnnotations.
type Annotated struct {
	// Annotations contains the map of annotations
	// where the annotations are stored by name.
	Annotations map[string][]Ann
}

func (a *Annotated) findAnnotation(cm string) (Ann, error) {
	an, err := ParseAnnotation(cm)
	if err != nil {
		return nil, err
	}
	return an, nil
}

// FindAnnotations finds annotations in a comment group.
func (a *Annotated) FindAnnotations(sl []string) error {
	if a.Annotations == nil {
		a.Annotations = map[string][]Ann{}
	}
	for _, c := range sl {
		an, err := a.findAnnotation(c)
		if err != nil {
			return err
		}
		if an != nil {
			a.Annotations[an.Name()] = append(a.Annotations[an.Name()], an)
		}
	}
	return nil
}

// GetAnnotation returns an annotation given the name, if that annotation does not
// exist it will return nil.
func (a Annotated) GetAnnotation(name string) []Ann {
	if v, ok := a.Annotations[name]; ok {
		return v
	}
	return nil
}

// HasAllAnnotations checks if we found all the required annotations.
func (a Annotated) HasAllAnnotations(definitions []AnnDefinition) error {
	for _, v := range definitions {
		if v.Required() && a.Annotations == nil {
			return fmt.Errorf("annotation `%s` is required", v.Name())
		}
		if _, ok := a.Annotations[v.Name()]; !ok && v.Required() {
			return fmt.Errorf("annotation `%s` is required", v.Name())
		}
	}
	return nil
}

// CheckAnnotations checks annotations.
func (a Annotated) CheckAnnotations(definitions []AnnDefinition) error {
	for _, d := range definitions {
		ann := a.GetAnnotation(d.Name())
		return a.checkDefinition(ann, d)
	}
	return nil
}

func (a Annotated) checkAnnotationParameters(ann Ann, d AnnDefinition) error {
	for k := range ann.Parameters() {
		if !d.HasParameter(k) {
			return fmt.Errorf("unknown parameter: `%s` in `@%s()` annotation", k, d.Name())
		}
	}
	for k, p := range d.Parameters() {
		if ann.Get(k) == "" && p {
			return fmt.Errorf("the `%s` parameter is required for @%s() annotation", k, d.Name())
		}
	}
	return nil
}

func (a Annotated) checkDefinition(ann []Ann, d AnnDefinition) error {
	if (ann == nil || len(ann) == 0) && d.Required() {
		return fmt.Errorf("annotation %s is required", d.Name())
	}
	if ann != nil {
		if err := a.checkIfMultipleIsAllowed(ann, d); err != nil {
			return err
		}
	}
	for _, an := range ann {
		if err := a.checkAnnotationParameters(an, d); err != nil {
			return err
		}
	}
	return nil
}

func (a Annotated) checkIfMultipleIsAllowed(ann []Ann, d AnnDefinition) error {
	if len(ann) > 1 && !d.AllowMultiple() {
		return fmt.Errorf(
			"there can not be more than one instance of `%s` annotation",
			d.Name(),
		)
	}
	return nil
}

// AddAnnotation adds an annotation.
func (a *Annotated) AddAnnotation(an Ann) {
	if a.Annotations == nil {
		a.Annotations = map[string][]Ann{}
	}
	a.Annotations[an.Name()] = append(a.Annotations[an.Name()], an)
}
