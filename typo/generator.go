package typo

import (
	"log"
)

func TypoGenerator(baseline string, flags TypoFlags) []string {
	if flags == NewTypoFlags() {
		log.Fatal("ERROR: At least one typo flag must be specified")
	}
	t := NewTypos(baseline)
	if flags.ExtraKey {
		t.InsertedKey()
	}
	if flags.Skip {
		t.SkipLetter()
	}
	if flags.Double {
		t.DoubleLetter()
	}
	if flags.Reverse {
		t.ReverseLetter()
	}
	if flags.Vowel {
		t.WrongVowel()
	}
	if flags.Key {
		t.WrongKey()
	}
	return t.Typos
}
