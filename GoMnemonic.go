package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// Словарь BIP39
var bip39Words = []string{
	"abandon",
	"ability",
	"able",
	"about",
	"above",
	"absent",
	"absorb",
	"abstract",
	"absurd",
	"abuse",
}

func main() {
	// Генерация мнемонической фразы из 12 слов
	mnemonic := generateMnemonicPhrase(12)
	fmt.Println("Generated mnemonic phrase:", mnemonic)
}

// Функция для генерации мнемонической фразы
func generateMnemonicPhrase(wordCount int) string {
	var words []string
	for i := 0; i < wordCount; i++ {
		wordIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(bip39Words))))
		words = append(words, bip39Words[wordIndex.Int64()])
	}
	return strings.Join(words, " ")
}
