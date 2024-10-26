# Todo API

Go の練習

# 開発用コマンド

## データベース関連

- DB の起動・終了

  - `docker compose up -d`:
  - `docker compose down`:

- マイグレーション関連

  - `atlas schema inspect --env dev`: 現在のスキーマを確認
  - `atlas schema apply --env dev`: 理想スキーマの適応
  - `atlas migrate diff --env dev`: 現在のスキーマと理想スキーマの差分を保存
  - `atlas migrate hash --env dev`: マイグレーションファイルのハッシュを生成
  - `atlas migrate validate --env dev`: ハッシュ検証

- postgres 関連
  - `psql -h 127.0.0.1 -U postgres -d mydb`: postgres クライアントからコンテナ DB に接続
  - `dt`: 現存のリレーション一覧の表示
  - `d [tablename]`: リレーションの詳細を表示
