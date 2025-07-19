package handlers

import (
	wallet "crypto-wallet/crypto"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
)

type SignUpRequestBody struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type PhraseResponse struct {
	Phrase string `json:"phrase"`
	Seed string `json:"seed"`
}

func GeneratePhrase(w http.ResponseWriter, r *http.Request) {
	var requestBody SignUpRequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var password string = requestBody.Password
	fmt.Println(password)

	mnemonic, err := wallet.GeneratePhrase()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
	mnemonicStr := strings.Join(mnemonic, " ")
	seed := wallet.GenerateSeed(mnemonicStr, password)

	responseBody := PhraseResponse{
		Phrase: mnemonicStr,
		Seed: hex.EncodeToString(seed),
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}
