package storage

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_readWords(t *testing.T) {
	type args struct {
		dict io.Reader
	}
	tests := []struct {
		name      string
		args      args
		wantWords []string
	}{
		{"Oneline", args{dict: strings.NewReader("holy moly")}, []string{"holy moly"}},
		{"Multiline", args{dict: strings.NewReader("holy\nmoly")}, []string{"holy", "moly"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWords := readWords(tt.args.dict); !reflect.DeepEqual(gotWords, tt.wantWords) {
				t.Errorf("readWords() = %v, want %v", gotWords, tt.wantWords)
			}
		})
	}
}

func Test_parseDict(t *testing.T) {
	type args struct {
		dict io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"Oneline", args{dict: strings.NewReader("holy moly")}, []string{"holy moly"}, false},
		{"Multiline", args{dict: strings.NewReader("holy\nmoly")}, []string{"holy", "moly"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDict(tt.args.dict)
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
