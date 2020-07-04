package printer

import "testing"

func Test_detectNewline(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "given a string with windows line endings",
			args: args{" I am a \r\n windows \r\n string"},
			want: "\r\n",
		},
		{
			name: "given a string with unix line endings",
			args: args{"I aam a \n unix \n string"},
			want: "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectNewline(tt.args.text); got != tt.want {
				t.Errorf("detectNewline() = %v, want %v", got, tt.want)
			}
		})
	}
}
