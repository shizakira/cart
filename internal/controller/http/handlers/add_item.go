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
		log.Warn().Err(err).Msg("user id strconv err")

		return
	}

	skuID, err := strconv.Atoi(chi.URLParam(r, "sku_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("sku id strconv err")

		return
	}

	var body Body
	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("body decode err")

		return
	}

	input := dto.AddItemInput{
		UserID: userID,
		SkuID:  skuID,
		Count:  body.Count,
	}

	if err = input.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn().Err(err).Msg("validate body err")

		return
	}

	if err = h.usecase.AddItem(r.Context(), input); err != nil {
		if errors.Is(err, model.ErrProductNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			log.Warn().Err(err).Msg("product not found err")

			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Warn().Err(err).Msg("usecase.AddItem err")

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
