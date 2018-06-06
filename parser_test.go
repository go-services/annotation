package annotation

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *Annotation
		wantErr bool
	}{
		{
			name: "Should parse normal annotation string",
			args: args{
				s: "@MyAnnotation()",
			},
			want: &Annotation{
				Name:       "MyAnnotation",
				parameters: map[string]attrValue{},
			},
		},
		{
			name: "Should parse normal annotation string with parameters",
			args: args{
				s: "@MyAnnotation(string='abc',int=1,bool=true,float=2.2)",
			},
			want: &Annotation{
				Name: "MyAnnotation",
				parameters: map[string]attrValue{
					"string": {
						Str: pointerString("abc"),
					},
					"int": {
						I: pointerInt(1),
					},
					"float": {
						F: pointerFloat(2.2),
					},
					"bool": {
						VTrue: true,
					},
				},
			},
		},
		{
			name: "Should parse multi line annotation string with parameters",
			args: args{
				s: `@Annotation(
					Name = "Benjamin Franklin",
					date = "3/27/2003"
					)`,
			},
			want: &Annotation{
				Name: "Annotation",
				parameters: map[string]attrValue{
					"Name": {
						Str: pointerString("Benjamin Franklin"),
					},
					"date": {
						Str: pointerString("3/27/2003"),
					},
				},
			},
		},
		{
			name: "Should return an error if annotation not found",
			args: args{
				s: "No annotation here",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should return an error if annotation is malformed",
			args: args{
				s: "@MyAnnotation(string=)",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_attrValue_String(t *testing.T) {
	type fields struct {
		Str    *string
		I      *int
		F      *float64
		VTrue  bool
		VFalse bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should return the string if string value exists",
			fields: fields{
				Str: pointerString("test"),
			},
			want: "test",
		},
		{
			name: "Should return the string representation of int if int exists",
			fields: fields{
				I: pointerInt(3),
			},
			want: "3",
		},
		{
			name: "Should return the string representation of float if float exists",
			fields: fields{
				F: pointerFloat(3.2),
			},
			want: "3.2000",
		},
		{
			name: "Should return the string representation of bool if bool exists",
			fields: fields{
				VTrue: true,
			},
			want: "true",
		},
		{
			name: "Should return the string representation of bool if bool exists",
			fields: fields{
				VFalse: true,
			},
			want: "false",
		},
		{
			name:   "Should return an empty string for unknown cases",
			fields: fields{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := attrValue{
				Str:    tt.fields.Str,
				I:      tt.fields.I,
				F:      tt.fields.F,
				VTrue:  tt.fields.VTrue,
				VFalse: tt.fields.VFalse,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("attrValue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_attrValue_Int(t *testing.T) {
	type fields struct {
		Str    *string
		I      *int
		F      *float64
		VTrue  bool
		VFalse bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Should return the int if int value exists",
			fields: fields{
				I: pointerInt(4),
			},
			want: 4,
		},
		{
			name: "Should return the int if float exists",
			fields: fields{
				F: pointerFloat(4.6),
			},
			want: 4,
		},
		{
			name: "Should return the int value for any string representation of int",
			fields: fields{
				Str: pointerString("3"),
			},
			want: 3,
		},
		{
			name: "Should return zero for all other cases",
			fields: fields{
				Str: pointerString("test"),
			},
			want: 0,
		},
		{
			name: "Should return zero for all other cases",
			fields: fields{
				VFalse: true,
			},
			want: 0,
		},
		{
			name: "Should return zero for all other cases",
			fields: fields{
				VTrue: true,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := attrValue{
				Str:    tt.fields.Str,
				I:      tt.fields.I,
				F:      tt.fields.F,
				VTrue:  tt.fields.VTrue,
				VFalse: tt.fields.VFalse,
			}
			if got := v.Int(); got != tt.want {
				t.Errorf("attrValue.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_attrValue_Float(t *testing.T) {
	type fields struct {
		Str    *string
		I      *int
		F      *float64
		VTrue  bool
		VFalse bool
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Should return the float if float value exists",
			fields: fields{
				F: pointerFloat(4.6),
			},
			want: 4.6,
		},
		{
			name: "Should return the float if int exists",
			fields: fields{
				I: pointerInt(4),
			},
			want: 4.0,
		},
		{
			name: "Should return float value for any string representation of float ",
			fields: fields{
				Str: pointerString("2.3"),
			},
			want: 2.3,
		},
		{
			name: "Should return zero for all other cases",
			fields: fields{
				Str: pointerString("test"),
			},
			want: 0,
		},
		{
			name: "Should return zero for all other cases",
			fields: fields{
				VFalse: true,
			},
			want: 0,
		},
		{
			name: "Should return zero for all other cases",
			fields: fields{
				VTrue: true,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := attrValue{
				Str:    tt.fields.Str,
				I:      tt.fields.I,
				F:      tt.fields.F,
				VTrue:  tt.fields.VTrue,
				VFalse: tt.fields.VFalse,
			}
			if got := v.Float(); got != tt.want {
				t.Errorf("attrValue.Float() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_attrValue_Bool(t *testing.T) {
	type fields struct {
		Str    *string
		I      *int
		F      *float64
		VTrue  bool
		VFalse bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Should return the bool value if it exists",
			fields: fields{
				VFalse: true,
			},
			want: false,
		},
		{
			name: "Should return the bool value if it exists",
			fields: fields{
				VTrue: true,
			},
			want: true,
		},
		{
			name: "Should return false for all other cases",
			fields: fields{
				Str: pointerString("test"),
			},
			want: false,
		},
		{
			name: "Should return false for all other cases",
			fields: fields{
				I: pointerInt(2),
			},
			want: false,
		},
		{
			name: "Should return false for all other cases",
			fields: fields{
				F: pointerFloat(2.2),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := attrValue{
				Str:    tt.fields.Str,
				I:      tt.fields.I,
				F:      tt.fields.F,
				VTrue:  tt.fields.VTrue,
				VFalse: tt.fields.VFalse,
			}
			if got := v.Bool(); got != tt.want {
				t.Errorf("attrValue.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_attrValue_Type(t *testing.T) {
	type fields struct {
		Str    *string
		I      *int
		F      *float64
		VTrue  bool
		VFalse bool
	}
	tests := []struct {
		name   string
		fields fields
		want   ValueType
	}{
		{
			name: "Should return type of string if string exists",
			fields: fields{
				Str: pointerString("test"),
			},
			want: STRING,
		},
		{
			name: "Should return type of int if int exists",
			fields: fields{
				I: pointerInt(2),
			},
			want: INT,
		},
		{
			name: "Should return type of float if int exists",
			fields: fields{
				F: pointerFloat(2.4),
			},
			want: FLOAT,
		},
		{
			name: "Should return type of bool if bool exists",
			fields: fields{
				VTrue: true,
			},
			want: BOOL,
		},
		{
			name: "Should return type of bool if bool exists",
			fields: fields{
				VFalse: true,
			},
			want: BOOL,
		},
		{
			name: "Should return type of unknown for any unknown case",
			want: UNKNOWN,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := attrValue{
				Str:    tt.fields.Str,
				I:      tt.fields.I,
				F:      tt.fields.F,
				VTrue:  tt.fields.VTrue,
				VFalse: tt.fields.VFalse,
			}
			if got := v.Type(); got != tt.want {
				t.Errorf("attrValue.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type ErrStruct struct {
		Something string
	}
	type args struct {
		a interface{}
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Should throw error if the structure provided does not have correct parser tags",
			wantErr: true,
			args: args{
				a: &ErrStruct{},
				s: "",
			},
		},
		{
			name:    "Should throw error if the object sent is not a struct pointer",
			wantErr: true,
			args: args{
				a: ErrStruct{},
				s: "",
			},
		},
		{
			name:    "Should throw an error if the string that needs to be parsed does not have any annotation",
			wantErr: true,
			args: args{
				a: &ann{},
				s: "Some random string",
			},
		},
		{
			name:    "Should not throw error if the string is an annotation",
			wantErr: false,
			args: args{
				a: &ann{},
				s: "@MyAnnotation()",
			},
		},
		{
			name:    "Should parse string parameter",
			wantErr: false,
			args: args{
				a: &ann{},
				s: "@MyAnnotation(my_string=\"test\")",
			},
		},
		{
			name:    "Should parse int parameter",
			wantErr: false,
			args: args{
				a: &ann{},
				s: "@MyAnnotation(my_int=2)",
			},
		},
		{
			name:    "Should parse annotation multiple lines",
			wantErr: false,
			args: args{
				a: &ann{},
				s: `@Annotation(
					Name = "Benjamin Franklin",
					date = "3/27/2003"
					)`,
			},
		},
		{
			name:    "Should parse float parameter",
			wantErr: false,
			args: args{
				a: &ann{},
				s: "@MyAnnotation(my_float=2.2)",
			},
		},
		{
			name:    "Should parse bool parameter",
			wantErr: false,
			args: args{
				a: &ann{},
				s: "@MyAnnotation(my_bool=true)",
			},
		},
		{
			name:    "Should parse multiple parameter",
			wantErr: false,
			args: args{
				a: &ann{},
				s: "@MyAnnotation(my_string=\"test\",my_int=2,my_float=2.2,my_bool=true)",
			},
		},
		{
			name:    "Should allow single quote string parameter",
			wantErr: false,
			args: args{
				a: &ann{},
				s: "@MyAnnotation(my_string='test')",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parse(tt.args.a, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_prepareString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should return the string without spaces",
			args: args{
				s: "		  @Annotation()		",
			},
			want: "@Annotation()",
		},
		{
			name: "Should not change the inner annotation space",
			args: args{
				s: "		  @Annotation(test=\"Some text   here\")		",
			},
			want: "@Annotation(test=\"Some text   here\")",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareString(tt.args.s); got != tt.want {
				t.Errorf("prepareString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func pointerString(s string) *string {
	return &s
}

func pointerInt(i int) *int {
	return &i
}

func pointerFloat(f float64) *float64 {
	return &f
}
