FROM golang:1.16-buster as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /app/mybin main.go

FROM alpine
WORKDIR /

EXPOSE 8000

COPY --from=builder /app/mybin /app

CMD /app