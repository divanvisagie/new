package git

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func getTarget(projectName string) string {
	dir, _ := os.Getwd()

	target := strings.Join([]string{dir, projectName}, separator)
	return target
}

func TestGetGitArgs(t *testing.T) {
	type args struct {
		projectURL  string
		projectName string
	}
	tests := []struct {
		name string
		args args
		want *Args
	}{
		{
			name: "given github short url",
			args: args{projectName: "newProject", projectURL: "divanvisagie/new"},
			want: &Args{
				url:    "https://github.com/divanvisagie/new.git",
				target: getTarget("newProject"),
			},
		},
		{
			name: "given full url",
			args: args{projectName: "myNewProject", projectURL: "https://gitlab.com/divan/new.git"},
			want: &Args{
				url:    "https://gitlab.com/divan/new.git",
				target: getTarget("myNewProject"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetArgs(tt.args.projectURL, tt.args.projectName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGitArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
