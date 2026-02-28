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

func (h Handlers) Checkout(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("strconv.Atoi")
		return
	}

	input := dto.CheckoutInput{UserID: userID}
	if err = input.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("input.Validate")
		return
	}

	output, err := h.usecase.Checkout(r.Context(), input)
	if errors.Is(err, domain.ErrCartNotFound) || errors.Is(err, domain.ErrCartIsEmpty) {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("usecase.Checkout")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Err(err).Msg("usecase.Checkout")
		return
	}

	render.JSON(w, output, http.StatusOK)
}
