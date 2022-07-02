package handler

import (
	"fmt"
	"net/http"
	"test-project/internal/model"

	"github.com/gorilla/mux"
)

const STRING_IDENTIFIER_LENGTH = 16

func (h *Handler) createRandomIdentifier(w http.ResponseWriter, r *http.Request) {
	stringIdentifier := generateRandStringBytes(STRING_IDENTIFIER_LENGTH)
	description := generateRandStringBytes(2 * STRING_IDENTIFIER_LENGTH)
	input := model.Identifier{
		StringIdentifier: stringIdentifier,
		Description:      description,
	}

	stringIdentifier, err := h.services.Identifier.CreateRandom(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("Created string_identifier: %s", stringIdentifier)
	fmt.Fprint(w, response)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getIdentifierDescription(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringIdentifier := vars["string_identifier"]

	description, err := h.services.Identifier.GetDescription(stringIdentifier)
	if err != nil {
		// newErrorResponse(w, http.StatusInternalServerError, err.Error())
		response := fmt.Sprintf("string_identifier '%s' not found", stringIdentifier)
		fmt.Fprint(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("string_identifier '%s' has description '%s'", stringIdentifier, description)
	fmt.Fprint(w, response)
	w.WriteHeader(http.StatusOK)
}
