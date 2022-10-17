FROM golang:1.16-alpine
ADD . /go-git
WORKDIR /go-git
RUN go build ./main.go
CMD ["./main"]