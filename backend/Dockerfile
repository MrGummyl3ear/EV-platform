FROM golang:1.21.7-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./
RUN go mod download
RUN go build -o EV-app ./cmd/main.go

CMD ["./EV-app"]