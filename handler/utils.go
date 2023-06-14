package handler

import (
	"strings"
)

func CreateErrorResponseObject(errStr string) *ErrorResponse {
	output := &ErrorResponse{
		Message: errStr,
	}

	return output
}

// AuthorizationHeaderからtokenを取り出す
func ExtractBearerToken(authHeader string) string {
	if authHeader == "" {
		return ""
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || !strings.EqualFold(authParts[0], "bearer") {
		return ""
	}

	return authParts[1]
}
