FROM golang:alpine AS builder

#为镜像设置环境变量
ENV Go111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录 /app
WORKDIR /app

COPY . .

RUN go build -tags "doc" -o main main.go
RUN go build -o /app/migrate /app/cmd/migrate/migrate.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/wait-for.sh .
COPY --from=builder /app/migrate .
COPY --from=builder /app/internal/dao/mysql/migration ./migration
COPY --from=builder /app/main .
COPY --from=builder /app/start.sh .

RUN chmod +x wait-for.sh
RUN chmod +x start.sh

EXPOSE 8080

#ENTRYPOINT ["/app/start.sh"]
#CMD ["/app/main"]
