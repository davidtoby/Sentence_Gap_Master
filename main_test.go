package main

import "testing"

func Test_chooseWord(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name       string
		args       args
		wantPrefix string
		wantSubfix string
		wantWord   string
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				sentence: "The quick brown fox jumps over the lazy dog",
			},
			wantPrefix: "The quick brown fox jumps over the lazy",
			wantSubfix: "quick brown fox jumps over the lazy dog",
			wantWord:   "dog",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrefix, gotSubfix, gotWord := chooseWord(tt.args.sentence)
			if gotPrefix != tt.wantPrefix {
				t.Errorf("chooseWords() gotPrefix = %v, want %v", gotPrefix, tt.wantPrefix)
			}
			if gotSubfix != tt.wantSubfix {
				t.Errorf("chooseWords() gotSubfix = %v, want %v", gotSubfix, tt.wantSubfix)
			}
			if gotWord != tt.wantWord {
				t.Errorf("chooseWords() gotWord = %v, want %v", gotWord, tt.wantWord)
			}
		})
	}
}
