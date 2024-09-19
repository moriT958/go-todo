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
