//go:build integration

package test

import (
	"net/http"
	"strings"
)

func (s *Suite) Test_AddItem() {
	tests := []struct {
		name       string
		path       string
		body       string
		wantStatus int
	}{
		{
			name:       "ok",
			path:       "/user/1007/cart/2008",
			body:       `{"count":1}`,
			wantStatus: http.StatusNoContent,
		},
		{
			name:       "product not found",
			path:       "/user/1007/cart/999999",
			body:       `{"count":1}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid user id",
			path:       "/user/abc/cart/2008",
			body:       `{"count":1}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid sku id",
			path:       "/user/1007/cart/abc",
			body:       `{"count":1}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid json body",
			path:       "/user/1007/cart/2008",
			body:       `{"count":`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "zero count",
			path:       "/user/1007/cart/2008",
			body:       `{"count":0}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "negative count",
			path:       "/user/1007/cart/2008",
			body:       `{"count":-1}`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			req, err := http.NewRequest(
				http.MethodPost,
				rootPath+tt.path,
				strings.NewReader(tt.body),
			)
			s.NoError(err)

			req.Header.Set("Content-Type", "application/json")

			resp, err := s.Client.Do(req)
			s.NoError(err)
			defer resp.Body.Close()

			s.Equal(tt.wantStatus, resp.StatusCode)
		})
	}

}
