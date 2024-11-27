package caesar_cipher

import (
	"log"
	"strings"
)

type CaesarCipher struct {
	shift int
}

func (s *CaesarCipher) detectShift(encryptedText string) {
	expectedFrequencies := map[rune]float64{
		'e': 12.7, 't': 9.06, 'a': 8.17, 'o': 7.51, 'i': 6.97, 'n': 6.75,
		's': 6.33, 'h': 6.09, 'r': 5.99, 'd': 4.25, 'l': 4.03, 'c': 2.78,
	}

	letterCount := make(map[rune]int)
	totalLetters := 0
	for _, char := range encryptedText {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			normalizedChar := char | 32
			letterCount[normalizedChar]++
			totalLetters++
		}
	}

	bestShift := 0
	bestScore := 0.0
	for shift := 0; shift < 26; shift++ {
		score := 0.0
		for char, count := range letterCount {
			decodedChar := (char-'a'-rune(shift)+26)%26 + 'a'
			score += float64(count) * expectedFrequencies[decodedChar]
		}
		if score > bestScore {
			bestScore = score
			bestShift = shift
		}
	}

	log.Printf("Deslocamento detectado: %d", bestShift)
	s.shift = bestShift
}

func (s *CaesarCipher) Decrypt(encryptedText string, shiftAutoDetect ...bool) string {
	if shiftAutoDetect != nil && shiftAutoDetect[0] == true {
		s.detectShift(encryptedText)
	}

	shift := s.shift % 26
	var decrypted strings.Builder
	for _, char := range encryptedText {
		if char >= 'a' && char <= 'z' {
			decrypted.WriteRune(((char-'a'-rune(shift)+26)%26 + 'a'))
		} else if char >= 'A' && char <= 'Z' {
			decrypted.WriteRune(((char-'A'-rune(shift)+26)%26 + 'A'))
		} else {
			decrypted.WriteRune(char)
		}
	}
	return decrypted.String()
}

func (s *CaesarCipher) Encrypt(plainText string) string {
	shift := s.shift % 26
	var encrypted strings.Builder
	for _, char := range plainText {
		if char >= 'a' && char <= 'z' {
			encrypted.WriteRune(((char-'a')+rune(shift))%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			encrypted.WriteRune(((char-'A')+rune(shift))%26 + 'A')
		} else {
			encrypted.WriteRune(char)
		}
	}
	return encrypted.String()
}

func NewCaeserCipher(shift int) *CaesarCipher {
	return &CaesarCipher{
		shift,
	}
}
