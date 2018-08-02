## Mysqlをインストール

``brew install mysql2``

## 実行

``bundle install --path vendor/bundle``

## データベースのセットアップ

``rails db:create``

``rails db:migrate``

## 仕様書

### ユーザ一覧取得
###  GET /users
### response 200
```
[
    {
        "id": 1,
        "name": "松岡裕典0",
        "sex": "female",
        "mail": "sample0@sample.com",
        "birthday": "1996/0/05",
        "university": "同志社",
        "created_at": "2018-08-01T18:13:26.000Z",
        "updated_at": "2018-08-01T18:13:26.000Z"
    },
    {
        "id": 3,
        "name": "松岡裕典2",
        "sex": "female",
        "mail": "sample2@sample.com",
        "birthday": "1996/2/05",
        "university": "同志社",
        "created_at": "2018-08-01T18:13:26.000Z",
        "updated_at": "2018-08-01T18:13:26.000Z"
    },
 ]
 ```
 
 ### 指定のユーザ取得
 ### GET /users/:id
 ### response 200
 
 ```
 #id=1の時(/users/1)
 {
        "id": 1,
        "name": "松岡裕典0",
        "sex": "female",
        "mail": "sample0@sample.com",
        "birthday": "1996/0/05",
        "university": "同志社",
        "created_at": "2018-08-01T18:13:26.000Z",
        "updated_at": "2018-08-01T18:13:26.000Z"
    }
 
 ```
 
 ### ユーザ登録
 ### POST /users
 ### response 201
 ### key
 ```
 name - 氏名 not null
 sex - 性別(male/female) not null
 mail - メアド not null
 birthday - 生年月日(xxxx/xx/xx) nut null
 university - 大学
 
 ```
 
 ### ユーザ情報編集
 ### PUT /users
 ### response 200
 
