package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/shizakira/cart/internal/dto"
	"github.com/shizakira/cart/internal/model"
)

func (h Handlers) AddItem(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		Count int `json:"count"`
	}

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

	var body Body
	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("json.NewDecoder")

		return
	}

	input := dto.AddItemInput{
		UserID: userID,
		SkuID:  skuID,
		Count:  body.Count,
	}

	if err = input.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("input.Validate")

		return
	}

	if err = h.usecase.AddItem(r.Context(), input); err != nil {
		if errors.Is(err, model.ErrProductNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			log.Warn().Err(err).Msg("usecase.AddItem")

			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Err(err).Msg("usecase.AddItem")

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
