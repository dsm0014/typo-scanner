package typo

import (
	"log"
)

func TypoGenerator(baseline string, genFlags GeneratorFlags) []string {
	if genFlags.Typo == NewTypoFlags() {
		log.Fatal("ERROR: At least one typo flag must be specified")
	}
	t := NewTypos(baseline, genFlags.Excluded)
	if genFlags.Typo.ExtraKey {
		t.InsertedKey()
	}
	if genFlags.Typo.Skip {
		t.SkipLetter()
	}
	if genFlags.Typo.Double {
		t.DoubleLetter()
	}
	if genFlags.Typo.Reverse {
		t.ReverseLetter()
	}
	if genFlags.Typo.Vowel {
		t.WrongVowel()
	}
	if genFlags.Typo.Key {
		t.WrongKey()
	}
	return t.Typos
}
