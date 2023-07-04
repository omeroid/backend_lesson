# Backend Lesson (Chat App)

***
## 概要
このリポジトリは、Golangを使用したWebバックエンドの講座用サンプルプログラムです。  
Chatアプリの実装を通じて、GolangやWebのアーキテクチャについて学ぶことができます。

***
## Chatアプリ概要
### 必要な環境
- Node.js (バージョン: 18.16.1）
- Go（バージョン: 1.20.5）
- SQLite（バージョン: 3.42.0）

### システム概要
以下の図は、Chatアプリのシステム構成を示しています。

![Chatアプリのシステム概要図](https://github.com/omeroid/backend_lesson/assets/54432132/c3140af9-adde-40a4-917e-4c729fee7c87)

### 機能一覧
以下の表は、Chatアプリの機能一覧を示しています。

| 機能 | HTTPメソッド | URL |
| --- | --- | --- |
| ユーザ登録 | POST | /user/signup |
| ログイン | POST | /user/signin |
| チャットルーム一覧取得 | GET | /chatRooms |
| チャットルーム詳細取得 | GET | /chatRooms/{chatRoomId} |
| チャットルーム作成 | POST | /chatRooms |
| メッセージ送信 | POST | /message |
| メッセージ削除 | DELETE | /message |
| メッセージ一覧取得 | GET | /message/{chatRoomId} |

***
## 補足資料
- バックエンドの起動手順
- フロントエンドの起動手順
