package annotation

import (
	"reflect"
	"testing"
)

func TestNewParameterDefinition(t *testing.T) {
	type args struct {
		name          string
		required      bool
		parameterType ValueType
	}
	tests := []struct {
		name string
		args args
		want ParameterDefinition
	}{
		{
			name: "Should return new parameter definition",
			args: args{
				name:          "param",
				required:      true,
				parameterType: STRING,
			},
			want: ParameterDefinition{
				name:     "param",
				required: true,
				tp:       STRING,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParameterDefinition(tt.args.name, tt.args.required, tt.args.parameterType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParameterDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDefinition(t *testing.T) {
	type args struct {
		name                   string
		allowUnknownParameters bool
		parameters             []ParameterDefinition
	}
	tests := []struct {
		name string
		args args
		want Definition
	}{
		{
			name: "Should return a new definition",
			args: args{
				name: "MyDefinition",
				allowUnknownParameters: true,
			},
			want: Definition{
				name: "MyDefinition",
				allowUnknownParameters: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefinition(tt.args.name, tt.args.allowUnknownParameters, tt.args.parameters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefinition_allowParameter(t *testing.T) {
	type fields struct {
		name                   string
		allowUnknownParameters bool
		parameters             []ParameterDefinition
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Should return true if unknown parameters are allowed",
			fields: fields{
				name: "SomeAnnotation",
				allowUnknownParameters: true,
				parameters:             []ParameterDefinition{},
			},
			args: args{
				name: "param",
			},
			want: true,
		},
		{
			name: "Should return true if unknown parameters are not allowed but parameter exists",
			fields: fields{
				name: "SomeAnnotation",
				allowUnknownParameters: false,
				parameters: []ParameterDefinition{
					{
						name:     "param",
						required: true,
						tp:       STRING,
					},
				},
			},
			args: args{
				name: "param",
			},
			want: true,
		},
		{
			name: "Should return false if unknown parameters are not allowed and the parameter does not exists",
			fields: fields{
				name: "SomeAnnotation",
				allowUnknownParameters: false,
				parameters:             []ParameterDefinition{},
			},
			args: args{
				name: "param",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Definition{
				name: tt.fields.name,
				allowUnknownParameters: tt.fields.allowUnknownParameters,
				parameters:             tt.fields.parameters,
			}
			if got := d.allowParameter(tt.args.name); got != tt.want {
				t.Errorf("Definition.allowParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefinition_Check(t *testing.T) {
	type fields struct {
		name                   string
		allowUnknownParameters bool
		parameters             []ParameterDefinition
	}
	type args struct {
		annotation Annotation
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should not return error if the annotation matches the definition",
			fields: fields{
				name: "SomeAnnotation",
				allowUnknownParameters: true,
				parameters: []ParameterDefinition{
					{
						name:     "param",
						required: true,
						tp:       STRING,
					},
				},
			},
			args: args{
				annotation: Annotation{
					Name: "SomeAnnotation",
					parameters: map[string]attrValue{
						"param": {
							Str: pointerString("test"),
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Should not return error if the definition allows unknown parameter",
			fields: fields{
				name: "SomeAnnotation",
				allowUnknownParameters: true,
				parameters:             []ParameterDefinition{},
			},
			args: args{
				annotation: Annotation{
					Name: "SomeAnnotation",
					parameters: map[string]attrValue{
						"param": {
							Str: pointerString("test"),
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Should return error if the annotation Name does not matches the definition Name",
			fields: fields{
				name: "SomeOtherAnnotation",
				allowUnknownParameters: true,
				parameters: []ParameterDefinition{
					{
						name:     "param",
						required: true,
						tp:       STRING,
					},
				},
			},
			args: args{
				annotation: Annotation{
					Name: "SomeAnnotation",
					parameters: map[string]attrValue{
						"param": {
							Str: pointerString("test"),
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error if the definition does not allow unknown parameters but they exist",
			fields: fields{
				name: "SomeAnnotation",
				allowUnknownParameters: false,
				parameters:             []ParameterDefinition{},
			},
			args: args{
				annotation: Annotation{
					Name: "SomeAnnotation",
					parameters: map[string]attrValue{
						"param": {
							Str: pointerString("test"),
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error if the definition parameter definitions return error",
			fields: fields{
				name: "SomeAnnotation",
				allowUnknownParameters: false,
				parameters: []ParameterDefinition{
					{
						name:     "param",
						required: true,
						tp:       STRING,
					},
				},
			},
			args: args{
				annotation: Annotation{
					Name:       "SomeAnnotation",
					parameters: map[string]attrValue{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Definition{
				name: tt.fields.name,
				allowUnknownParameters: tt.fields.allowUnknownParameters,
				parameters:             tt.fields.parameters,
			}
			if err := d.Check(tt.args.annotation); (err != nil) != tt.wantErr {
				t.Errorf("Definition.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterDefinition_checkParameter(t *testing.T) {
	type fields struct {
		name     string
		required bool
		tp       ValueType
	}
	type args struct {
		annotation Annotation
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should not return error if the parameter matches the definition",
			fields: fields{
				name:     "param",
				required: true,
				tp:       STRING,
			},
			args: args{
				annotation: Annotation{
					Name: "SomeAnnotation",
					parameters: map[string]attrValue{
						"param": {
							Str: pointerString("test"),
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Should not return error if the parameter is missing but it is not required",
			fields: fields{
				name:     "param",
				required: false,
				tp:       STRING,
			},
			args: args{
				annotation: Annotation{
					Name:       "SomeAnnotation",
					parameters: map[string]attrValue{},
				},
			},
			wantErr: false,
		},
		{
			name: "Should return error if the parameter is required but it is missing",
			fields: fields{
				name:     "param",
				required: true,
				tp:       STRING,
			},
			args: args{
				annotation: Annotation{
					Name:       "SomeAnnotation",
					parameters: map[string]attrValue{},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error if the type of the parameter does not match",
			fields: fields{
				name:     "param",
				required: true,
				tp:       STRING,
			},
			args: args{
				annotation: Annotation{
					Name: "SomeAnnotation",
					parameters: map[string]attrValue{
						"param": {
							I: pointerInt(2),
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ParameterDefinition{
				name:     tt.fields.name,
				required: tt.fields.required,
				tp:       tt.fields.tp,
			}
			if err := p.checkParameter(tt.args.annotation); (err != nil) != tt.wantErr {
				t.Errorf("ParameterDefinition.checkParameter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
