FROM golang:1.17

WORKDIR /app
COPY . .

RUN go mod init github.com/lbullen21/ecommerce

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build main.go

CMD ["./main"]

# EXPOSE 3030