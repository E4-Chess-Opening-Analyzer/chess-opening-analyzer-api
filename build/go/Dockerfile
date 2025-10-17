FROM golang:1.24-alpine3.21

LABEL maintainer="Luc Lacotte <luc.lonqq@gmail.com>"
LABEL description="golang:1.24-alpine3.21 with git and gow installed"
LABEL version="1.0"

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/mitranim/gow@latest

ENV PATH="/go/bin:${PATH}"

CMD ["go", "run", "src/main.go"]