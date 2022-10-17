package typo

type TypoFlags struct {
	ExtraKey bool //InsertedKey()
	Skip     bool //SkipLetter()
	Double   bool //DoubleLetter()
	Reverse  bool //ReverseLetter()
	Vowel    bool //WrongVowel()
	Key      bool //WrongKey()
}

func NewTypoFlags() TypoFlags {
	return TypoFlags{
		ExtraKey: false,
		Skip:     false,
		Double:   false,
		Reverse:  false,
		Vowel:    false,
		Key:      false,
	}
}

type GeneratorFlags struct {
	Typo         TypoFlags
	Excluded     []string
	SuppressLogs bool
}
