package lexer

import (
	"reflect"
	"testing"
)

func TestLex(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		want    Tokens
		wantErr bool
	}{
		{
			name:    "empty string",
			payload: "",
			want:    Tokens{},
			wantErr: false,
		},
		{
			name:    "string token",
			payload: `"hello"`,
			want:    Tokens{"hello"},
			wantErr: false,
		},
		{
			name:    "number token",
			payload: "123",
			want:    Tokens{123},
			wantErr: false,
		},
		{
			name:    "boolean token",
			payload: "true",
			want:    Tokens{true},
			wantErr: false,
		},
		{
			name:    "null token",
			payload: "null",
			want:    Tokens{nil},
			wantErr: false,
		},
		{
			name:    "unexpected character",
			payload: "!",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "multiple tokens",
			payload: `"hello", 123, true, null`,
			want: Tokens{
				"hello",
				',',
				123,
				',',
				true,
				',',
				nil,
			},
			wantErr: false,
		},
		{
			name:    "multiple tokens with json syntax",
			payload: `{"name": "Ahmed", "age": null}`,
			want: Tokens{
				'{',
				"name",
				':',
				"Ahmed",
				',',
				"age",
				':',
				nil,
				'}',
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Lex(tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lex() = %v, want %v", got, tt.want)
			}
		})
	}
}
