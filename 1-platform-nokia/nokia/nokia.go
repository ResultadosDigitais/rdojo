package nokia

var numberLetterMap = map[rune]string{
	'A': "2",
	'B': "22",
	'C': "222",
	'D': "3",
	'E': "33",
	'F': "333",
	'G': "4",
	'H': "44",
	'I': "444",
	'J': "5",
	'K': "55",
	'L': "555",
	'M': "6",
	'N': "66",
	'O': "666",
	'P': "7",
	'Q': "77",
	'R': "777",
	'S': "7777",
	'T': "8",
	'U': "88",
	'V': "888",
	'W': "9",
	'X': "99",
	'Y': "999",
	'Z': "9999",
	' ': "0",
}

type UnsuportedCharacterError struct {
}

type MaximumLengthError struct {
}

func (u UnsuportedCharacterError) Error() string {
	return "Unsuported character!"
}

func (u MaximumLengthError) Error() string {
	return "Text exceeds the limit of 255"
}

func getTypedSequence(text string) (string, error) {
	if len(text) > 255 {
		return "", MaximumLengthError{}
	}

	var output string
	var last byte
	for _, letter := range text {
		transliterated, ok := numberLetterMap[letter]
		if !ok {
			return "", UnsuportedCharacterError{}
		}
		if last == transliterated[0] {
			output += "_"
		}
		output += transliterated
		last = transliterated[0]
	}

	return output, nil
}
