package prom

import (
	"testing"
)

func Test_gigabytesToBytes(t *testing.T) {
	type args struct {
		gb float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"gb to bytes",
			args{
				gb: 2,
			},
			2e+9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gigabytesToBytes(tt.args.gb); got != tt.want {
				t.Errorf("gigabytesToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_microsecondsToSeconds(t *testing.T) {
	type args struct {
		microSeconds float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"µs to s",
			args{
				microSeconds: 2e+6,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := microsecondsToSeconds(tt.args.microSeconds); got != tt.want {
				t.Errorf("microsecondsToSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_millisecondsToSeconds(t *testing.T) {
	type args struct {
		milliseconds float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"ms to s",
			args{
				milliseconds: 2e+3,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := millisecondsToSeconds(tt.args.milliseconds); got != tt.want {
				t.Errorf("millisecondsToSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}
