# Docker Hubを使ってみよう
ハンズオンの中で使用するdockerコマンド一覧

``` sh
# dockerイメージにタグ（名前）を付ける
docker tag [イメージID] [Docker HubのアカウントID]/[イメージの名前]

# Docker HubにイメージをPush（保存）
docker push [Docker HubのアカウントID]/[イメージの名前]

# Docker Hubからイメージをダウンロード
# 今回は、「takumaron/othlo_whale」
docker pull [ダウンロードしたいイメージ名]
```
