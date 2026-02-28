//go:build integration

package test

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (s *Suite) Test_Checkout() {
	req, err := http.NewRequest(
		http.MethodPost,
		rootPath+"/user/1007/cart/1076963",
		strings.NewReader(`{"count":1}`),
	)
	s.NoError(err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.Client.Do(req)
	s.NoError(err)
	resp.Body.Close()

	tests := []struct {
		name       string
		path       string
		wantStatus int
	}{
		{
			name:       "ok",
			path:       "/user/1007/cart/checkout",
			wantStatus: http.StatusOK,
		},
		{
			name:       "cart not found",
			path:       "/user/9999/cart/checkout",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid user id",
			path:       "/user/abc/cart/checkout",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			req, err := http.NewRequest(
				http.MethodPost,
				rootPath+tt.path,
				nil,
			)
			s.NoError(err)

			resp, err := s.Client.Do(req)
			s.NoError(err)
			defer resp.Body.Close()

			s.Equal(tt.wantStatus, resp.StatusCode)

			if tt.wantStatus == http.StatusOK {
				var body struct {
					OrderID int64 `json:"order_id"`
				}
				s.NoError(json.NewDecoder(resp.Body).Decode(&body))
				s.Greater(body.OrderID, int64(0))
			}
		})
	}
}
