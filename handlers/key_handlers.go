package handlers

import (
	wallet "crypto-wallet/crypto"
	"fmt"
	"net/http"
	"strings"
)

type SignUpRequestBody struct {
	Name  string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type PhraseResponse struct {
	Phrase string `json:"phrase"`
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
	fmt.Println(mnemonic)

	responseBody := PhraseResponse{
		Phrase: mnemonicStr,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}

