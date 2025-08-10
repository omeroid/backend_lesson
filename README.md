# Backend Lesson (Chat App)

## 概要
このリポジトリは、Golangを使用したWebバックエンドの講座用サンプルプログラムです。  
Chatアプリの実装を通じて、GolangやWebのアーキテクチャについて学ぶことができます。

## 🎯 講座の構成

### ブランチ構成
- **`main`ブランチ**: 完成形のコード（解答例）
- **`question`ブランチ**: ハンズオン用の穴埋めコード（受講者が実装する部分が空欄になっています）

### 講座の流れ
1. 環境構築
2. `question`ブランチから作業用ブランチを作成
3. 穴埋め部分を実装しながらGolangとバックエンド開発を学習
4. 完成後、`main`ブランチのコードと比較して答え合わせ

## 🚀 セットアップ手順

### 1. 必要な環境の準備
以下のツールを事前にインストールしてください：

- **Node.js** (バージョン: 18.0.0以上、推奨: 20.x LTS）
  - インストール確認: `node --version`
  - [ダウンロード](https://nodejs.org/)
  
- **Go**（バージョン: 1.23.0以上）
  - インストール確認: `go version`
  - [ダウンロード](https://go.dev/dl/)
  
- **SQLite**（バージョン: 3.38.0以上）
  - インストール確認: `sqlite3 --version`
  - macOS: `brew install sqlite`
  - Ubuntu/Debian: `sudo apt-get install sqlite3`
  - Windows: [ダウンロード](https://www.sqlite.org/download.html)

- **Git**
  - インストール確認: `git --version`
  - [ダウンロード](https://git-scm.com/)

### 2. リポジトリのクローン
```bash
# HTTPSの場合
git clone https://github.com/omeroid/backend_lesson.git

# SSHの場合（推奨）
git clone git@github.com:omeroid/backend_lesson.git

# クローンしたディレクトリに移動
cd backend_lesson
```

### 3. 依存関係のインストール

#### バックエンド（Go）のセットアップ
```bash
# backendディレクトリに移動
cd backend

# Go依存関係のインストール
go mod download

# インストールの確認
go mod verify
```

#### フロントエンド（Node.js）のセットアップ
```bash
# プロジェクトルートに戻る
cd ..

# frontendディレクトリに移動
cd frontend

# npm依存関係のインストール
npm install
```

## 📝 ハンズオンの進め方

### 1. 作業用ブランチの作成
講座を始める前に、`question`ブランチから自分の作業用ブランチを作成します：

```bash
# プロジェクトルートに移動
cd backend_lesson

# questionブランチに切り替え
git checkout question

# 最新の状態を取得
git pull origin question

# 自分の作業用ブランチを作成（名前は自由に変更してください）
git checkout -b feature/your-name-handson

# ブランチが切り替わったことを確認
git branch
# * feature/your-name-handson となっていればOK
```

### 2. 穴埋め箇所の確認
`question`ブランチでは、実装が必要な箇所に以下のようなコメントが記載されています：

```go
// TODO: ここに実装を追加
// HINT: ユーザー認証のロジックを実装してください
```

これらのTODOコメントを探して、順番に実装していきます。

### 3. 実装の進め方

#### TODO箇所を検索する
```bash
# バックエンドのTODOを検索
cd backend
grep -r "TODO" --include="*.go" .

# または、VSCodeなどのエディタの検索機能を使用
```

#### 実装のヒント
- 各TODOには実装のヒントが記載されています
- わからない場合は`main`ブランチの同じファイルを参考にできます
- 講義資料のスライドも参照してください

### 4. 動作確認

#### バックエンドの起動
```bash
# backendディレクトリで実行
cd backend
go run cmd/main.go
# サーバーが http://localhost:1323 で起動します
```

#### フロントエンドの起動
別のターミナルウィンドウを開いて：
```bash
# frontendディレクトリで実行
cd frontend
npm start
# アプリが http://localhost:3000 で起動します
```

### 5. テストユーザーでログイン

アプリケーションを起動後、以下のテストユーザーでログインできます：

- **ユーザー名**: `omeroid`
- **パスワード**: `backend`

### 6. コードの保存（コミット）
実装が一段落したら、こまめにコミットしましょう：

```bash
# 変更内容を確認
git status

# 変更をステージング
git add .

# コミット（メッセージは実装内容に合わせて変更）
git commit -m "feat: ユーザー認証機能を実装"
```

### 7. 答え合わせ
実装が完了したら、`main`ブランチのコードと比較して答え合わせができます：

```bash
# 特定のファイルを比較
git diff main backend/path/to/file.go

# または、GitHubで比較
# https://github.com/omeroid/backend_lesson/compare/question...main
```

## 🔧 トラブルシューティング

### よくある問題と解決方法

#### 1. ポートが既に使用されている場合
```bash
# 1323ポートを使用しているプロセスを確認（macOS/Linux）
lsof -i :1323

# プロセスを終了
kill -9 [PID]

# フロントエンドの3000ポートも同様に確認
lsof -i :3000
```

#### 2. データベースエラーが発生する場合
```bash
# backendディレクトリで
rm -f chat.db  # 既存のDBファイルを削除
go run cmd/main.go  # 再起動で新しいDBが作成されます
```

#### 3. npmインストールでエラーが出る場合
```bash
# node_modulesとpackage-lock.jsonを削除
rm -rf node_modules package-lock.json

# キャッシュをクリア
npm cache clean --force

# 再インストール（依存関係の互換性の問題がある場合）
npm install --legacy-peer-deps
```

## 📦 使用技術スタック

### バックエンド
- **Go** 1.23.0+
- **Echo** v4.13.4 (Webフレームワーク)
- **GORM** v1.30.1 (ORM)
- **SQLite** (データベース)
- **golang.org/x/crypto** (パスワードハッシュ化)

### フロントエンド
- **React** 19.1.1
- **React Router** v7.8.0
- **Material-UI (MUI)** v6.3.0
- **Axios** v1.11.0 (HTTP クライアント)
- **SWR** v2.3.5 (データフェッチング)
- **React Hot Toast** v2.5.2 (通知)

## 📚 Chatアプリ概要

### システム概要

![Chatアプリのシステム概要図](https://github.com/omeroid/backend_lesson/assets/54432132/c3140af9-adde-40a4-917e-4c729fee7c87)

### 機能一覧

| 機能 | HTTPメソッド | URL |
| --- | --- | --- |
| ユーザ登録 | POST | /user/signup |
| ログイン | POST | /user/signin |
| チャットルーム一覧取得 | GET | /rooms |
| チャットルーム詳細取得 | GET | /rooms/{roomId} |
| チャットルーム作成 | POST | /rooms |
| メッセージ送信 | POST | /rooms/{roomId}/messages |
| メッセージ削除 | DELETE | /chatRooms/{roomId}/messages/{messageId} |
| メッセージ一覧取得 | GET | /rooms/{roomId}/messages/ |

## 📝 ハンズオン問題集

### 問題概要
`question`ブランチには、以下の3つの機能を実装する課題があります。各機能はメッセージ管理に関するものです。

### 問題1: メッセージ一覧取得機能の実装
**ファイル**: `backend/cmd/main.go`
**内容**: チャットルーム内のメッセージ一覧を取得するAPIエンドポイントを追加します。

- HTTPメソッド: GET
- パス: `/rooms/:roomId/messages`
- ハンドラー: `handler.ListMessage`

---

### 問題2: メッセージ作成機能の実装
**ファイル**: `backend/cmd/main.go`
**内容**: チャットルームに新しいメッセージを投稿するAPIエンドポイントを追加します。

- HTTPメソッド: POST
- パス: `/rooms/:roomId/messages`
- ハンドラー: `handler.CreateMessage`

---

### 問題3: メッセージ削除機能の実装
**ファイル**: `backend/cmd/main.go`
**内容**: 特定のメッセージを削除するAPIエンドポイントを追加します。

- HTTPメソッド: DELETE
- パス: `/rooms/:roomId/messages/:messageId`
- ハンドラー: `handler.DeleteMessage`

---

### 問題A: メッセージ作成処理の実装
**ファイル**: `backend/handler/message.go`
**関数名**: `CreateMessage`
**内容**: メッセージを作成するハンドラー関数を実装します。

**実装手順**:
1. データベース接続を取得（`c.Get("db")`を`*gorm.DB`に型変換）
2. Authorizationヘッダーからトークンを取得し、`util.ExtractBearerToken`で抽出
3. `IsSessionValid`でトークンを検証
4. リクエストボディから`CreateMessageInput`を取得
5. URLパラメータから`roomId`を取得し、`strconv.Atoi`で整数に変換
6. UserIDに該当するユーザーをデータベースから検索
7. メッセージをデータベースに保存
8. `CreateMessageOutput`を作成して返却（`http.StatusCreated`）

---

### 問題B: メッセージ一覧取得処理の実装
**ファイル**: `backend/handler/message.go`
**関数名**: `ListMessage`
**内容**: 指定されたルームの全メッセージを取得するハンドラー関数を実装します。

**実装手順**:
1. データベース接続を取得（`c.Get("db")`を`*gorm.DB`に型変換）
2. Authorizationヘッダーからトークンを取得し、`util.ExtractBearerToken`で抽出
3. `IsSessionValid`でトークンを検証
4. URLパラメータから`roomId`を取得し、`strconv.Atoi`で整数に変換
5. 指定されたルームIDのメッセージをデータベースから取得
6. 各メッセージにユーザー情報を追加して`ListMessageOutput`を作成
7. レスポンスを返却（`http.StatusOK`）

---

### 問題C: メッセージ削除処理の実装
**ファイル**: `backend/handler/message.go`
**関数名**: `DeleteMessage`
**内容**: 指定されたメッセージを削除するハンドラー関数を実装します。

**実装手順**:
1. データベース接続を取得（`c.Get("db")`を`*gorm.DB`に型変換）
2. Authorizationヘッダーからトークンを取得し、`util.ExtractBearerToken`で抽出
3. `IsSessionValid`でトークンを検証
4. URLパラメータから`messageId`と`roomId`を取得し、`strconv.Atoi`で整数に変換
5. 指定されたメッセージをデータベースから削除
6. レスポンスを返却（`http.StatusNoContent`）

---

### 💡 実装のポイント
- 必要なパッケージのインポートを忘れずに（コメントアウトされているものを有効化）
- エラーハンドリングを適切に実装
- HTTPステータスコードを正しく使用（作成:201、取得:200、削除:204）
- トークン検証を忘れずに実装

## 🔧 デバッグ方法と動作確認

### 1. コンパイルエラーの確認
```bash
# backendディレクトリで実行
cd backend
go build ./cmd/main.go

# エラーが出た場合の対処法
# - インポート忘れ: undefined エラーが出たパッケージをインポート
# - 型エラー: 型変換を確認（例: c.Get("db").(*gorm.DB)）
# - 関数名のタイポ: 大文字小文字を確認
```

### 2. サーバー起動の確認
```bash
# バックエンドサーバーを起動
go run cmd/main.go

# 正常に起動した場合の表示
# ⇨ http server started on [::]:1323

# よくあるエラー
# - "bind: address already in use": 既に起動中なので、既存プロセスを停止
# - "panic: runtime error": ルーティング設定のミスを確認
```

### 3. APIエンドポイントの動作確認

#### テスト用アカウントでログイン
```bash
# まずログインしてトークンを取得
curl -X POST http://localhost:1323/user/signin \
  -H "Content-Type: application/json" \
  -d '{"userName":"omeroid","password":"backend"}'

# レスポンス例
# {"id":1,"name":"omeroid","token":"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"}
# このtokenを以降のリクエストで使用
```

#### 問題1: メッセージ一覧取得のテスト
```bash
# メッセージ一覧を取得（tokenは上記で取得したものを使用）
curl -X GET http://localhost:1323/rooms/1/messages \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"

# 期待されるレスポンス
# {"messages":[{"id":1,"text":"Welcome to the omeroid lecture!","user":{"id":1,"name":"omeroid"},"createdAt":"..."}]}

# エラーの場合
# 401 Unauthorized: トークンが無効 → Authorization ヘッダーを確認
# 404 Not Found: エンドポイントが未実装 → main.goのルーティングを確認
# 500 Internal Server Error: DB接続エラー → handler内の実装を確認
```

#### 問題2: メッセージ作成のテスト
```bash
# メッセージを投稿
curl -X POST http://localhost:1323/rooms/1/messages \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{"userId":1,"text":"テストメッセージです"}'

# 期待されるレスポンス（201 Created）
# {"id":2,"text":"テストメッセージです","user":{"id":1,"name":"omeroid"},"createdAt":"..."}

# エラーの場合
# 400 Bad Request: リクエストボディの形式エラー → CreateMessageInput構造体を確認
# 401 Unauthorized: トークンが無効
# 404 Not Found: エンドポイントが未実装
```

#### 問題3: メッセージ削除のテスト
```bash
# メッセージを削除（メッセージID:2を削除する例）
curl -X DELETE http://localhost:1323/rooms/1/messages/2 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"

# 期待されるレスポンス（204 No Content）
# ボディは空

# エラーの場合
# 401 Unauthorized: トークンが無効
# 404 Not Found: エンドポイントが未実装 または メッセージが存在しない
# 500 Internal Server Error: DB削除処理のエラー
```

### 4. ログを使ったデバッグ

実装中にデバッグが必要な場合、以下のようにログを追加：

```go
// handler/message.go内でデバッグ用ログを追加
func CreateMessage(c echo.Context) error {
    conn := c.Get("db").(*gorm.DB)
    
    // デバッグ: DB接続確認
    fmt.Println("DB接続取得成功")
    
    authHeader := c.Request().Header.Get("Authorization")
    token := util.ExtractBearerToken(authHeader)
    
    // デバッグ: トークン確認
    fmt.Printf("取得したトークン: %s\n", token)
    
    // ... 以下実装
}
```

### 5. フロントエンドからの動作確認

バックエンドとフロントエンドを両方起動した状態で確認：

```bash
# ターミナル1: バックエンド起動
cd backend
go run cmd/main.go

# ターミナル2: フロントエンド起動
cd frontend
npm start
```

ブラウザで http://localhost:3000 にアクセスして確認：
1. omeroid/backend でログイン
2. チャットルームを選択
3. メッセージの表示・投稿・削除が正常に動作するか確認

### 6. よくあるエラーと対処法

| エラー | 原因 | 対処法 |
|-------|------|--------|
| `undefined: handler.ListMessage` | 関数が未実装 | handler/message.goに関数を実装 |
| `cannot use conn (type interface {}) as type *gorm.DB` | 型アサーション忘れ | `.(*gorm.DB)`を追加 |
| `undefined: strconv` | インポート忘れ | import文に"strconv"を追加 |
| `http: panic serving` | nilポインタアクセス | エラーチェックを追加 |
| `Error 1: no such table: messages` | DB初期化失敗 | chat.dbを削除して再起動 |

### 🎯 答え合わせ
実装が完了したら、`main`ブランチの同じファイルと比較して答え合わせをしてください：
```bash
git diff main backend/cmd/main.go
git diff main backend/handler/message.go
```

## 📖 補足資料
- [講義資料（スライド）](https://docs.google.com/presentation/d/10xLeueQwx0gD7bsZ947B9Ukv_f-4tCtVBS-XD8l0SDA)
- [問題文・ハンズオン手順書](https://docs.google.com/document/d/1lBfKX0FiuU1kX6Njy5Ss0Cdu3FvXSs8Kd7PeH9fPGAA)
- [バックエンドの起動手順](https://github.com/omeroid/backend_lesson/blob/feat/readme/docs/backend.md)
- [フロントエンドの起動手順](https://github.com/omeroid/backend_lesson/blob/feat/readme/docs/frontend.md)
