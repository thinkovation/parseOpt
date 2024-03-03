package main

import (
	"parse/brcparsers"
	"strconv"
	"testing"
)

//FastParserBareBones

func Test_FastParserBareBones(t *testing.T) {

	type testStruct struct {
		input string
		want  int64
	}

	tests := []testStruct{}
	for i := -999; i <= 999; i++ {
		var ts testStruct
		num := float64(i) / 10.0

		ts.input = strconv.FormatFloat(num, 'f', 1, 64)
		ts.want = int64(i)

	}

	for _, tt := range tests {
		t.Run("Test", func(t *testing.T) {
			got, _ := brcparsers.FastParserBareBones(tt.input)
			if got != tt.want {
				t.Errorf("FastParserBareBones= %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_customStringToIntParserChat(t *testing.T) {

	type testStruct struct {
		input string
		want  int64
	}

	tests := []testStruct{}
	for i := -999; i <= 999; i++ {
		var ts testStruct
		num := float64(i) / 10.0

		ts.input = strconv.FormatFloat(num, 'f', 1, 64)
		ts.want = int64(i)

	}

	for _, tt := range tests {
		t.Run("Test", func(t *testing.T) {
			got, _ := brcparsers.CustomBRCParserChatGPT(tt.input)
			if got != tt.want {
				t.Errorf("customStringToIntParserChat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_customStringToIntParserIfElse(t *testing.T) {
	type testStruct struct {
		input string
		want  int64
	}

	tests := []testStruct{}
	for i := -999; i <= 999; i++ {
		var ts testStruct
		num := float64(i) / 10.0

		ts.input = strconv.FormatFloat(num, 'f', 1, 64)
		ts.want = int64(i)

	}

	for _, tt := range tests {
		t.Run("Test", func(t *testing.T) {
			if got, _ := brcparsers.FastParserOptimised(tt.input); got != tt.want {
				t.Errorf("customStringToIntParserChat() = %v, want %v", got, tt.want)
			}
		})
	}
}
