# 概要
<p>
2022年2月21日 - 2022年3月4日<br>
CA Tech Dojo Online で利用するAPIベース実装リポジトリ
</p>

CA Tech Dojo Onlineでは「スゴリらん！」というゲームのAPI実装を通してGo言語を使用したバックエンド開発のノウハウを学んでいきます。<br>
「スゴリらん！」は8つのAPIを必要とし、そのうち7つのAPIを期間中に実装していきます。<br>

課題の具体的な内容は[こちら](./TASKS.md)。

【ゲーム画面】
![ゲーム画面](./img/game_view.png)

【画面遷移図】
![画面遷移図](./img/transition.png)

## API仕様
API仕様はSwaggerUIを利用して閲覧します。
```
$ docker-compose up swagger-ui
```
を実行することでローカルのDocker上にSwaggerUIサーバが起動します。<br>
<br>
SwaggerUIサーバ起動後以下のURLからSwaggerUIへアクセスすることができます。

SwaggerUI: <http://localhost:3000/> <br> 
定義ファイル: `./api-document.yaml`<br>

# 事前準備
## goimportsとgolangci-lintのinstall
自分の書いたソースコードがプロジェクトのコード規約に則って記述されているか確認したり、整形したりするツールとして**gofmt** 、**goimports**、**golangci-lint**を使用します。<br>
**gofmt**はgoの標準コマンドであるため、goの環境構築を完了させていれば使用できるようになっていると思います。<br>
**goimports**、**golangci-lint**は別途installする必要があるため、以下のコマンドを実行しましょう。
```
$ make local-install
```
コードを整形するときは
```
$ make fmt
```
コード規約に則っているか確認するときは
```
$ make lint
```
を実行して開発を進めましょう。<br>
それぞれのコマンドが何を行っているか知りたいときはこのリポジトリ内のMakefileを見てみましょう。

## docker-composeを利用したMySQLとRedisの準備
### MySQL
MySQLはリレーショナルデータベースの1つです。
```
$ docker-compose up mysql
```
を実行することでローカルのDocker上にMySQLサーバが起動します。<br>
<br>
初回起動時に db/init ディレクトリ内のDDL, DMLを読み込みデータベースの初期化を行います。<br>
(DDL(DataDefinitionLanguage)とはデータベースの構造や構成を定義するためのSQL文)<br>
(DML(DataManipulationLanguage)とはデータの管理・操作を定義するためのSQL文)

#### PHPMyAdmin
MySQLデータベースのテーブルやレコードの閲覧、変更するためのツールとしてPHPMyAdminを用意しています。
```
$ docker-compose up phpmyadmin
```
を実行することでローカルのDocker上にPHPMyAdminサーバが起動します。<br>
PHPMyAdminサーバ起動後以下のURLからアクセスすることができます。

PHPMyAdmin: <http://localhost:4000/>

#### MySQLWorkbenchの設定
※ dockerの環境設定が上手くいかなかった場合に利用推奨<br>
Download: https://www.mysql.com/jp/products/workbench/

MySQLへの接続設定をします。
1. MySQL Connections の + を選択
2. 以下のように接続設定を行う
    ```
    Connection Name: 任意 (dojo_api等)
    Connection Method: Standard (TCP/IP)
    Hostname: 127.0.0.1 (localhost)
    Port: 3306
    Username: root
    Password: ca-tech-dojo
    Default Schema: dojo_api

#### API用のデータベースの接続情報を設定する
環境変数にデータベースの接続情報を設定します。<br>
ターミナルのセッション毎に設定したり、.bash_profileで設定を行います。

Macの場合
```
$ export MYSQL_USER=root \
    MYSQL_PASSWORD=ca-tech-dojo \
    MYSQL_HOST=127.0.0.1 \
    MYSQL_PORT=3306 \
    MYSQL_DATABASE=dojo_api
```

Windowsの場合<br>
※ それぞれの環境によって環境変数を設定するコマンドが異なる場合があるので注意
```
$ SET MYSQL_USER=root
$ SET MYSQL_PASSWORD=ca-tech-dojo
$ SET MYSQL_HOST=127.0.0.1
$ SET MYSQL_PORT=3306
$ SET MYSQL_DATABASE=dojo_api
```

#### Redis
Redisはインメモリデータベースの1つです。<br>
必須ではありませんが課題の中ででキャッシュやランキングなどの機能でぜひ利用してみましょう。<br>
```
$ docker-compose up redis
```
を実行することでローカルのDocker上にMySQLサーバが起動します。

# APIローカル起動方法
```
$ go run ./cmd/main.go
```

## ローカル環境のAPIを使用したゲームプレイ方法
自身が開発しているローカルのAPIを使用して、実際にゲームをプレイする方法は二つあります。場合によって使い分けてください。

**どうして二つあるの？**
今回使用するゲームのクライアントは、このリポジトリには含まれておらず、インターネット上に公開されています。インターネット上に公開されているアプリケーションが公開されている場所以外にアクセスしようとすると、セキュリティ上の問題があるため一部のブラウザではアクセスをブロックする設定になっています。(CORS)

それを防ぐために、インターネット上にあるものをあたかもローカルに存在するように見せるproxyを用意しました。

### docker-compose up proxyでプレイ
Dockerを利用してプレイする方法です。ブラウザに左右されずに動作させることができます。

```
$ docker-compose up -d proxy

// APIローカルも起動させる必要があります。
$ go run ./cmd/main.go
```

ブラウザから下記URLにアクセスしてください。
[http://localhost:3010/app](http://localhost:3010/app)

ID・パスワードはともに `ca-tech-dojo` です。
API接続先には `http://localhost:3010` と入力します。(ブラウザ直接利用の場合と異なるので注意！)

### ブラウザから直接プレイ
Dockerがうまく動かない場合には直接プレイすることもできます。ただし、ChromeやEdgeは利用できません。

- macOSユーザーの場合: Safari or Firefox
- Windowsユーザーの場合: Firefox

```
// APIローカルは起動させる必要があります
$ go run ./cmd/main.go
```

ブラウザから下記URLにアクセスしてください。
[http://13.114.176.9/](http://13.114.176.9/)

ID・パスワードはともに `ca-tech-dojo` です。
API接続先には `http://localhost:8080` と入力します。(proxy利用の場合と異なるので注意！)

## ビルド方法
作成したAPIを実際にをサーバ上にデプロイする場合は、<br>
ビルドされたバイナリファイルを配置して起動することでデプロイを行います。
### ローカルビルド
Macの場合
```
$ GOOS=linux GOARCH=amd64 go build -o dojo-api ./cmd/main.go
```

Windowsの場合
```
$ SET GOOS=linux
$ SET GOARCH=amd64
$ go build -o dojo-api ./cmd/main.go
```

このコマンドの実行で `dojo-api` という成果物を起動するバイナリファイルが生成されます。<br>
GOOS,GOARCHで「Linux用のビルド」を指定しています。
