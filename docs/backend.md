## バックエンド起動手順

1. プロジェクトのルートディレクトリに `.env` ファイルを作成し、以下の内容を入力して保存します。

   ```
   DATABASE_NAME=chatapp.sqlite
   ```

2. ユーザのホームディレクトリに `.sqliterc` ファイルを作成し、以下の内容を入力して保存します。これにより、SQLiteの外部キー制約が有効になります。

   ```
   PRAGMA foreign_keys=ON;
   ```

3. プロジェクトのルートディレクトリで以下のコマンドを実行し、依存パッケージをダウンロードします。

   ```
   go get github.com/omeroid/backend_backend_lesson
   ```

4. プロジェクトのルートディレクトリで以下のコマンドを実行すると、サーバが起動し、`localhost:1323` でアクセスできるようになります。

   ```
   go run main.go
   ```
## バックエンド起動確認
次のコマンドをコマンドラインで実行することで、バックエンドの動作を確認しましよう。

### /user/signup POST - サインアップ
**REQUEST**
```mac
curl --location 'http://localhost:1323/user/signup' \
--header 'Content-Type: application/json' \
--data '{
    "username": "testuser",
    "password": "test"
}'
```

**RESPONSE**
```
{
  "id": 2,
  "name": "testuser",
  "createdAt": "2023-06-19T10:29:02.464221+09:00"
}
```

### /user/signin POST - サインイン
**REQUEST**
```mac
curl --location 'http://localhost:1323/user/signin' \
--header 'Content-Type: application/json' \
--data '{
    "username":"testuser",
    "password":"test"
}'
```

**RESPONSE**
```
{
  "userId": 2,
  "userName": "testuser",
  "token": "7c26b436-01b7-415e-96f3-c164e37f3f1d"
}
```


### /rooms GET - ルーム情報全件取得

**REQUEST**
```mac
curl --location --request GET 'http://localhost:1323/rooms' \
--header 'Authorization: Bearer {token}'
```

> **Note**  
> Authorizationに含まれる{token}には「サインイン」時のリスポンスのtokenを使用してください

**RESPONSE**
```
{
  "Rooms": [
    {
      "id": 1,
      "name": "雑談",
      "description": "どんな話題でも OK! 雑談ルーム",
      "createdAt": "2023-06-19T10:28:46.054979+09:00"
    }
  ]
}
```

### /rooms POST - ルーム作成

**REQUEST**
```mac
curl --location 'http://localhost:1323/rooms' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {token}' \
--data '{
    "name":"testroom",
    "description":"description"
}'
```

**RESPONSE**
```
{
  "id": 2,
  "name": "testroom",
  "description": "description",
  "createdAt": "2023-06-19T10:38:30.888965+09:00"
}
```

### /rooms/{roomId} GET - 指定したIDのルームの情報取得
**REQUEST**
```mac
curl --location 'http://localhost:1323/rooms/{roomId}' \
--header 'Authorization: Bearer {token}'
```

> **Note**  
> urlに含まれる{roomId}には「ルーム作成」時のリスポンスのidを使用してください

**RESPONSE**
```
{
  "id": 2,
  "name": "test room",
  "description": "chat room",
  "createdAt": "2023-06-19T10:38:30.888965+09:00"
}
```

### /rooms/{roomId}/messages POST - 指定したルームでメッセージを送る
**REQUEST**
```mac
curl --location 'http://localhost:1323/rooms/{roomId}/messages' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {token}' \
--data '{
    "userId":{userId},
    "text":"hello"
}'
```

> **Note**  
> dataに含まれる{userId}には「ログイン」時のリスポンスのidを使用してください

**RESPONSE**
```
{
  "id": 2,
  "text": "hello",
  "user": {
    "id": 2,
    "name": "testuser",
    "createdAt": "2023-06-19T10:29:02.464221+09:00"
  },
  "createdAt": "2023-06-19T10:47:25.288945+09:00"
}
```

### /rooms/{roomId}/messages GET - 指定したルームのメッセージを全件取得
**REQUEST**
```mac
curl --location 'http://localhost:1323/rooms/{roomId}/messages' \
--header 'Authorization: Bearer {token}'
```

**RESPONSE**
```
{
  "messages": [
    {
      "id": 2,
      "text": "hello",
      "user": {
        "id": 2,
        "name": "testuser",
        "createdAt": "2023-06-19T10:29:02.464221+09:00"
      },
      "createdAt": "2023-06-19T10:47:25.288945+09:00"
    },
    {
      "id": 3,
      "text": "hello~",
      "user": {
        "id": 2,
        "name": "testuser",
        "createdAt": "2023-06-19T10:29:02.464221+09:00"
      },
      "createdAt": "2023-06-19T10:51:18.973032+09:00"
    }
  ]
}
```

### /rooms/{roomId}/messages/{messageId} GET - 指定したルームのメッセージを削除
**REQUEST**
```
curl -X DELETE -H "Authorization: Bearer 7c26b436-01b7-415e-96f3-c164e37f3f1d" http://localhost:1323/rooms/2/messages/3
```

**RESPONSE**
```mac
{
  "id": 3,
  "text": "Hello!!!!!!!",
  "user": {
    "id": 2,
    "name": "omeroid",
    "createdAt": "2023-06-19T10:29:02.464221+09:00"
  },
  "createdAt": "2023-06-19T10:51:18.973032+09:00"
}
```
