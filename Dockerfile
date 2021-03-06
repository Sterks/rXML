FROM golang:latest as builder

WORKDIR /app
ADD . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o rXML

FROM alpine:latest AS production
WORKDIR /app
ENV APPLICATION=production
COPY --from=builder /app .
CMD ["./rXML"]
