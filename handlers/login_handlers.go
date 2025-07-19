package handlers

import (
	"crypto-wallet/wallet"
	"net/http"
)

type LoginRequestBody struct {
	Password          string `json:"password" validate:"required"`
	EncryptedMnemonic string `json:"encryptedMnemonic" validate:"required"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var requestBody LoginRequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	wallet.RetriveExistingWallet(requestBody.Password, requestBody.EncryptedMnemonic)
}