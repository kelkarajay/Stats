FROM golang:1.19

ARG BUILD_CMD

WORKDIR /app

ADD . .

COPY go.mod go.sum ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/${BUILD_CMD}/main.go 
RUN chmod +x app

EXPOSE 8080

# Run
CMD ["./app"]