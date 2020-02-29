# scraping_moneyforward

マネーフォワード MEをスクレイピングし、今月の収支をslackに送信する。

## install

```shell script
# 環境変数のセットが必要。direnvを使うのが良いです。
$ cp .envrc.sample .envrc
$ vim .envrc # いい感じに

# バイナリを生成
$ make build
$ bin/scraping_moneyforward

# dockerでやるなら
$ make docker-build
$ make docker-run
```
