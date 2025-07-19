package handlers

import (
	"crypto-wallet/wallet"
	"net/http"
)

type ImportRequestBody struct {
	Phrase   string `json:"phrase" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ImportResponseBody struct {
	Address    string `json:"address"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

func ImportHandlers(w http.ResponseWriter, r *http.Request) {
	var requestBody ImportRequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := wallet.ImportWallet(requestBody.Phrase, requestBody.Password)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseBody := LoginResponseBody{
		Address:    data.Address,
		PublicKey:  data.PublicKey,
		PrivateKey: data.PrivateKey,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}
