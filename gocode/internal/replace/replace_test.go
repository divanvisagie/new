package replace

import (
	"reflect"
	"testing"
)

func Test_getAllFilePathsInDirectory(t *testing.T) {
	type args struct {
		targetFolder string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "given test_data example tree",
			args: args{"../../test/testdata/example_tree"},
			want: []string{
				"../../test/testdata/example_tree/onedeep/other_file.txt",
				"../../test/testdata/example_tree/test.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllFilePathsInDirectory(tt.args.targetFolder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllFilePathsInDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}
