package main

import "testing"

func mockUserInput(name string, description string) string {
	return "[[mock-data]]"
}

func Test_fetchRepository(t *testing.T) {

	type args struct {
		seed string
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testbed",
			args: args{
				name: "testbed",
				seed: "divanvisagie/kotlin-tested-seed",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fetchRepository(tt.args.seed, tt.args.name, mockUserInput)
		})
	}
}
