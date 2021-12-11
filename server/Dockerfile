FROM golang:latest

RUN mkdir /go/src/work
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/work
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/work
#COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

#CMD ["app"]