package util

import (
	"path/filepath"
	"runtime"
	"strings"
)

// AuthorizationHeaderからtokenを取り出す
func ExtractBearerToken(authHeader string) string {
	authParts := strings.Fields(authHeader)
	if len(authParts) != 2 || !strings.EqualFold(authParts[0], "bearer") {
		return ""
	}

	return authParts[1]
}

func JoinWithBackendRoot(path string) string {
	_, b, _, _ := runtime.Caller(0)
	utilPath := filepath.Dir(b)
	pkgPath := filepath.Dir(utilPath)
	backendPath := filepath.Dir(pkgPath)
	return filepath.Join(backendPath, path)
}
