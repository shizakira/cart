//go:build integration

package test

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (s *Suite) Test_GetItems_EmptyCart() {
	req, _ := http.NewRequest(
		http.MethodGet,
		rootPath+"/user/9999/cart",
		nil,
	)

	resp, err := s.Client.Do(req)
	s.NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusNotFound, resp.StatusCode)
}

func (s *Suite) Test_GetItems_OK() {
	req1, _ := http.NewRequest(http.MethodPost,
		rootPath+"/user/1001/cart/2008",
		strings.NewReader(`{"count":2}`),
	)
	req1.Header.Set("Content-Type", "application/json")

	resp1, err := s.Client.Do(req1)
	s.NoError(err)
	resp1.Body.Close()
	s.Equal(http.StatusNoContent, resp1.StatusCode)

	req2, _ := http.NewRequest(
		http.MethodPost,
		rootPath+"/user/1001/cart/2958025",
		strings.NewReader(`{"count":1}`),
	)
	req2.Header.Set("Content-Type", "application/json")

	resp2, err := s.Client.Do(req2)
	s.NoError(err)
	resp2.Body.Close()
	s.Equal(http.StatusNoContent, resp2.StatusCode)

	req, err := http.NewRequest(
		http.MethodGet,
		rootPath+"/user/1001/cart",
		nil,
	)
	s.NoError(err)

	resp, err := s.Client.Do(req)
	s.NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusOK, resp.StatusCode)

	var body struct {
		Items []struct {
			SkuID int    `json:"sku_id"`
			Name  string `json:"name"`
			Count int    `json:"count"`
			Price int    `json:"price"`
		} `json:"items"`
		TotalPrice int `json:"total_price"`
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	s.NoError(err)
	s.Len(body.Items, 2)
	s.True(body.Items[0].SkuID < body.Items[1].SkuID)

	for _, item := range body.Items {
		s.NotZero(item.Name)
		s.Positive(item.Count)
		s.Positive(item.Price)
	}

	s.Positive(body.TotalPrice)
}
