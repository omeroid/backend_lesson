package db

import (
	"github.com/glebarez/sqlite"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// データベースに接続するための関数を定義
func InitDB(dbName string) (*gorm.DB, error) {
	// gormを使って、指定された名前のSQLiteデータベースに接続を試みます
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	// もし、何らかのエラーが発生した場合、nilとエラー内容を返します
	if err != nil {
		return nil, err
	}

	// エラーが発生しなければ、接続したデータベースの情報を返します
	return db, nil
}

// Echoフレームワークに対して、データベースに接続するためのミドルウェアを設定します。
func DBMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	// ミドルウェア関数を定義します。
	// この関数は、リクエストが行われる度に実行されます
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		// 上記で設定したDBに接続する処理を行います。
		return func(c echo.Context) error {

			// このリクエストのコンテキストに、DB接続をセットします。
			// これにより、これ以降の処理でデータベース操作を行いやすくなります
			c.Set("db", db)

			// 次のミドルウェアまたはルートハンドラにリクエストを渡します
			return next(c)
		}
	}
}
