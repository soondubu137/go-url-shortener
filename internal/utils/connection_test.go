package utils

import "testing"

func TestCanConnect(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Google", args{url: "https://www.google.com"}, true},
		{"Invalid", args{url: "https://invalid.google.com"}, false},
		{"404", args{url: "https://github.com/SoonDubu923/does-not-exist"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanConnect(tt.args.url); got != tt.want {
				t.Errorf("TestConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
