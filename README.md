# fotune-slipper

# Dockerの起動方法
- ビルド
`docker compose build`
- 起動
`docker compose up -d`
- Goのクライアント起動
`docker compose exec server bash`

- Vueのクライアント起動
`docker compose exec client /bin/sh`
`npm run serve`