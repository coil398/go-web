FROM golang:1.17-stretch AS dev
COPY . /go/src/web
WORKDIR /go/src/web
RUN apt update -y && apt upgrade -y && apt install -y git mysql-client \
    && go get -u \
    && apt clean \
    && rm -rf /var/lib/apt/lists/* \
    && go get github.com/cosmtrek/air

ENV SERVER_PORT=3000

FROM golang:1.17-stretch AS build
COPY . /go/src/web
WORKDIR /go/src/web
RUN apt update -y && apt upgrade -y && apt install -y git \
    && go get -u \
    && GOOS=linux GOARCH=amd64 go build -o app . \
    && apt clean && rm -rf /var/lib/apt/lists/*

FROM gcr.io/distroless/base
WORKDIR /go/src/web
COPY --from=build /go/src/web/app .

EXPOSE 3000
ENV SERVER_PORT=3000
EXPOSE 3000
CMD ["./app"]
