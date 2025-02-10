FROM golang:1.23

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

RUN go mod tidy

COPY . .

RUN go build -o server ./cmd/main/main.go

EXPOSE 8088

CMD ["/app/server"]
