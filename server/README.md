# Mysqlをインストール

``brew install mysql2``

# 実行

``bundle install --path vendor/bundle``

# データベースのセットアップ

``make db/init ``

# sampleデータ入れ

``make db/seed_user``

# 仕様書
<!--
## ユーザ一覧取得
###  GET /users
### response 200
```
[
  {
    "id": 1,
    "name": "松岡裕典0",
    "sex": "female",
    "email": "sample0@sample.com",
    "birthday": "1996/0/05",
    "university": "同志社",
    "created_at": "2018-08-01T18:13:26.000Z",
    "updated_at": "2018-08-01T18:13:26.000Z"
  },
  {
    "id": 3,
    "name": "松岡裕典2",
    "sex": "female",
    "email": "sample2@sample.com",
    "birthday": "1996/2/05",
    "university": "同志社",
    "created_at": "2018-08-01T18:13:26.000Z",
    "updated_at": "2018-08-01T18:13:26.000Z"
  },
]
 ```
-->

## ログイン
### POST api/v1/login
入力

```
JSON

{"auth": {"email": "メアド", "password": "パスワード"}}

```

### response 201

出力
```
例
{
  "jwt": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1MzUzNDgyMzgsInN1YiI6MX0.z0jmRZXz2dbBVVCB_oZiS71ljrAuSNtHj1Q4qYSgTzI"
}
```

## ユーザ登録
### POST api/v1/singup

入力
```
name - 氏名 not null
sex - 性別(male/female) not null
email - メアド not null
password - パスワード
password_digest - パスワード確認用
birthday - 生年月日(xxxx/xx/xx) nut null
university - 大学
```

### response 201

出力
```
例
{
  "jwt": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1MzUzNDgyMzgsInN1YiI6MX0.z0jmRZXz2dbBVVCB_oZiS71ljrAuSNtHj1Q4qYSgTzI"
}
```

##ユーザ情報
### GET api/v1/user
### response 200

出力
```
{
    "user": {
        "id": 1,
        "name": "新垣結衣",
        "sex": "female",
        "email": "yuiyui1@sample.com",
        "birthday": "1988-06-11T00:00:00.000Z",
        "university": "けんたチュキチュキクラブ"
    },
    "plans": [
        {
            "id": 1,
            "user_id": 1,
            "date": "2018-07-16T00:00:00.000Z",
            "place": "大阪",
            "drink": "あり",
            "course": "高学歴コース",
            "status": 0,
            "created_at": "2018-08-26T05:24:22.000Z",
            "updated_at": "2018-08-26T05:24:22.000Z"
        },
        {
            "id": 2,
            "user_id": 1,
            "date": "2018-07-18T00:00:00.000Z",
            "place": "兵庫",
            "drink": "あり",
            "course": "低学歴コース",
            "status": 0,
            "created_at": "2018-08-26T05:24:22.000Z",
            "updated_at": "2018-08-26T05:24:22.000Z"
        }
    ]
}

```

## ユーザ情報編集
### PUT api/v1/user
### response 200

## ユーザ削除
### DELETE api/v1/user
### response 204

## 予定を取得
### GET api/v1/user/plans/:id
### response 200

## 予定の詳細を取得
### GET api/v1/user/plans/:id
### response 200

出力
```
{
    "id": 1,
    "user_id": 1,
    "date": "2018-07-16T00:00:00.000Z",
    "place": "大阪",
    "drink": "あり",
    "course": "高学歴コース",
    "status": 0,
    "created_at": "2018-08-26T05:24:22.000Z",
    "updated_at": "2018-08-26T05:24:22.000Z"
}
```

## 予定を追加
### POST api/v1/user/plans
### response 200

```
{
    "id": 1,
    "user_id": 1,
    "date": "2018-07-16T00:00:00.000Z",
    "place": "大阪",
    "drink": "あり",
    "course": "高学歴コース",
    "status": 0,
    "created_at": "2018-08-26T05:24:22.000Z",
    "updated_at": "2018-08-26T05:24:22.000Z"
}
```

## 予定の追加
### POST /user/plans
### response 201
### key

```
date - 日付 not null (xxxx-xx-xxT00:00:00.000Z)
place - 場所 not null
drink - 飲酒 not null
course - なんかコース的な not null
status - マッチングしたかしてないか not null
```

## 予定の変更
### PUT /user/plans/:id
### response 200

## 予定の削除
### DELETE /user/plans/:id
### response 204
