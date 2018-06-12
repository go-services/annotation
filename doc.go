// Package annotation is used to parse and check java style annotation (e.x @SomeAnnotation(param="value"))
//
// The package can only be used for string that only contain the annotation, it can not find the annotation in an arbitrary string.
//
// Usage
//	annotationString := `@Annotation(
//                            stringParam='String Value',
//                            someInt=2,
//                            someBool=true,
//                            someFloat=2.5
//                        )`
//	ann, _ := annotation.Parse(annotationString)
//	fmt.Printf("Annotation Name: %s\n", ann.Name)                                // Annotation Name: Annotation
//	fmt.Printf("Annotation stringParam = %s\n", ann.Get("stringParam").String()) // Annotation stringParam = String Value
//	fmt.Printf("Annotation someInt = %d\n", ann.Get("someInt").Int())            // Annotation someInt = 2
//	fmt.Printf("Annotation someBool = %t\n", ann.Get("someBool").Bool())         // Annotation someBool = true
//	fmt.Printf("Annotation someFloat = %.4f\n", ann.Get("someFloat").Float())    // Annotation someInt = 2.5000
//
// You can check the annotation using definitions
//
// 	annotationString := `@Annotation(
//                        stringParam='String Value',
//                        someInt=2,
//                        someBool=true,
//                        someFloat=2.5
//                    )`
//	ann, _ := annotation.Parse(annotationString)
//	stringParam := annotation.NewParameterDefinition("stringParam", true, annotation.STRING)
//	someInt := annotation.NewParameterDefinition("someInt", true, annotation.INT)
//	someBool := annotation.NewParameterDefinition("someBool", true, annotation.BOOL)
//	someFloat := annotation.NewParameterDefinition("someFloat", true, annotation.FLOAT)
//	definition := annotation.NewDefinition(
//		"Annotation",
//		false,
//		stringParam,
//		someInt,
//		someBool,
//		someFloat,
//	)
//	err := definition.Check(*ann)
//
//	fmt.Println(err) // <nil>
//
//	annotationMissingParam := "@Annotation(someInt=2, someBool=true, someFloat=2.5)"
//	annWrong, _ := annotation.Parse(annotationMissingParam)
//	err = definition.Check(*annWrong)
//	fmt.Println(err) // the `stringParam` parameter is required for @Annotation() Annotation
package annotation
