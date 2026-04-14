package transliterate

import (
	"strings"
	"unicode"
)

// table maps Uzbek Cyrillic runes to their Latin equivalents.
// Longer replacements (digraphs) are handled via multiReplace.
var table = map[rune]string{
	'А': "A", 'а': "a",
	'Б': "B", 'б': "b",
	'В': "V", 'в': "v",
	'Г': "G", 'г': "g",
	'Д': "D", 'д': "d",
	'Е': "E", 'е': "e",
	'Ё': "Yo", 'ё': "yo",
	'Ж': "J", 'ж': "j",
	'З': "Z", 'з': "z",
	'И': "I", 'и': "i",
	'Й': "Y", 'й': "y",
	'К': "K", 'к': "k",
	'Л': "L", 'л': "l",
	'М': "M", 'м': "m",
	'Н': "N", 'н': "n",
	'О': "O", 'о': "o",
	'П': "P", 'п': "p",
	'Р': "R", 'р': "r",
	'С': "S", 'с': "s",
	'Т': "T", 'т': "t",
	'У': "U", 'у': "u",
	'Ф': "F", 'ф': "f",
	'Х': "X", 'х': "x",
	'Ц': "Ts", 'ц': "ts",
	'Ч': "Ch", 'ч': "ch",
	'Ш': "Sh", 'ш': "sh",
	'Щ': "Sh", 'щ': "sh",
	'Ъ': "\u02bc", 'ъ': "\u02bc", // modifier-letter apostrophe
	'Ь': "", 'ь': "",
	'Э': "E", 'э': "e",
	'Ю': "Yu", 'ю': "yu",
	'Я': "Ya", 'я': "ya",
	// Uzbek-specific letters
	'Ў': "O\u02bc", 'ў': "o\u02bc", // Oʻ / oʻ
	'Қ': "Q", 'қ': "q",
	'Ғ': "G\u02bc", 'ғ': "g\u02bc", // Gʻ / gʻ
	'Ҳ': "H", 'ҳ': "h",
}

// multiReplace handles Cyrillic digraphs that must be replaced before
// individual rune substitution.
var multiReplace = []struct{ from, to string }{
	{"НГ", "NG"},
	{"Нг", "Ng"},
	{"нг", "ng"},
}

// HasCyrillic reports whether s contains at least one Cyrillic character.
func HasCyrillic(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
}

// HasLatin reports whether s contains at least one Latin character.
func HasLatin(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Latin, r) {
			return true
		}
	}
	return false
}

// Do converts Uzbek Cyrillic text to the official Uzbek Latin script.
// Non-Cyrillic characters are passed through unchanged.
func Do(text string) string {
	// Handle digraphs first
	for _, m := range multiReplace {
		text = strings.ReplaceAll(text, m.from, m.to)
	}

	var b strings.Builder
	b.Grow(len(text) * 2)
	for _, r := range text {
		if lat, ok := table[r]; ok {
			b.WriteString(lat)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}
