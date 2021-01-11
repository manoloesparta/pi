FROM golang:1.13-alpine

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080 

CMD ["/build/main"]
