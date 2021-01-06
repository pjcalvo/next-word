package main

import "testing"

func TestExistingFile(t *testing.T) {
	word, err := getWord("easy")
	if err != nil {
		t.Error("Failed when it shouldnt")
	}
	if len(word) == 0 {
		t.Error("No word was returned")
	}
}
