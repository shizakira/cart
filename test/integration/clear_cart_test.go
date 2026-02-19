//go:build integration

package test

import (
	"net/http"
	"strings"
)

func (s *Suite) Test_ClearCart_OK() {
	reqAdd, _ := http.NewRequest(
		http.MethodPost,
		rootPath+"/user/4001/cart/1076963",
		strings.NewReader(`{"count":3}`),
	)
	reqAdd.Header.Set("Content-Type", "application/json")

	respAdd, err := s.Client.Do(reqAdd)
	s.NoError(err)
	respAdd.Body.Close()
	s.Equal(http.StatusNoContent, respAdd.StatusCode)

	reqClear, _ := http.NewRequest(
		http.MethodDelete,
		rootPath+"/user/4001/cart",
		nil,
	)

	respClear, err := s.Client.Do(reqClear)
	s.NoError(err)
	respClear.Body.Close()

	s.Equal(http.StatusNoContent, respClear.StatusCode)

	reqGet, _ := http.NewRequest(
		http.MethodGet,
		rootPath+"/user/4001/cart",
		nil,
	)

	respGet, err := s.Client.Do(reqGet)
	s.NoError(err)
	respGet.Body.Close()

	s.Equal(http.StatusNotFound, respGet.StatusCode)
}

func (s *Suite) Test_ClearCart_Empty_OK() {
	req, _ := http.NewRequest(
		http.MethodDelete,
		rootPath+"/user/5001/cart",
		nil,
	)

	resp, err := s.Client.Do(req)
	s.NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusNoContent, resp.StatusCode)
}
