package utils

import "testing"

func TestGetBaseURL(t *testing.T) {
	type args struct {
		u string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Basic", args{u: "https://example.com/1a2b3c"}, "1a2b3c", false},
		{"BasicWithDepth", args{u: "https://example.com/1a2b3c/4d5e6f"}, "4d5e6f", false},
		{"BasicWithQuery", args{u: "https://example.com/1a2b3c?arg=val"}, "1a2b3c", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBaseURL(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBaseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
