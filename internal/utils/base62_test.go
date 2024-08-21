package utils

import (
	"testing"
)

var key string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestEncodeBase62(t *testing.T) {
	type args struct {
		id  int64
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{0, key}, "0"},
		{"Test2", args{10, key}, "a"},
		{"Test3", args{61, key}, "Z"},
		{"Test4", args{62, key}, "10"},
		{"Test5", args{100, key}, "1C"},
		{"Test6", args{1000, key}, "g8"},
		{"Test7", args{12345, key}, "3d7"},
		{"Test8", args{100000, key}, "q0U"},
		{"Test9", args{139648545098945, key}, "DEADBEEF"},
		{"Test9", args{10105241, key}, "GoPL"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeBase62(tt.args.id, tt.args.key); got != tt.want {
				t.Errorf("EncodeBase62() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeBase62(t *testing.T) {
	type args struct {
		s   string
		key string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Test1", args{"0", key}, 0},
		{"Test2", args{"a", key}, 10},
		{"Test3", args{"Z", key}, 61},
		{"Test4", args{"10", key}, 62},
		{"Test5", args{"1C", key}, 100},
		{"Test6", args{"g8", key}, 1000},
		{"Test7", args{"3d7", key}, 12345},
		{"Test8", args{"q0U", key}, 100000},
		{"Test9", args{"DEADBEEF", key}, 139648545098945},
		{"Test10", args{"GoPL", key}, 10105241},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeBase62(tt.args.s, tt.args.key); got != tt.want {
				t.Errorf("DecodeBase62() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{"abc"}, "cba"},
		{"Test2", args{"abcd"}, "dcba"},
		{"Test3", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
