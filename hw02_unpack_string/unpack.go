package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var unpackedString strings.Builder
	runeCount := utf8.RuneCountInString(str)
	runeSliceStr := []rune(str)

	for i := 0; i < runeCount; {
		curSymbolInRune := runeSliceStr[i]

		if unicode.IsNumber(curSymbolInRune) {
			return "", ErrInvalidString
		}

		if i < runeCount-1 && unicode.IsNumber(runeSliceStr[i+1]) {
			numberOfRepeating := int(runeSliceStr[i+1] - '0')

			for repeatIndex := 0; repeatIndex < numberOfRepeating; repeatIndex++ {
				unpackedString.WriteRune(curSymbolInRune)
			}
			i += 2
			continue
		}

		unpackedString.WriteRune(curSymbolInRune)
		i++
	}

	return unpackedString.String(), nil
}
