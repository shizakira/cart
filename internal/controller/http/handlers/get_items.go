package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/shizakira/cart/internal/domain"
	"github.com/shizakira/cart/internal/dto"
	"github.com/shizakira/cart/pkg/render"
)

func (h Handlers) GetItems(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("user id strconv err")

		return
	}

	input := dto.GetItemsInput{UserID: userID}
	if err = input.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("validate body err")

		return
	}

	output, err := h.usecase.GetItems(r.Context(), input)
	if err != nil {
		if errors.Is(err, domain.ErrCartNotFound) || errors.Is(err, domain.ErrCartIsEmpty) {
			w.WriteHeader(http.StatusNotFound)
			log.Info().Err(err).Msg("cart err")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error().Err(err).Msg("cart err")
		}

		return
	}

	render.JSON(w, output, http.StatusOK)
}
