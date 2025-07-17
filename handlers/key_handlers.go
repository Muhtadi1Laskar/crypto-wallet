package handlers

import (
	keys "crypto-wallet/Keys"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

type KeyRequestBody struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type KeyResponse struct {
	Phrase string `json:"phrase"`
}

func hashFunction(text string) string {
	byteMessage := []byte(text)
	hash := sha256.New()
	hash.Write(byteMessage)

	hashedBytes := hash.Sum(nil)
	encodedStr := hex.EncodeToString(hashedBytes)

	return encodedStr
}

func GeneratePhrase(w http.ResponseWriter, r *http.Request) {
	var requestBody KeyRequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	privateKey, publicKey, err := keys.GenerateKeys()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(privateKey)
	fmt.Println(publicKey)

	fmt.Println(requestBody.Name)
	fmt.Println(requestBody.Email)

	fmt.Println(hashFunction(requestBody.Name))
	fmt.Println(hashFunction(requestBody.Email))
}
