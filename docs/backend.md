## バックエンド起動手順

1. プロジェクトのルートディレクトリに `.env` ファイルを作成し、以下の内容を入力して保存します。
```
DATABASE_NAME=chatapp.sqlite
```

2. ユーザのホームディレクトリに `.sqliterc` ファイルを作成し、以下の内容を入力して保存します。これにより、SQLiteの外部キー制約が有効になります。
- ファイル作成場所  
Windows: C:/Users/{username}/.sqliterc  
mac: ~/.sqliterc

```
PRAGMA foreign_keys=ON;
```

3. プロジェクトの「backend」フォルダで以下のコマンドを実行し、依存パッケージをダウンロードします。
```
go get all
```

4. プロジェクトのルートディレクトリで以下のコマンドを実行すると、サーバが起動し、`localhost:1323` でアクセスできるようになります。
```
go run main.go
```

## Postmanセットアップ手順
Postmanは、API開発者やテストエンジニアがAPIをテスト、デバッグ、ドキュメント化するための人気のあるツールです。
この講座では作成したAPIのテストにPostmanを使用します。
このセクションではPostmanの使用方法について説明します。


1. [google drive](https://drive.google.com/drive/u/0/folders/1irK9IZINkFQeqF0tMxpxwc2qRcvBAtsL)から二つのjsonファイルをダウンロードします。

2. Postmanを開き、先ほどダウンロードした`collection.json`と`environment.json`をドラック&ドロップします。  

3. 画面左「Collections」と画面右上「No Environment」プルダウンに`backend_lesson`が追加されていれば成功です。

4. 最後に画面右上「No Environment」をクリックし、`No Environment`から`backend_lesson`に切り替えましょう

## Postmanの使用方法
### パラメータを設定
URLパラメータは、URL内の特定の部分に動的な値を埋め込むために使用されます。一般的に、URLパスやクエリ文字列の一部として指定されます。

例1：ユーザーIDを含むユーザープロフィールページのURL
```
https://example.com/users/123
```
この例では、`users`というパスの後にユーザーIDが指定されています。ユーザーIDはURLパラメータであり、この場合は`123`です。

例2：商品IDとカテゴリーIDを含む商品ページのURL
```
https://example.com/products/456?category=789
```
この例では、`products`というパスの後に商品IDが指定されており、クエリ文字列の`category`パラメータにカテゴリーIDが指定されています。

下の画像は、Postmanを使用してURLパラメータを指定する方法を示しています。URLの一部に`:roomId`というパラメータが含まれており、値を指定する必要があります。

Postmanの右側のパネルで、「Params」を選択し、その下にある「Path Variables」セクションに移動します。そこで、`roomId`とその値を指定します。

例えば、`roomId`の値を`123`に設定する場合、`roomId`の値の横に`123`を入力します。

このようにして、Postmanを使用してURLパラメータを指定し、APIエンドポイントに対してリクエストを送信することができます。

![image](https://github.com/omeroid/backend_lesson/assets/54432132/8516d581-57b8-4f83-aa24-06862c4b37b6)

### リクエストボディの設定

リクエストボディは、HTTPリクエストでサーバに送信されるデータです。主にPOSTやPUTメソッドで使用され、フォームデータやJSONデータなどの形式で表現されます。

下の画像は、チャットアプリにログインするためのリクエストです。このリクエストではボディにユーザ名、パスワードを指定しています。

Postmanの右側のパネルで、「Body」を選択し、その下にある「raw」をクリックし、bodyに含まれるデータを変更することができます。

![image](https://github.com/omeroid/backend_lesson/assets/54432132/9d985a78-b2d8-406b-8c58-5e3928ded122)

### リクエストの送信
上の2項目を設定すればリクエストを送信するだけです。URL横の「Send」ボタンをクリックしリクエストを送信してみましょう。
レスポンスのボディは画面下の「Body」から確認することができます。

![image](https://github.com/omeroid/backend_lesson/assets/54432132/e0144de7-39f3-4632-b49f-e0e5909ed095)

## API確認
上の「Postmanの使用方法」を参考にAPIを叩いてみましょう。

### /user/signup POST - アカウント作成
**成功リスポンス例**
```
{
    "id": 2,
    "name": "testuser",
    "createdAt": "2023-07-05T18:32:53.94961+09:00"
}
```

### /user/signin POST - ログイン
**成功リスポンス例**
```
{
    "userId": 2,
    "userName": "testuser",
    "token": "9b5098e5-f790-4431-beb2-ea861527f473"
}
```

> **Note**  
> リスポンスに含まれる「token」は次のリクエストを送信するために必要ですが、今回は自動で使用するようになっています。

### /rooms GET - チャットルーム一覧取得
**成功リスポンス例**
```
{
    "rooms": [
        {
            "id": 1,
            "name": "雑談",
            "description": "どんな話題でもOK!　雑談ルーム",
            "createdAt": "2023-07-05T17:14:16.712466+09:00"
        }
    ]
}
```

### /rooms POST - チャットルーム作成
**成功リスポンス例**
```
{
    "id": 2,
    "name": "testroom",
    "description": "description",
    "createdAt": "2023-07-05T18:34:27.977218+09:00"
}
```

### /rooms/{roomId} GET - チャットルーム詳細取得
**成功リスポンス例**
```
{
    "id": 2,
    "name": "testroom",
    "description": "description",
    "createdAt": "2023-07-05T17:28:08.196615+09:00"
}
```

> **Note**  
> Path VariablesのroomIdは「チャットルーム作成」のリスポンスを参照

### /rooms/{roomId}/messages POST - メッセージ送信
**成功リスポンス例**
```
{
    "id": 1,
    "text": "hello",
    "user": {
        "id": 2,
        "name": "testuser",
        "createdAt": "2023-07-05T17:16:17.976631+09:00"
    },
    "createdAt": "2023-07-05T18:38:26.401928+09:00"
}
```

> **Note**  
> userIdはログイン時に環境変数にセットされているので、編集しないでok

### /rooms/{roomId}/messages GET - メッセージ一覧取得
**成功リスポンス例**
```
{
    "messages": [
        {
            "id": 1,
            "text": "Welcome to the omeroid lecture!",
            "user": {
                "id": 1,
                "name": "omeroid",
                "createdAt": "2023-07-05T17:14:16.711721+09:00"
            },
            "createdAt": "2023-07-05T17:14:16.712848+09:00"
        },
        {
            "id": 2,
            "text": "hello",
            "user": {
                "id": 2,
                "name": "testuser",
                "createdAt": "2023-07-05T17:16:17.976631+09:00"
            },
            "createdAt": "2023-07-05T17:30:48.628124+09:00"
        },
    ]
}
```

### /rooms/{roomId}/messages/{messageId} GET - メッセージ削除
**成功リスポンス例**
```
{
    "id": 2,
    "text": "hello",
    "user": {
        "id": 2,
        "name": "testuser",
        "createdAt": "2023-07-05T17:16:17.976631+09:00"
    },
    "createdAt": "2023-07-05T18:38:26.401928+09:00"
}
```
