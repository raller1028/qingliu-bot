FROM golang:1.19 as builder

WORKDIR /app

COPY . .
# CGO_ENABLED=0 一定不要忘记,会因为依赖库的问题导构建后无法运行
RUN CGO_ENABLED=0 go build -v -o /go/bin/lark-bot


# final (target) stage
# FROM scratch
FROM alpine

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/lark-bot /go/bin/lark-bot

LABEL VERSION="0.1"

EXPOSE 8000

ENTRYPOINT ["/go/bin/lark-bot"]