FROM golang:1.17.0-alpine3.14
WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY main.go .
RUN go build -o kubevela-nocalhost-demo-app main.go

FROM alpine:3.10
WORKDIR /app
COPY --from=builder /app/kubevela-nocalhost-demo-app /app/kubevela-nocalhost-demo-app
ENTRYPOINT ./kubevela-nocalhost-demo-app

EXPOSE 9080

CMD ["./app"]
