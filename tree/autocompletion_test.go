package tree

import (
	"slices"
	"testing"
)

/*
TestAutocompletion tests solution(s) with the following signature and problem description:

	func (t *trie) AutoComplete(word string) []string

Given a dictionary of words and a word, return autocomplete suggestions from the dictionary that
start with the given word.

For example given {"car","caravan","card","carpet","cap","ca"}, and the word "car", return
{"avan","d","pet"} because they are all words that start with "car".
*/
func TestAutocompletion(t *testing.T) {
	tests := []struct {
		input       string
		dict        []string
		suggestions []string
	}{
		{"a", []string{}, []string{}},
		{"a", []string{"a", "aa", "aaa", "aaaa", "aaaaa"}, []string{"a", "aa", "aaa", "aaaa"}},
		{"car", []string{"car", "caravan", "card", "carpet", "cap", "ca"}, []string{"avan", "d", "pet"}},
		{"a", []string{"aaaaa", "aaaa", "aaa", "aa", "a"}, []string{"a", "aa", "aaa", "aaaa"}},
		{"a", []string{"abc", "abd", "abe", "bbb"}, []string{"bc", "bd", "be"}},
		{"a", []string{"abc", "abd", "abe"}, []string{"bc", "bd", "be"}},
		{"a", []string{"aac", "abc", "acc"}, []string{"ac", "bc", "cc"}},
		{"ab", []string{"abcdfg", "abdefg", "abefgh"}, []string{"cdfg", "defg", "efgh"}},
		{"ab", []string{"abc", "abd", "abe"}, []string{"c", "d", "e"}},
		{"ab", []string{"abc", "abcc", "abccd", "abd", "abe"}, []string{"c", "cc", "ccd", "d", "e"}},
		{"ab", []string{"abc", "abcc", "abcd", "abd", "abe"}, []string{"c", "cc", "cd", "d", "e"}},
		{"ab", []string{"abc", "abcd", "abd", "abcc", "abe"}, []string{"c", "cc", "cd", "d", "e"}},
		{"car", []string{"car", "caravan", "card", "carpet", "cap", "ca"}, []string{"avan", "d", "pet"}},
	}

	for i, test := range tests {
		trie := newTrie(test.dict)
		got := trie.Autocompletion(test.input)
		if !slices.Equal(got, test.suggestions) {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.suggestions, got)
		}
	}
}
