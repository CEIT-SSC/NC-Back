FROM golang:1.17

WORKDIR /go/src/nc_back
COPY . .

RUN go build ./main.go

CMD ["./main"]