# TODO
- 商品画像のリンクを貼りたい
- SNS認証を追加したい
  - 参考：https://qiita.com/momotaro98/items/90d12c10a655f4026d82
  - fireBaseなど
- 理想的なシステム図
  - https://qiita.com/shunp/items/abea7fa01e7a664c85da#firebase
- OR mapperを使いたい（データ構造の変化による変更箇所が増えるため）
- DB関係の処理は全てまとめてみた
    - 循環参照を避けるためにentityとdataを別にした
    - entityはdbのデータをオブジェクトで表現したもの
    - dataはtemplateに渡すentityオブジェクトと同じデータ構造で，templateで実行するメソッドを実装したもの
      - templateで実行するメソッドが不要なバックエンドAPIだけならばdataに分けなくていいはず
      - templateでメソッドを実行したかったので別で用意した
    - golangでは互いにimportし合うのは禁止されている
    - 結果的にメソッドにDB処理を持たせる方が楽だと思った

```bash
$ go mod init github.com/saanai/util-sys
$ go mod tidy
```