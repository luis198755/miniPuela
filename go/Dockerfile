FROM arm32v6/golang:1.21-alpine

WORKDIR /go
RUN go mod init miniPuebla

#COPY go.mod go.mod
COPY main.go main.go
RUN go mod tidy -v
#RUN go mod tidy

RUN go build -o main .

CMD ["./miniPuebla"]