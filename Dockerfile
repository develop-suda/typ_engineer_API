# ベースとなるDockerイメージ指定
FROM golang:1.18 as build

ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV GOBIN=$GOPATH/bin
ENV PATH $PATH:$GOROOT:$GOPATH:$GOBIN
ENV GO111MODULE=on
ENV TZ=Asia/Tokyo

# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/typ_engineer_api
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/typ_engineer_api
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/typ_engineer_api

RUN apt-get update
RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs
RUN go get github.com/ramya-rao-a/go-outline
RUN go get github.com/cweill/gotests/...
RUN go get github.com/fatih/gomodifytags
RUN go get github.com/josharian/impl
RUN go get github.com/haya14busa/goplay/cmd/goplay
RUN go get github.com/go-delve/delve/cmd/dlv
RUN go get github.com/golangci/golangci-lint/cmd/golangci-lint
RUN go get golang.org/x/tools/gopls
# div-dap のインストール方法は次のドキュメントを参考にしました:
# https://github.com/golang/vscode-go/blob/v0.26.0/docs/dlv-dap.md#updating-dlv-dap
RUN GOBIN=/tmp/ go get github.com/go-delve/delve/cmd/dlv@master
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/joho/godotenv
RUN go get -u github.com/golang-module/carbon/v2

RUN go mod tidy
# Airをインストール
RUN go install github.com/cosmtrek/air@v1.27.3

# airコマンドでGoファイルを起動
CMD ["air"]