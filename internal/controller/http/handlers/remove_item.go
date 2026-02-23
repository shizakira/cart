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
		log.Warn().Err(err).Msg("strconv.Atoi")

		return
	}

	skuID, err := strconv.Atoi(chi.URLParam(r, "sku_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("strconv.Atoi")

		return
	}

	input := dto.RemoveItemInput{
		UserID: userID,
		SkuID:  skuID,
	}
	if err = input.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("input.Validate")

		return
	}

	if err = h.usecase.RemoveItem(r.Context(), input); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Err(err).Msg("usecase.RemoveItem")

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
