package handlers

import (
	"crypto-wallet/wallet"
	"net/http"
)

type LoginRequestBody struct {
	Password          string `json:"password" validate:"required"`
	EncryptedMnemonic string `json:"encryptedMnemonic" validate:"required"`
}

type LoginResponseBody struct {
	Address    string `json:"address"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var requestBody LoginRequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	data := wallet.RetriveExistingWallet(requestBody.Password, requestBody.EncryptedMnemonic)

	responseBody := LoginResponseBody{
		Address:    data.Address,
		PublicKey:  data.PublicKey,
		PrivateKey: data.PrivateKey,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}
