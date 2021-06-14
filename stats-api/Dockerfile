FROM golang:latest AS builder

WORKDIR /app
COPY . ./
RUN go mod download
RUN go build -o stats-server cmd/stats-server/main.go

FROM ubuntu:20.04

RUN apt-get update && apt-get -y upgrade && apt-get -y install ca-certificates
COPY --from=builder /app/stats-server /usr/local/bin/stats-server

EXPOSE 8000
ENV PORT 8000
USER nobody
CMD ["stats-server", "--scheme=http", "--host=0.0.0.0", "--port=8000"]

