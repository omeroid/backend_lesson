## バックエンド起動手順
1 `git clone https://github.com/omeroid/backend_lesson.git`する

2 `.env`をプロジェクトのルートディレクトリに作成し、`DATABASE_NAME=chatapp.sqlite`と入力し保存する

3 ホームディレクトリ(Macなら`~`、Windowsなら`C:\Users\{ユーザ名}`)に`.sqliterc`を作成し`PRAGMA foreign_keys=ON;`と入力し保存(sqliteの外部キー制約をonにするため)

4 プロジェクトのルートディレクトリで`go get github.com/omeroid/backend_backend_lesson`を実行し、依存パッケージをダウンロードする

5 プロジェクトのルートディレクトリで`go run main.go`を実行すると`localhost:1323`でサーバが起動する
## 各エンドポイントに対するリクエストとレスポンスの例(curlコマンド)

### /user/signup POST　サインアップ
REQUEST
```
curl -X POST -H "Content-Type: application/json" -d '{
"userName": "omeroid",
"password": "backend"
}
http://localhost:1323/user/signup
```

RESPONSE 
```
{
"id":2,
"name":"omeroid",
"createdAt":"2023-06-19T10:29:02.464221+09:00"
}
```
  
### /user/signin POST　サインイン
REQUEST
```
curl -X POST -H "Content-Type: application/json" -d '{
"userName": "omeroid",
"password": "backend"
}' http://localhost:1323/user/signin
```

RESPONSE
```
{
"userId":2,
"userName":"omeroid",
"token":"7c26b436-01b7-415e-96f3-c164e37f3f1d"
}
```

### /rooms GET　ルーム情報全件取得
REQUEST
```
curl -X GET -H "Authorization: Bearer 7c26b436-01b7-415e-96f3-c164e37f3f1d" http://localhost:1323/rooms
```

RESPONSE
```
{
"Rooms":[{"id":1,
"name":"雑談",
"description":"どんな話題でも OK!　雑談ルーム",
"createdAt":"2023-06-19T10:28:46.054979+09:00"}]
}
```

### /rooms POST　ルーム作成
REQUEST
```
curl -X POST -H "Authorization: Bearer 7c26b436-01b7-415e-96f3-c164e37f3f1d" -H "Content-Type: application/json" -d 
'{
"name": "test room",
"description": "chat room"
}'  http://localhost:1323/rooms
```

RESPONSE
```
{
"id":2,
"name":"test room",
"description":"chat room",
"createdAt":"2023-06-19T10:38:30.888965+09:00"
}
```

### /rooms/{roomId} GET　指定したIDのルームの情報取得
REQUEST
```
curl -X GET -H "Authorization: Bearer 7c26b436-01b7-415e-96f3-c164e37f3f1d" http://localhost:1323/rooms/2
```

RESPONSE
```
{
"id":2,
"name":"test room",
"description":"chat room",
"createdAt":"2023-06-19T10:38:30.888965+09:00"
}
```

### /rooms/{roomId}/messages POST　指定したルームでメッセージを送る
REQUEST
```
curl -X POST -H "Authorization: Bearer 7c26b436-01b7-415e-96f3-c164e37f3f1d" -H "Content-Type: application/json" -d '{
"userId": 2,
"text": "Hello!"
}'  http://localhost:1323/rooms/2/messages
```

RESPONSE
```
{
"id":2,
"text":"Hello!",
"user":{
  "id":2,
  "name":"omeroid",
  "createdAt":"2023-06-19T10:29:02.464221+09:00"
},
"createdAt":"2023-06-19T10:47:25.288945+09:00"
}
```

### /rooms/{roomId}/messages GET　指定したルームのメッセージを全件取得
REQUEST 
```
curl -X GET -H "Authorization: Bearer 7c26b436-01b7-415e-96f3-c164e37f3f1d" http://localhost:1323/rooms/2/messages
```

RESPONSE
```
{"messages":[{
"id":2,
"text":"Hello!",
"user":{
  "id":2,
  "name":"omeroid",
  "createdAt":"2023-06-19T10:29:02.464221+09:00"
},
"createdAt":"2023-06-19T10:47:25.288945+09:00"
},
{
”id":3,
"text":"Hello!!!!!!!",
"user":{
  "id":2,
  "name":"omeroid",
  "createdAt":"2023-06-19T10:29:02.464221+09:00"
},
"createdAt":"2023-06-19T10:51:18.973032+09:00"}]
}
```

### /room/{roomId}/messages/{messageId} GET　指定したルームのメッセージを削除
REQUEST
```
curl -X GET -H "Authorization: Bearer 7c26b436-01b7-415e-96f3-c164e37f3f1d" http://localhost:1323/rooms/2/messages/3
```

RESPONSE
```
{
"id":3,
"text":"Hello!!!!!!!",
"user":{
  "id":2,
  "name":"omeroid",
  "createdAt":"2023-06-19T10:29:02.464221+09:00"
},
"createdAt":"2023-06-19T10:51:18.973032+09:00"
}
```
 
 
