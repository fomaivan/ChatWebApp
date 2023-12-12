package handler

import (
	"chat/internal/entity"
	"net/http"
	"strings"
)

var AUTHZ_HEADER_NAME = "Authorization"

func (h *Handler) GetСlaimsFromAuthHeader(r *http.Request) (*map[string]string, error) {
	jwtClaims := &map[string]string{}
	authzHeader := r.Header.Get(AUTHZ_HEADER_NAME)
	if authzHeader == "" {
		return jwtClaims, entity.ErrEmptyAuthHeader
	}
	headerParts := strings.Split(authzHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return jwtClaims, entity.ErrInvalidAuthHeader
	}
	jwtClaims, err := h.auth.FetchAuth(headerParts[1])
	return jwtClaims, err
}
