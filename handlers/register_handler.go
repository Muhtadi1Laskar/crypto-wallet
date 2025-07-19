package handlers

import (
	"crypto-wallet/wallet"
	"fmt"
	"net/http"
)

type SignUpRequestBody struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type PhraseResponse struct {
	Phrase string `json:"phrase"`
	Address string `json:"address"`
	EncryptedMnemonic string `json:"encryptedMnemonic"`
	EncryptedPrivateKey string `json:"encryptedPrivateKey"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var requestBody SignUpRequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var password string = requestBody.Password
	fmt.Println(password)

	data := wallet.CreateNewWallet(password)

	responseBody := PhraseResponse{
		Phrase: data.Phrase,
		Address: data.Address,
		EncryptedMnemonic: data.EncryptedMnemonic,
		EncryptedPrivateKey: data.EncryptedPrivateKey,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}
