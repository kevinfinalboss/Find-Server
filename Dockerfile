FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o app

EXPOSE 8080

HEALTHCHECK --interval=5s \
            --timeout=3s \
            CMD curl --fail http://localhost:8080/health || exit 1

CMD ["go", "build"]