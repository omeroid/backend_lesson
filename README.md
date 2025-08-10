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

### 5. コードの保存（コミット）
実装が一段落したら、こまめにコミットしましょう：

```bash
# 変更内容を確認
git status

# 変更をステージング
git add .

# コミット（メッセージは実装内容に合わせて変更）
git commit -m "feat: ユーザー認証機能を実装"
```

### 6. 答え合わせ
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
- **React** 18.3.1
- **React Router** v6.30.1
- **Material-UI** v5.16.7
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

## 補足資料
- [講義資料](https://docs.google.com/presentation/d/10xLeueQwx0gD7bsZ947B9Ukv_f-4tCtVBS-XD8l0SDA)
- [バックエンドの起動手順](https://github.com/omeroid/backend_lesson/blob/feat/readme/docs/backend.md)
- [フロントエンドの起動手順](https://github.com/omeroid/backend_lesson/blob/feat/readme/docs/frontend.md)
