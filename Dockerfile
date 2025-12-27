FROM golang:1.25.5-alpine

WORKDIR /app

# install git (needed for go modules sometimes)
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

EXPOSE 8080

CMD ["./server"]
