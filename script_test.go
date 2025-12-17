package whatlanggo

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func TestDetectScript(t *testing.T) {
	tests := map[string]*unicode.RangeTable{
		"123456789-=?":  nil,
		"Hello, world!": unicode.Latin,
		"Привет всем!":  unicode.Cyrillic,
		"ქართული ენა მსოფლიო ":         unicode.Georgian,
		"県見夜上温国阪題富販":                   unicode.Han,
		" ككل حوالي 1.6، ومعظم الناس ": unicode.Arabic,
		"हिमालयी वन चिड़िया (जूथेरा सालिमअली) चिड़िया की एक प्रजाति है": unicode.Devanagari,
		"היסטוריה והתפתחות של האלפבית העברי":                            unicode.Hebrew,
		"የኢትዮጵያ ፌዴራላዊ ዴሞክራሲያዊሪፐብሊክ":                                     unicode.Ethiopic,
		"Привет! Текст на русском with some English.":                   unicode.Cyrillic,
		"Russian word любовь means love.":                               unicode.Latin,
		"আমি ভালো আছি, ধন্যবাদ!":                                        unicode.Bengali,
	}

	for text, want := range tests {
		got := DetectScript(text)
		require.Equalf(t, want, got, "text=%q", text)
	}
}

func TestIsLatin(t *testing.T) {
	tests := map[rune]bool{
		'z': true, 'A': true, 'č': true, 'š': true, 'Ĵ': true, 'ж': false,
	}

	for r, want := range tests {
		got := isLatin(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}

func TestIsEthiopic(t *testing.T) {
	tests := map[rune]bool{
		'ፚ': true, 'ᎀ': true, 'а': false, 'L': false,
	}

	for r, want := range tests {
		got := isEthiopic(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}

func TestIsGeorgian(t *testing.T) {
	tests := map[rune]bool{
		'რ': true, 'Я': false,
	}

	for r, want := range tests {
		got := isGeorgian(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}

func TestIsBengali(t *testing.T) {
	tests := map[rune]bool{
		'а': false, 'ই': true,
	}

	for r, want := range tests {
		got := isBengali(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}

func TestIsHiraganaKatakana(t *testing.T) {
	tests := map[rune]bool{
		'カ': true, 'Ґ': false,
		'ｴ': true, 'ᄁ': false,
		'ひ': true, 'Ꙕ': false,
		'ゐ': true, 'ф': false,
	}

	for r, want := range tests {
		got := isHiraganaKatakana(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}

func TestIsHangul(t *testing.T) {
	tests := map[rune]bool{
		'ᄁ': true, 't': false,
	}

	for r, want := range tests {
		got := isHangul(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}

func TestIsGreek(t *testing.T) {
	tests := map[rune]bool{
		'φ': true, 'ф': false,
	}

	for r, want := range tests {
		got := isGreek(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}

func TestIsKannada(t *testing.T) {
	tests := map[rune]bool{
		'ಡ': true, 'S': false,
	}

	for r, want := range tests {
		got := isKannada(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}

func TestIsTamil(t *testing.T) {
	tests := map[rune]bool{
		'ஐ': true, 'Ж': false,
	}

	for r, want := range tests {
		got := isTamil(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}

func TestIsThai(t *testing.T) {
	tests := map[rune]bool{
		'ก': true, '๛': true, 'Ґ': false,
	}

	for r, want := range tests {
		got := isThai(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}

func TestIsGujarati(t *testing.T) {
	tests := map[rune]bool{
		'ઁ': true, '૱': true, 'l': false,
	}

	for r, want := range tests {
		got := isGujarati(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}

func TestIsGurmukhi(t *testing.T) {
	tests := map[rune]bool{
		'ਁ': true, 'ੴ': true, 'Ж': false,
	}

	for r, want := range tests {
		got := isGurmukhi(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}

func TestIsTelugu(t *testing.T) {
	tests := map[rune]bool{
		'ఁ': true, '౿': true, 'l': false,
	}

	for r, want := range tests {
		got := isTelugu(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}

func TestIsOriya(t *testing.T) {
	tests := map[rune]bool{
		'ଐ': true, '୷': true, 'l': false,
	}

	for r, want := range tests {
		got := isOriya(r)
		require.Equalf(t, want, got, "r=%#U", r)
	}
}
