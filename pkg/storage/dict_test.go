package storage

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_parseDict(t *testing.T) {
	type args struct {
		handle io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"Oneline", args{handle: strings.NewReader("holy moly")}, []string{"holy moly"}, false},
		{"Multiline", args{handle: strings.NewReader("holy\nmoly")}, []string{"holy", "moly"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDict(tt.args.handle)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_onlyAlpha(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Only alphabecital characters", args{s: "abcdef"}, true},
		{"One numerical, remaining alphabetical characters", args{s: "1abcdef"}, false},
		{"Only numerical characters", args{s: "12345"}, false},
		{"Only punctuation characters", args{s: "!?@#"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := onlyAlpha(tt.args.s); got != tt.want {
				t.Errorf("onlyAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}
