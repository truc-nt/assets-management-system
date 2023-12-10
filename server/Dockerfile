FROM golang:1.20
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod vendor

COPY . ./

RUN go build -o server cmd/main.go

EXPOSE 8080

CMD ["./server"]


