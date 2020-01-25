package annotation

import (
	"reflect"
	"testing"
)

func TestNewAnnotation(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want Annotation
	}{
		{
			name: "Should return a new annotation",
			args: args{
				name: "MyAnnotation",
			},
			want: Annotation{
				Name:       "MyAnnotation",
				parameters: map[string]attrValue{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAnnotation(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnnotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnnotation_Get(t *testing.T) {
	type fields struct {
		name       string
		parameters map[string]attrValue
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Value
	}{
		{
			name: "Should return the value with specific type",
			fields: fields{
				name: "MyAnnotation",
				parameters: map[string]attrValue{
					"test": {
						Str: pointerString("abc"),
					},
				},
			},
			args: args{
				name: "test",
			},
			want: attrValue{
				Str: pointerString("abc"),
			},
		},
		{
			name: "Should return an empty value if the parameter does not exist",
			fields: fields{
				name:       "MyAnnotation",
				parameters: map[string]attrValue{},
			},
			args: args{
				name: "test",
			},
			want: attrValue{},
		},
		{
			name: "Should return an empty value if the parameters is nil",
			fields: fields{
				name: "MyAnnotation",
			},
			args: args{
				name: "test",
			},
			want: attrValue{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ad := &Annotation{
				Name:       tt.fields.name,
				parameters: tt.fields.parameters,
			}
			if got := ad.Get(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Annotation.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnnotation_set(t *testing.T) {
	type fields struct {
		name       string
		parameters map[string]attrValue
	}
	type args struct {
		name  string
		value attrValue
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]attrValue
	}{
		{
			name: "Should Set the attribute",
			fields: fields{
				name:       "MyAnnotation",
				parameters: map[string]attrValue{},
			},
			args: args{
				name: "test",
				value: attrValue{
					Str: pointerString("abc"),
				},
			},
			want: map[string]attrValue{
				"test": {
					Str: pointerString("abc"),
				},
			},
		},
		{
			name: "Should handle nil parameters",
			fields: fields{
				name: "MyAnnotation",
			},
			args: args{
				name: "test",
				value: attrValue{
					Str: pointerString("abc"),
				},
			},
			want: map[string]attrValue{
				"test": {
					Str: pointerString("abc"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ad := &Annotation{
				Name:       tt.fields.name,
				parameters: tt.fields.parameters,
			}
			ad.Set(tt.args.name, tt.args.value)
			if !reflect.DeepEqual(ad.parameters, tt.want) {
				t.Errorf("Annotation.parameters is %v, want %v", ad.parameters, tt.want)
			}
		})
	}
}

func TestAnnotation_String(t *testing.T) {
	type fields struct {
		Name       string
		parameters map[string]attrValue
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should return the correct string representation of the annotation",
			fields: fields{
				Name:       "MyAnnotation",
				parameters: map[string]attrValue{},
			},
			want: "@MyAnnotation()",
		},
		{
			name: "Should return the correct string representation of the annotation",
			fields: fields{
				Name: "MyAnnotation",
				parameters: map[string]attrValue{
					"string": {
						Str: pointerString("abc"),
					},
				},
			},
			want: "@MyAnnotation(string=\"abc\")",
		},
		{
			name: "Should return the correct string representation of the annotation",
			fields: fields{
				Name: "MyAnnotation",
				parameters: map[string]attrValue{
					"int": {
						I: pointerInt(1),
					},
				},
			},
			want: "@MyAnnotation(int=1)",
		},
		{
			name: "Should return the correct string representation of the annotation",
			fields: fields{
				Name: "MyAnnotation",
				parameters: map[string]attrValue{
					"float": {
						F: pointerFloat(2.2),
					},
				},
			},
			want: "@MyAnnotation(float=2.2000)",
		},
		{
			name: "Should return the correct string representation of the annotation",
			fields: fields{
				Name: "MyAnnotation",
				parameters: map[string]attrValue{
					"bool": {
						VTrue: true,
					},
				},
			},
			want: "@MyAnnotation(bool=true)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ad := &Annotation{
				Name:       tt.fields.Name,
				parameters: tt.fields.parameters,
			}
			if got := ad.String(); got != tt.want {
				t.Errorf("Annotation.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
