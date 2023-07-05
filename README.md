# Backend Lesson (Chat App)

## 概要
このリポジトリは、Golangを使用したWebバックエンドの講座用サンプルプログラムです。  
Chatアプリの実装を通じて、GolangやWebのアーキテクチャについて学ぶことができます。

## Chatアプリ概要
### 必要な環境
- Node.js (バージョン: 18.16.1）
- Go（バージョン: 1.20.5）
- SQLite（バージョン: 3.42.0）

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
- [バックエンドの起動手順](https://github.com/omeroid/backend_lesson/blob/feat/readme/docs/backend.md)
- [フロントエンドの起動手順](https://github.com/omeroid/backend_lesson/blob/feat/readme/docs/frontend.md)
