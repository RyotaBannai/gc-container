FROM golang:1.25 AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o hello-world .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/hello-world .
ENV PORT=8080

CMD ["./hello-world"]
