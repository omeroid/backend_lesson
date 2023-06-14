package handler

import (
	"encoding/json"
	"strings"
)

func ThrowError(errStr string) string {
	res := ErrorResponse{
		Message: errStr,
	}

	var output []byte
	output, _ = json.Marshal(res)
	return string(output)
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
