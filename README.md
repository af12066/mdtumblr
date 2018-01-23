# mdtumblr

Markdown で書かれたファイルを Tumblr に投稿するだけの CLI ツール

## インストール

```bash
$ go install github.com/af12066/mdtumblr
```

## 使い方

### API キーの設定

Tumblr の API キーを以下の環境変数に設定してください:

- `TUMBLR_CLIENT_ID`: コンシューマキー
- `TUMBLR_CLIENT_SECRET`: コンシューマシークレット
- `TUMBLR_ACCESS_TOKEN`: OAuth 認証後に取得されるアクセストークン
- `TUMBLR_ACCESS_SECRET`: OAuth 認証後に取得されるシークレット

### 実行

```bash
$ mdtumblr ./README.md
$ mdtumblr -t PostTitle -u host-name ./README.md
```

## Option

- `-t`: ポストのタイトルを指定します。オプションがない場合は、コマンド実行後に対話形式で入力します。
- `-u`: ブログのホスト名を入力します。オプションがない場合は、コマンド実行後に対話形式で入力します。
- `-s`: ポストの状態を指定します。`published`, `draft`, `queue`, `private` のいずれかで、デフォルトは `published` です。

## License

MIT
