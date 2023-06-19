package handler

import "strings"

// AuthorizationHeaderからtokenを取り出す
func ExtractBearerToken(authHeader string) string {
	authParts := strings.Fields(authHeader)
	if len(authParts) != 2 || !strings.EqualFold(authParts[0], "bearer") {
		return ""
	}

	return authParts[1]
}
