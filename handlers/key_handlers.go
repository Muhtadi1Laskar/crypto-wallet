package handlers

import (
	"fmt"
	"net/http"
)

type KeyRequestBody struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type KeyResponse struct {
	Phrase string `json:"phrase"`
}

func GeneratePhrase(w http.ResponseWriter, r *http.Request) {
	var requestBody KeyRequestBody

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(requestBody.Name)
	fmt.Println(requestBody.Email)
}