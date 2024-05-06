package parser

import (
	"reflect"
	"testing"

	"github.com/AhmedGaber77/json-parser/pkg/lexer"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		tokens  lexer.Tokens
		isRoot  bool
		want    interface{}
		want1   lexer.Tokens
		wantErr bool
	}{
		{
			name:    "empty tokens",
			tokens:  lexer.Tokens{},
			isRoot:  true,
			want:    nil,
			want1:   lexer.Tokens{},
			wantErr: true,
		},
		{
			name:    "root not an object",
			tokens:  lexer.Tokens{lexer.JSON_LEFTBRACKET},
			isRoot:  true,
			want:    nil,
			want1:   lexer.Tokens{lexer.JSON_LEFTBRACKET},
			wantErr: true,
		},
		{
			name:    "parse array",
			tokens:  lexer.Tokens{lexer.JSON_LEFTBRACKET, "hello", lexer.JSON_RIGHTBRACKET},
			isRoot:  false,
			want:    []interface{}{"hello"},
			want1:   lexer.Tokens{},
			wantErr: false,
		},
		{
			name:    "parse object",
			tokens:  lexer.Tokens{lexer.JSON_LEFTBRACE, "key", lexer.JSON_COLON, "value", lexer.JSON_RIGHTBRACE},
			isRoot:  false,
			want:    map[string]interface{}{"key": "value"},
			want1:   lexer.Tokens{},
			wantErr: false,
		},
		{
			name:    "parse object with multiple key-value pairs",
			tokens:  lexer.Tokens{lexer.JSON_LEFTBRACE, "key1", lexer.JSON_COLON, "value1", lexer.JSON_COMMA, "key2", lexer.JSON_COLON, "value2", lexer.JSON_RIGHTBRACE},
			isRoot:  false,
			want:    map[string]interface{}{"key1": "value1", "key2": "value2"},
			want1:   lexer.Tokens{},
			wantErr: false,
		},
		{
			name:    "parse nested object",
			tokens:  lexer.Tokens{lexer.JSON_LEFTBRACE, "key", lexer.JSON_COLON, lexer.JSON_LEFTBRACE, "nestedKey", lexer.JSON_COLON, "nestedValue", lexer.JSON_RIGHTBRACE, lexer.JSON_RIGHTBRACE},
			isRoot:  false,
			want:    map[string]interface{}{"key": map[string]interface{}{"nestedKey": "nestedValue"}},
			want1:   lexer.Tokens{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Parse(tt.tokens, tt.isRoot)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
