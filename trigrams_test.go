package whatlanggo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCount(t *testing.T) {
	tests := map[string]map[string]int{
		"":             {"": 0},
		",":            {"": 0},
		"a":            {" a ": 1},
		"-a-":          {" a ": 1},
		"yes":          {" ye": 1, "yes": 1, "es ": 1},
		"Give - IT...": {" gi": 1, "giv": 1, "ive": 1, "ve ": 1, " it": 1, "it ": 1},
	}

	for key, value := range tests {
		got := Count(key)

		for key1, value1 := range value {
			require.Equalf(t, value1, got[key1], "text=%q trigram=%q", key, key1)
		}
	}
}

func TestToTrigramChar(t *testing.T) {
	tests := map[rune]rune{
		'a': 'a', 'z': 'z', 'A': 'A', 'Z': 'Z', 'Ж': 'Ж', 'ß': 'ß',
		// punctuation, digits, ... etc
		'\t': ' ', '\n': ' ', ' ': ' ', '.': ' ', '0': ' ', '9': ' ', ',': ' ', '@': ' ',
		'[': ' ', ']': ' ', '^': ' ', '\\': ' ', '`': ' ', '|': ' ', '{': ' ', '}': ' ', '~': ' ',
	}

	for r, want := range tests {
		got := toTrigramChar(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}
