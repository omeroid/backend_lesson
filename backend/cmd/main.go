package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/omeroid/backend_lesson/backend/handler"
	"github.com/omeroid/backend_lesson/backend/pkg/db"
	"github.com/omeroid/backend_lesson/backend/pkg/util"
)

// プログラムのエントリーポイント（最初に実行される部分）を定義しています。
func main() {
	e := echo.New() // Echoインスタンスを作成。EchoはWebアプリケーションを効率よく作成するためのGo言語用ライブラリーです。

	databaseFilePath := util.JoinWithBackendRoot("chatapp.sqlite") // データベースへのパスを指定します。
	var hasDatabaseFile bool                                       // データベースファイルが存在するかどうかをチェックするフラグを定義します。
	if _, err := os.Stat(databaseFilePath); err == nil {
		hasDatabaseFile = true // データベースファイルが存在する場合、フラグを真に設定します。
	}

	conn, err := db.InitDB(databaseFilePath) // データベースへ接続します。
	if err != nil {
		e.Logger.Fatalf("DBの接続失敗: %s", err.Error()) // データベースへの接続が失敗した場合にログを吐き出し、プログラムを終了します。
	}

	if !hasDatabaseFile { // データベースファイルが存在しない場合、以下の操作を行います。
		if err := db.Migrate(conn); err != nil {
			e.Logger.Fatalf("DBへのmigration失敗: %s", err.Error()) // データベースへのマイグレーションが失敗した場合にログを吐き出し、プログラムを終了します。
		}
		if err := db.InsertSampleRecord(conn); err != nil {
			e.Logger.Fatalf("DBへのサンプルレコードのinsert失敗: %s", err.Error()) // サンプルレコードの挿入が失敗した場合にログを吐き出し、プログラムを終了します。
		}
		e.Logger.Print("Migration Successful") // マイグレーションが成功した際のログを表示します。
	}

	e.Use(db.DBMiddleware(conn)) // DBへの接続情報を含むミドルウェアをEchoインスタンスに登録します。
	e.Use(middleware.BodyDump(func(ctx echo.Context, reqBody, resBody []byte) {
		if 0 < len(reqBody) {
			fmt.Printf("Request Body: %v\n", string(reqBody))
		}
		if 0 < len(resBody) {
			fmt.Printf("Response Body: %v\n", string(resBody))
		}
	})) // BodyDumpミドルウェアをEchoインスタンスに登録します。これによりリクエストやレスポンスのボディがロギングされます。
	e.Use(middleware.Logger())  // LoggerミドルウェアをEchoインスタンスに登録します。これによりリクエストやレスポンスの情報がロギングされます。
	e.Use(middleware.Recover()) // RecoverミドルウェアをEchoインスタンスに登録します。これによりパニック時にサーバーがクラッシュするのを防ぎます。

	// CORS(Cross-Origin Resource Sharing)の設定を行います。これにより、ブラウザが安全に異なるオリジンからリソースを取得することが可能になります。
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},                                                // 全てのオリジンからのアクセスを許可します。
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete}, // これらのHTTPメソッドの使用を許可します。
	}))

	// それぞれのパスとHTTPメソッドに対して実行するハンドラーを定義します。
	e.POST("/user/signup", handler.SignUp)
	e.POST("/user/signin", handler.SignIn)
	e.GET("/rooms", handler.ListRoom)
	e.POST("/rooms", handler.CreateRoom)
	e.GET("/rooms/:roomId", handler.GetRoom)

	// TODO 1. メッセージ一覧取得機能を実装してください
	// (ヒント: handler.ListMessageを実装してください)

	// TODO 2. メッセージ作成機能を実装してください
	// (ヒント: handler.CreateMessageを実装してください)

	// TODO 3. メッセージの削除機能を実装してください。
	// (ヒント: handler.DeleteMessageを実装してください)

	e.Logger.Fatal(e.Start(":1323")) // サーバーを1323ポートで起動します。なお、サーバー起動時にエラーが発生した場合はログを出力してプログラムを終了します。
}
