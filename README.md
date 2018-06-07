# Annotate [![Go Report Card](https://goreportcard.com/badge/github.com/go-services/annotation)](https://goreportcard.com/report/github.com/go-services/annotation) [![Coverage Status](https://coveralls.io/repos/github/go-services/annotation/badge.svg?branch=master)](https://coveralls.io/github/go-services/annotation?branch=master) [![Build Status](https://travis-ci.org/go-services/annotation.svg?branch=master)](https://travis-ci.org/go-services/annotation)

Annotate is a go package that allows parsing and read `Java style` annotation strings e.x `@SomeAnnotation(param="value")`

## Install
Using `go get`
```bash
go get -u github.com/go-services/annotate
```
If you are using `dep` 
```bash
dep ensure --add github.com/go-services/annotate
```

## Usage
The package can only be used for string that only contain the annotation, it can not find the annotation in an arbitrary string.

Example
```go
package main

import (
	"github.com/go-services/annotation"
	"fmt"
)

func main() {
	annotationString := "@Annotation(stringParam='String Value', someInt=2, someBool=true, someFloat=2.5)"
	ann, _ := annotation.Parse(annotationString)
	fmt.Printf("Annotation Name: %s\n", ann.Name)                                // Annotation Name: Annotation
	fmt.Printf("Annotation stringParam = %s\n", ann.Get("stringParam").String()) // Annotation stringParam = String Value
	fmt.Printf("Annotation someInt = %d\n", ann.Get("someInt").Int())            // Annotation someInt = 2
	fmt.Printf("Annotation someBool = %t\n", ann.Get("someBool").Bool())         // Annotation someBool = true
	fmt.Printf("Annotation someFloat = %.4f\n", ann.Get("someFloat").Float())    // Annotation someInt = 2.5000
}
```
## Definitions
Annotate also provides a way to check if the annotation has the correct parameters, name.

**ParameterDefinition**
```go
type ParameterDefinition struct {
	// name the parameter name
	name     string
	
	// required tells if the parameter is required
	required bool
	
	// tp shows the required type of the annotation
	tp       ValueType
}
```
**Definition**
```go
// Definition describes the Annotation definition.
type Definition struct {
	// Name is the Name of the Annotation e.x Hello for // @Hello().
	name string

	// should the definition allow unknown parameters
	allowUnknownParameters bool

	// parameters has a list of parameter definitions
	parameters []ParameterDefinition
}
```
Example
```go
package main

import (
	"github.com/go-services/annotation"
	"fmt"
)

func main() {
	annotationString := "@Annotation(stringParam='String Value', someInt=2, someBool=true, someFloat=2.5)"
	ann, _ := annotation.Parse(annotationString)
	stringParam := annotation.NewParameterDefinition("stringParam", true, annotation.STRING)
	someInt := annotation.NewParameterDefinition("someInt", true, annotation.INT)
	someBool := annotation.NewParameterDefinition("someBool", true, annotation.BOOL)
	someFloat := annotation.NewParameterDefinition("someFloat", true, annotation.FLOAT)
	definition := annotation.NewDefinition(
		"Annotation",
		false,
		stringParam,
		someInt,
		someBool,
		someFloat,
	)
	err := definition.Check(*ann)

	fmt.Println(err) // <nil>

	annotationMissingParam := "@Annotation(someInt=2, someBool=true, someFloat=2.5)"
	annWrong, _ := annotation.Parse(annotationMissingParam)
	err = definition.Check(*annWrong)
	fmt.Println(err) // the `stringParam` parameter is required for @Annotation() Annotation

}
```