# はじめてのDockerイメージ
ハンズオンの中で使用するdockerコマンド一覧

``` sh
# dockerイメージの作成
#   -t = 作成するdockerイメージに名前をつけるオプション
#   . = 直下にあるDockerfileを暗黙的に指定
docker build -t go-server .

# 自分のマシン上に存在するdockerイメージを表示
docker images

# go-serverという名前のdockerイメージを実行してコンテナを作成
#   -d = バックグラウンドで起動するためのオプション
#   -p = ポートフォワーディング設定（ホスト側ポート番号:コンテナ側ポート番号）
docker run -d -p 8080:8080 go-server

# 実行されているdockerコンテナ一覧を表示
docker ps

# 実行されていないdockerコンテナも合わせて全て表示
#   -a = 全てのdockerコンテナと指定するためのオプション
docker ps -a

# 起動しているdockerコンテナを停止
docker stop [コンテナID]

# 停止されているdockerコンテナを全て削除
#   -f = フォース。（強制的に実行）
docker prune -f

# dockerイメージの削除
docker rmi [イメージID]

# ローカルに保存されているdockerイメージを全て削除
docker image prune -f --all
```
