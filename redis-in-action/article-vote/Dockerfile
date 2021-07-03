FROM golang:1.16-alpine as builder
WORKDIR /go/src/test
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN go build -o /go/bin/test .

FROM alpine as release
COPY --from=builder /go/bin/test /test
EXPOSE 5050
ENTRYPOINT ["/test"]