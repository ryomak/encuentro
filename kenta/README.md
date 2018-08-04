# Mysqlをインストール

``brew install mysql2``

# 実行

``bundle install --path vendor/bundle``

# データベースのセットアップ

``rails db:create``

``rails db:migrate``

# 仕様書

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

## 指定のユーザ取得
### GET /users/:id
### response 200

```
#id=1の時(/users/1)
{
  "id": 1,
  "name": "松岡裕典0",
  "sex": "female",
  "email": "sample0@sample.com",
  "birthday": "1996/0/05",
  "university": "同志社",
  "created_at": "2018-08-01T18:13:26.000Z",
  "updated_at": "2018-08-01T18:13:26.000Z"
}
```

## ユーザ登録
### POST /users
### response 201
### key
```
name - 氏名 not null
sex - 性別(male/female) not null
email - メアド not null
birthday - 生年月日(xxxx/xx/xx) nut null
university - 大学
```

## ユーザ情報編集
### PUT /users
### response 200

## ユーザ削除
### DELETE /users/:id
### response 204

## あるユーザーの予定を取得
### GET /users/:user_id/plans
### response 200

```
[
  {
    "id": 1,
    "user_id": 1,
    "date": "2018-10-01T00:00:00.000Z",
    "place": "ラブホテル",
    "drink": "あり",
    "course": "おっぱいコース",
    "status": 0,
    "created_at": "2018-08-03T16:24:38.000Z",
    "updated_at": "2018-08-03T16:24:38.000Z"
  },
  {
    "id": 2,
    "user_id": 1,
    "date": "2018-10-02T00:00:00.000Z",
    "place": "ラブホテル",
    "drink": "あり",
    "course": "おっぱいコース",
    "status": 0,
    "created_at": "2018-08-03T16:24:38.000Z",
    "updated_at": "2018-08-03T16:24:38.000Z"
  }
]
```

## 予定の詳細取得
### GET /users/:user_id/plans/:id
### response 200

```
#user_id=1,id=1の時(/users/1/plans/1)
{
  "id": 1,
  "user_id": 1,
  "date": "2018-10-01T00:00:00.000Z",
  "place": "ラブホテル",
  "drink": "あり",
  "course": "おっぱいコース",
  "status": 0,
  "created_at": "2018-08-03T16:24:38.000Z",
  "updated_at": "2018-08-03T16:24:38.000Z"
}
```

## 予定の追加
### POST /users/:user_id/plans
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
### PUT /users/:user_id/plans/:id
### response 200

## 予定の削除
### DELETE /users/:user_id/plans/:id
### response 204
