package main

import (
	"testing"
)

func TestSetWords(t *testing.T) {
	w := []string{"a", "b", "c"}
	SetWords(w)

	for i, word := range w {
		if word != words[i] {
			t.Error("Not equal set words")
		}
	}
}

func TestAutoComplete(t *testing.T) {
	tests := []struct {
		search   string
		input    []string
		expected []string
	}{
		{
			search:   "",
			input:    []string{"foo", "bar"},
			expected: []string{"foo", "bar"},
		},
		{
			search:   "foo",
			input:    []string{"foo", "bar"},
			expected: []string{"foo"},
		},
		{
			search:   "dog",
			input:    []string{"door"},
			expected: []string{},
		},
		{
			search:   "do",
			input:    []string{"door"},
			expected: []string{"door"},
		},
		{
			search:   "do",
			input:    []string{"Door"},
			expected: []string{"Door"},
		},
		{
			search:   "Do",
			input:    []string{"door"},
			expected: []string{"door"},
		},
		{
			search:   "Dooooor",
			input:    []string{"door"},
			expected: []string{},
		},
		{
			search:   "Foo",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, test := range tests {
		w := test.input
		SetWords(w)
		results := AutoComplete(test.search)

		if len(test.expected) != len(results) {
			t.Error("Not equal size")
		}

		for i, word := range test.expected {
			if word != results[i] {
				t.Error("Not equal set words")
			}
		}
	}
}
