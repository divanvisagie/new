// Package targ means Typed Arg
package targ

import (
	"reflect"
	"testing"
)

func TestNewContainer(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want *Container
	}{
		{
			name: "given one flag",
			args: args{
				[]string{"--verbose"},
			},
			want: &Container{
				Args: []string{"--verbose"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContainer(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFlag(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "given double flag",
			args: args{"--verbose"},
			want: true,
		},
		{
			name: "given single flag",
			args: args{"-v"},
			want: true,
		},
		{
			name: "given arg with dash",
			args: args{"a-b"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFlag(tt.args.s); got != tt.want {
				t.Errorf("isFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_padToSize(t *testing.T) {
	type args struct {
		s    string
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "given string of size 3",
			args: args{"arg", 5},
			want: "arg  ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := padToSize(tt.args.s, tt.args.size); got != tt.want {
				t.Errorf("padToSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTflag_Bool(t *testing.T) {
	type fields struct {
		Arg         string
		name        string
		short       string
		description string
	}
	tests := []struct {
		name   string
		fields *Tflag
		want   bool
	}{
		{
			name: "given flag --help",
			fields: &Tflag{
				Arg:         "--help",
				name:        "--help",
				description: "",
				short:       "-h",
			},
			want: true,
		},
		{
			name: "given flag but no arg",
			fields: &Tflag{
				Arg:         "",
				name:        "--targ",
				description: "",
				short:       "-t",
			},
			want: false,
		},
		{
			name: "given short arg",
			fields: &Tflag{
				Arg:         "-h",
				name:        "--help",
				short:       "-h",
				description: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf := tt.fields
			if got := tf.Bool(); got != tt.want {
				t.Errorf("Tflag.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}
