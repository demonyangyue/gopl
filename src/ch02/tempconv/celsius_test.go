package tempconv

import "testing"

func TestCToK(t *testing.T) {
	type args struct {
		c Celsius
	}
	tests := []struct {
		name string
		args args
		want Kelvin
	}{
		{"Test 0 Celsius", args{Celsius(0)}, Kelvin(273.15)},
		{"Test 0 Kelvin", args{Celsius(-273.15)}, Kelvin(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CToK(tt.args.c); got != tt.want {
				t.Errorf("CToK() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestKToC(t *testing.T) {
	type args struct {
		k Kelvin
	}
	tests := []struct {
		name string
		args args
		want Celsius
	}{
		{"Test 0 Celsius", args{Kelvin(273.15)}, Celsius(0)},
		{"Test 0 Kelvin", args{Kelvin(0)}, Celsius(-273.15)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KToC(tt.args.k); got != tt.want {
				t.Errorf("KToC() = %v, want %v", got, tt.want)
			}
		})
	}
}