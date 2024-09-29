# Todo API
 Goの練習

## 作る機能
- todoのcrud操作
  - todoの作成: `POST /todo`
  - todoの一覧取得: 
    - ページ指定での取得: `GET /todo/page=`
    - id指定での取得: `GET /todo/{id}`
  - todoの更新: `PATCH /todo/{id}`
  - todoの削除: `DELETE /todo/{id}`

- ユーザ管理機能
  - jwtを使用して自前で作る
  - google認証

## 開発用コマンド
`docker image build -t [name:tag] [path]`: appイメージのビルド   
`docker container run --rm [imageName:tag]`: appコンテナの起動  

`docker compose up -d`: DBコンテナ立ち上げ  
`docker compose down`: DBコンテナを落とす  
`docker compose exec db psql -U postgres -d mydb`: DBコンテナに入る  

`make apply`: atlasでの理想スキーマをDBに反映  
`make migrate`: atlasでのスキーママイグレーションの更新  
`make inspect`: DBのスキーマをatlasでの理想スキーマファイルに反映 

