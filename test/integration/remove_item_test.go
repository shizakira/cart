//go:build integration

package test

import (
	"net/http"
	"strings"
)

func (s *Suite) Test_RemoveItem_OK() {
	reqAdd, _ := http.NewRequest(
		http.MethodPost,
		rootPath+"/user/2001/cart/2008",
		strings.NewReader(`{"count":2}`),
	)
	reqAdd.Header.Set("Content-Type", "application/json")

	respAdd, err := s.Client.Do(reqAdd)
	s.NoError(err)
	respAdd.Body.Close()
	s.Equal(http.StatusNoContent, respAdd.StatusCode)

	reqDel, err := http.NewRequest(
		http.MethodDelete,
		rootPath+"/user/2001/cart/2008",
		nil,
	)
	s.NoError(err)

	respDel, err := s.Client.Do(reqDel)
	s.NoError(err)
	respDel.Body.Close()

	s.Equal(http.StatusNoContent, respDel.StatusCode)

	reqGet, _ := http.NewRequest(
		http.MethodGet,
		rootPath+"/user/2001/cart",
		nil,
	)

	respGet, err := s.Client.Do(reqGet)
	s.NoError(err)
	respGet.Body.Close()

	s.Equal(http.StatusNotFound, respGet.StatusCode)
}
