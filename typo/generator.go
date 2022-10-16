package typo

func TypoGenerator(baseline string, flags TypoFlags) []string {
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
