FROM golang:1.17

WORKDIR /go/src/nc_back
COPY . .

#RUN go run ./main.go

CMD ["go","run","main.go"]