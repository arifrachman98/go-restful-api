package middleware

import (
	"net/http"

	"github.com/arifrachman98/go-restful-api/helper"
	"github.com/arifrachman98/go-restful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(h http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: h}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "RAHASIA" == r.Header.Get("X-API-Key") {
		//ok
		m.Handler.ServeHTTP(w, r)
	} else {
		//error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
