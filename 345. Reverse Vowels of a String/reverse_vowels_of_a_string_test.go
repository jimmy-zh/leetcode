package reverse_vowels_of_a_string

import (
	"testing"
)

func TestReverseVowels(t *testing.T) {
	input := []string{
		"hello",
		"ehllo",
		"hlleo",
		"eOhll",
		"hhLLh",
		"heEoh",
		"heheO",
		"sdfSODJFsdfjjOESRFNODsjHFOWNszdnoneqKVISsdnSLDoowO",
	}
	result := []string{
		"holle",
		"ohlle",
		"hlloe",
		"Oehll",
		"hhLLh",
		"hoEeh",
		"hOhee",
		"sdfSODJFsdfjjooSRFNIDsjHFeWNszdnonOqKVOSsdnSLDEOwO",
	}

	for i := range input {
		if reverseVowels(input[i]) != result[i] {
			t.Error("reverse error")
		}
	}
}
