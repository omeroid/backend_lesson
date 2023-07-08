package util

import (
	"path/filepath"
	"runtime"
	"strings"
)

// AuthorizationHeaderからtokenを取り出す関数を定義します。
func ExtractBearerToken(authHeader string) string {
	// 入力された文字列(authHeader)をスペースで分割し、その結果をauthPartsという変数に格納します。
	authParts := strings.Fields(authHeader)
	// 分割した結果が二つの要素でない、または最初の要素が "bearer" でない場合は空の文字列を戻します。
	if len(authParts) != 2 || !strings.EqualFold(authParts[0], "bearer") {
		return ""
	}

	// 上記の条件を満たす場合は二つ目の要素（トークン）を戻します。
	return authParts[1]
}

// パスとバックエンドのルートディレクトリを連結する関数を定義します。
func JoinWithBackendRoot(path string) string {
	// Caller関数を使用して、呼び出し元の情報を取得します。
	// この場合、関心があるのはファイルのパス(b)のみなので、他の返り値は無視します。
	_, b, _, _ := runtime.Caller(0)

	// utilのディレクトリパスを取得します。
	utilPath := filepath.Dir(b)

	// utilが存在するpkgのディレクトリパスを取得します。
	pkgPath := filepath.Dir(utilPath)

	// backendのディレクトリパスを取得します。
	backendPath := filepath.Dir(pkgPath)

	// backendのディレクトリパスと引数に渡されたパスを結合します。
	return filepath.Join(backendPath, path)
}
