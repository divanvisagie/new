package prompt

import (
	"reflect"
	"testing"
)

func Test_readYamlFile(t *testing.T) {

	type args struct {
		configFilePath string
	}
	tests := []struct {
		name string
		args args
		want *NewConfig
	}{
		// TODO: Add test cases.
		{
			name: "given test file",
			args: args{"../../test/testdata/.new.yml"},
			want: &NewConfig{
				Replace: Replace{
					Strings: []ReplacementString{
						ReplacementString{
							Description: "Your pojects favorite type of cheese",
							Match:       "{{CHEESE}}",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readConfigFile(tt.args.configFilePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readYamlFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
