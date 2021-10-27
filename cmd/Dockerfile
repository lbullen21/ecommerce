FROM golang:1.17

WORKDIR /app
COPY . .

RUN go mod init github.com/lbullen21/ecommerce

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build main.go
# RUN go run main.go .

EXPOSE 3030

CMD ["./main"]

