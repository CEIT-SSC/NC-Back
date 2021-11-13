FROM golang:1.17

WORKDIR /go/src/nc_back
COPY pkg .

RUN go build ./main.go

CMD ["./main"]