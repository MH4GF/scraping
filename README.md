# scraping_moneyforward

マネーフォワード MEをスクレイピングし、今月の収支をslackに送信する。

## 動作環境

```shell script
- ChromeDriver 84.0.4147.30 (48b3e868b4cc0aa7e8149519690b6f6949e110a8-refs/branch-heads/4147@{#310})
- go1.14.1
```

## deploy

### deploy to Heroku Container Registry

chromedriverのことをあんまり考えないようにするために、Docker Imageを使ってデプロイしている。  
[Heroku Scheduler](https://devcenter.heroku.com/articles/scheduler)を利用してスケジュール実行する。  
heroku.ymlでアドオンのインストールまで指定しているので、デプロイしてスケジュール間隔の設定をすれば良い

```shell script
$ heroku login
$ heroku container:login
$ heroku container:push worker
$ heroku container:release worker

# herokuのダッシュボードで環境変数とスケジュール間隔を設定
```

### deploy to AWS Fargate with Terraform

WIP

## build

```shell script
# 環境変数のセットが必要。direnvを使うのが良いです。
$ cp .envrc.sample .envrc
$ vim .envrc

# バイナリを生成して実行
$ make build
$ bin/scraping

# dockerでビルドと実行もできる
$ make docker-build
$ make docker-run
```

