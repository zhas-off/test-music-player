FROM golang
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go
CMD ["/app/main"]