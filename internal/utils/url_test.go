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
		{name: "Basic", args: args{u: "https://example.com/1a2b3c"}, want: "1a2b3c", wantErr: false},
		{name: "BasicWithDepth", args: args{u: "https://example.com/1a2b3c/4d5e6f"}, want: "4d5e6f", wantErr: false},
		{name: "BasicWithQuery", args: args{u: "https://example.com/1a2b3c?arg=val"}, want: "1a2b3c", wantErr: false},
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
