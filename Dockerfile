FROM golang as builder

WORKDIR /app

ADD . /app

RUN go build -o bin/server server/main.go

FROM alpine

COPY --from=builder /app/bin/server .

EXPOSE 8000

CMD ["./server"]
