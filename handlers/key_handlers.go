package handlers

import (
	crypto "crypto-wallet/crypto"
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

	mnemonic, err := crypto.GeneratePhrase()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
	mnemonicStr := strings.Join(mnemonic, " ")
	seed := crypto.GenerateSeed(mnemonicStr, password)

	responseBody := PhraseResponse{
		Phrase: mnemonicStr,
		Seed: hex.EncodeToString(seed),
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}
