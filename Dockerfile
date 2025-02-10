FROM golang:1.23

WORKDIR /internal
COPY go.mod go.sum ./
COPY cmd .

RUN go mod tidy

COPY . .

RUN go build -o server ./cmd/main/main.go

EXPOSE 8088

CMD ["/internal/server"]
