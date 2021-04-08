package expand

import "testing"

func Test_expand(t *testing.T) {
	type args struct {
		s string
		f func(string) string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
		"Single placeholder should be replaced successfully",
			args{ "hello $world",
				func(s string) string {
					return s + s
				},
			},
		"hello worldworld",
		},


		{
			"Multiple placeholders should be replaced successfully",
			args{ "hello $world ${boy}",
				func(s string) string {
					return s + s
				},
			},
			"hello worldworld boyboy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expand(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("expand() = %v, want %v", got, tt.want)
			}
		})
	}
}