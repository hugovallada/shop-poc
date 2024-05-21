FROM golang:1.22 AS builder
RUN apt-get install ca-certificates
WORKDIR /app
COPY . . 
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server -ldflags="-w -s" .

FROM scratch AS app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]

