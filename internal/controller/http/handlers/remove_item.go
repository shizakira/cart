package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/shizakira/cart/internal/dto"
)

func (h Handlers) RemoveItem(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("user id strconv err")

		return
	}

	skuID, err := strconv.Atoi(chi.URLParam(r, "sku_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("sku id strconv err")

		return
	}

	input := dto.RemoveItemInput{
		UserID: userID,
		SkuID:  skuID,
	}
	if err = input.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("validate body err")

		return
	}

	if err = h.usecase.RemoveItem(r.Context(), input); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Err(err).Msg("validate body err")

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
