package wallet

import (
	"crypto/rand"
	"fmt"
	"os"
	"strings"
)

func readWordList(filPath string) ([]string, error) {
	data, err := os.ReadFile(filPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read word list: %v", err)
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	return lines, nil
}

func bytesToBits(data []byte) []byte {
	bits := make([]byte, 0, len(data) * 8)
	for _, b := range data {
		bits = append(bits, bitToChar(b>>7))
		bits = append(bits, bitToChar(b>>6))
		bits = append(bits, bitToChar(b>>5))
		bits = append(bits, bitToChar(b>>4))
		bits = append(bits, bitToChar(b>>3))
		bits = append(bits, bitToChar(b>>2))
		bits = append(bits, bitToChar(b>>1))
		bits = append(bits, bitToChar(b))
	}
	return bits
}

func bitToChar(b byte) byte {
	return '0' + (b & 1)
}

func generateEntropy() ([]byte, error) {
	entropy := make([]byte, 16)
	_, err := rand.Read(entropy)
	if err != nil {
		return nil, fmt.Errorf("failed to generate entropy: %v", err)
	}
	return entropy, nil
}