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
	// Phrase string `json:"phrase"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
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

	privateKeyPEM := keys.PrivateKeyToPEM(privateKey)
	publicKeyPEM := keys.PublicKeyToPEM(publicKey)

	fmt.Println(privateKeyPEM)
	fmt.Println(publicKeyPEM)

	fmt.Println(requestBody.Name)
	fmt.Println(requestBody.Email)

	fmt.Println(hashFunction(requestBody.Name))
	fmt.Println(hashFunction(requestBody.Email))

	responseBody := KeyResponse{
		PrivateKey: privateKeyPEM,
		PublicKey: publicKeyPEM,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}
